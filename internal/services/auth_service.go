// Package services package services
package services

import (
	"errors"
	"log"
	"time"

	"github.com/JavascriptDev347/learning-go-shop/internal/config"
	"github.com/JavascriptDev347/learning-go-shop/internal/dto"
	"github.com/JavascriptDev347/learning-go-shop/internal/models"
	"github.com/JavascriptDev347/learning-go-shop/internal/utils"
	"gorm.io/gorm"
)

// AuthService struct includes db and config, and we need it to connect to another file with authservice
type AuthService struct {
	db     *gorm.DB
	config *config.Config
}

// NewAuthService is constructor function, and we need it to connect to another file with authservice
func NewAuthService(db *gorm.DB, config *config.Config) *AuthService {
	return &AuthService{
		db:     db,
		config: config,
	}
}

// Register function for register new user like a customer and create cart for this user.
func (s *AuthService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {

	// existingUser variable for check user if exists on db
	var existingUser models.User
	if err := s.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("user already exists")
	}

	// hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// create user
	user := models.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Role:      models.UserRoleCustomer,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	// create a cart
	cart := models.Cart{
		UserID: user.ID,
	}
	if err := s.db.Create(&cart).Error; err != nil {
		log.Println("Failed to create cart for user:", err)
	}

	return s.generateAuthResponse(&user)
	// generate token

}

// Login function for login user with email and password
func (s *AuthService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	var user models.User
	if err := s.db.Where("email = ? AND is_active = ?", req.Email, true).First(&user).Error; err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	return s.generateAuthResponse(&user)
}

// RefreshToken function for create new access token with refresh token, and we need it when access token is expired, and we can use refresh token to create new access token without login again.

func (s *AuthService) RefreshToken(req *dto.RefreshTokenRequest) (*dto.AuthResponse, error) {
	claims, err := utils.ValidateToken(req.RefreshToken, s.config.JWT.Secret)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	var refreshToken models.RefreshToken
	if err := s.db.Where("token = ? AND expires_at > ?", req.RefreshToken, time.Now()).First(&refreshToken).Error; err != nil {
		return nil, errors.New("refresh token not found or expired")
	}

	var user models.User
	if err := s.db.First(&user, claims.UserID).Error; err != nil {
		return nil, errors.New("user not found")
	}

	s.db.Delete(&refreshToken)

	return s.generateAuthResponse(&user)
}

// Logout function for logout user
func (s *AuthService) Logout(refreshToken string) error {
	return s.db.Where("token = ?", refreshToken).Delete(&models.RefreshToken{}).Error
}

func (s *AuthService) generateAuthResponse(user *models.User) (*dto.AuthResponse, error) {
	accessToken, refreshToken, err := utils.GenerateTokenPair(
		&s.config.JWT,
		user.ID,
		user.Email,
		string(user.Role),
	)

	if err != nil {
		return nil, err
	}

	refreshTokenModel := models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(s.config.JWT.RefreshTokenExpiresIn),
	}

	// Create the refresh token record and return error if fails
	if err := s.db.Create(&refreshTokenModel).Error; err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		User: dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Phone:     user.Phone,
			Role:      string(user.Role),
			IsActive:  user.IsActive,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
