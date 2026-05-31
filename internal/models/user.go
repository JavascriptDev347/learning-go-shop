package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex; not null"`
	Password  string    `json:"-" gorm:"not null"`
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  string    `json:"last_name" gorm:"not null"`
	Phone     string    `json:"phone"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	Role      UserRole  `json:"role" gorm:"default:customer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-" gorm:"index"`

	//relationship
	RefreshTokens []RefreshToken `json:"-"`
	Orders        []Order        `json:"-"`
	Cart          Cart           `json:"-"`
}

type UserRole string

var (
	UserRoleAdmin    UserRole = "admin"
	UserRoleCustomer UserRole = "customer"
)

type RefreshToken struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Token     string    `json:"token" gorm:"uniqueIndex; not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"-" gorm:"index"`

	User User `json:"-"`
}
