package services

import (
	"github.com/JavascriptDev347/learning-go-shop/internal/dto"
	"github.com/JavascriptDev347/learning-go-shop/internal/models"
	"gorm.io/gorm"
)

// UserService struct
type UserService struct {
	db *gorm.DB
}

// NewUserService constructor for you can use this constructor for another file
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// GetProfile func belongs to UserService, and it gets the user profile by userID. It returns a UserResponse struct and an error if any.
func (s *UserService) GetProfile(userID uint) (*dto.UserResponse, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.LastName,
		LastName:  user.FirstName,
		Phone:     user.Phone,
		Role:      string(user.Role),
		IsActive:  user.IsActive,
	}, nil
}

// UpdateProfile func belongs to UserService, and it updates the user profile by userID and the UpdateProfileRequest struct. It returns a UserResponse struct and an error if any.
func (s *UserService) UpdateProfile(userID uint, req *dto.UpdateProfileRequest) (*dto.UserResponse, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Phone = req.Phone

	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return s.GetProfile(userID)
}
