// Package models defines the data models for orders and carts in the e-commerce application. It includes the Order, OrderItem, Cart, and CartItem structs, along with their relationships and relevant fields. The OrderStatus type is also defined to represent the various states an order can be in.
package models

import (
	"time"

	"gorm.io/gorm"
)

// Order struct includes order details such as ID, UserID, Status, TotalAmount, CreatedAt, UpdatedAt, and DeletedAt. It also defines relationships to the User and OrderItems. The OrderStatus type is used to represent the status of an order, with possible values including pending, confirmed, shipped, delivered, and cancelled.
type Order struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id" gorm:"not null"`
	Status      OrderStatus    `json:"status" gorm:"default:pending"`
	TotalAmount float64        `json:"total_amount" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// relationship
	User       User        `json:"user"`
	OrderItems []OrderItem `json:"order_items"`
}

// OrderStatus for order status, includes pending,confirmed,shipped,canceled and delivered
type OrderStatus string

// const OrderStatus values
const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusConfirmed OrderStatus = "confirmed"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

// OrderItem struct includes order item details such as ID, OrderID, ProductID, Quantity, Price, CreatedAt, and DeletedAt. It also defines relationships to the Order and Product.
type OrderItem struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	OrderID   uint           `json:"order_id" gorm:"not null"`
	ProductID uint           `json:"product_id" gorm:"not null"`
	Quantity  int            `json:"quantity" gorm:"not null"`
	Price     float64        `json:"price" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// relationship
	Order   Order   `json:"-"`
	Product Product `json:"product"`
}

// Cart struct includes details such as ID, UserID, CreatedAt, UpdatedAt, DeletedAt. It also defines relationship to the CartItems.
type Cart struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// relationship
	CartItems []CartItem `json:"cart_items"`
}

// CartItem struct includes cartitem details such as ID, CartID, ProductID, Quantity, Price, CreatedAt, and DeletedAt. It also defines relationships to the Cart and Product.
type CartItem struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CartID    uint           `json:"cart_id" gorm:"not null"`
	ProductID uint           `json:"product_id" gorm:"not null"`
	Quantity  int            `json:"quantity" gorm:"not null"`
	Price     float64        `json:"price" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// relationship
	Cart    Cart    `json:"-"`
	Product Product `json:"product"`
}
