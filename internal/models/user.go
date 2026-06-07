package models

import (
	"time"

	"gorm.io/gorm"
)

// User struct for user. It includes user info such as ID, Email, Password, FirstName, LastName, Phone, IsActive, Role, CreatedAt, UpdatedAt, and DeletedAt. It also defines relationships to the RefreshTokens, Orders, and Cart. The Email field is unique and not null to ensure that each user can be uniquely identified. The Password field is not included in the JSON response for security reasons. The IsActive field indicates whether the user account is currently active. The Role field defines the user's role in the system, which can be either admin or customer.
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email" gorm:"uniqueIndex; not null"`
	Password  string         `json:"-" gorm:"not null"`
	FirstName string         `json:"first_name" gorm:"not null"`
	LastName  string         `json:"last_name" gorm:"not null"`
	Phone     string         `json:"phone"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	Role      UserRole       `json:"role" gorm:"default:customer"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationship
	RefreshTokens []RefreshToken `json:"-"`
	Orders        []Order        `json:"-"`
	Cart          Cart           `json:"-"`
}

// UserRole for role variable type of string
type UserRole string

// User roles admin and customer
var (
	UserRoleAdmin    UserRole = "admin"
	UserRoleCustomer UserRole = "customer"
)

// RefreshToken struct for user's refresh token. It also includes ID, UserID, Token, ExpiresAt, CreatedAt. It relationship with User
type RefreshToken struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	Token     string         `json:"token" gorm:"uniqueIndex; not null"`
	ExpiresAt time.Time      `json:"expires_at" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	User User `json:"-"`
}
