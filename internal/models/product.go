package models

import (
	"time"

	"gorm.io/gorm"
)

// Category struct includes details category for product such as ID, Name, Desctiption and others. It also defines a relationship to the Product. The IsActive field indicates whether the category is currently active and available for products.
type Category struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationship
	Products []Product `json:"-"`
}

// Product struct includes product details such as ID, CategoryID, Name, Description, Price, Stock, SKU, IsActive, CreatedAt, UpdatedAt, and DeletedAt. It also defines relationships to the Category, ProductImages, OrderItems, and CartItems. The SKU field is unique and not null to ensure that each product can be uniquely identified. The IsActive field indicates whether the product is currently active and available for purchase.
type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CategoryID  uint           `json:"category_id" gorm:"not null"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Price       float64        `json:"price" gorm:"not null"`
	Stock       int            `json:"stock" gorm:"default:0"`
	SKU         string         `json:"sku" gorm:"uniqueIndex;not null"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationship
	Category   Category       `json:"category"`
	Images     []ProductImage `json:"images"`
	OrderItems []OrderItem    `json:"-"`
	CartItems  []CartItem     `json:"-"`
}

// ProductImage struct includes details product image such as ID, ProductID, URL, AltText, IsPrimary, CreatedAt, and UpdatedAt. It also defines a relationship to the Product. The IsPrimary field indicates whether the image is the primary image for the product.
type ProductImage struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ProductID uint           `json:"product_id" gorm:"not null"`
	URL       string         `json:"url" gorm:"not null"`
	AltText   string         `json:"alt_text"`
	IsPrimary bool           `json:"is_primary" gorm:"default:false"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	// relationship
	Product Product `json:"-"`
}
