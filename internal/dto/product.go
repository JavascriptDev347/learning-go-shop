// Package dto for product dto
package dto

// CreateCategoryRequest defines the input payload required to create a new category.
type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=32"`
	Description string `json:"description"`
}

// UpdateCategoryRequest defines the input payload required to update an existing category.
type UpdateCategoryRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=32"`
	Description string `json:"description"`
	IsActive    *bool  `json:"is_active"`
}

// CategoryResponse represents the serialized data schema sent back to clients for a category.
type CategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

// CreateProductRequest defines the input payload required to create a new product.
type CreateProductRequest struct {
	CategoryID  uint    `json:"category_id" binding:"required"`
	Name        string  `json:"name" binding:"required,min=2,max=32"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"min=0"`
	SKU         string  `json:"sku" binding:"required"`
}

// UpdateProductRequest defines the input payload required to modify an existing product's details.
type UpdateProductRequest struct {
	CategoryID  uint    `json:"category_id" binding:"required"`
	Name        string  `json:"name" binding:"required,min=2,max=32"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"min=0"`
	IsActive    *bool   `json:"is_active"`
}

// ProductResponse represents the serialized detailed product structure, including nested relations, sent back to clients.
type ProductResponse struct {
	ID          uint                    `json:"id"`
	CategoryID  uint                    `json:"category_id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Price       float64                 `json:"price"`
	Stock       int                     `json:"stock"`
	SKU         string                  `json:"sku"`
	IsActive    bool                    `json:"is_active"`
	Category    CategoryResponse        `json:"category"`
	Images      []ProductImagesResponse `json:"images"`
}

// ProductImagesResponse represents the serialized asset information for a single product image.
type ProductImagesResponse struct {
	ID        uint   `json:"id"`
	URL       string `json:"url"`
	AltText   string `json:"alt_text"`
	IsPrimary bool   `json:"is_primary"`
}
