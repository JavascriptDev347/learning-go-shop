package services

import (
	"github.com/JavascriptDev347/learning-go-shop/internal/dto"
	"github.com/JavascriptDev347/learning-go-shop/internal/models"
	"github.com/JavascriptDev347/learning-go-shop/internal/utils"
	"gorm.io/gorm"
)

// ProductService struct for product
type ProductService struct {
	db *gorm.DB
}

// NewProductService is constructor for ProductService and if this not available ypu can not use this ProductService inside another file
func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{
		db: db,
	}
}

// CreateCategory Category add
func (s *ProductService) CreateCategory(req dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	category := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.db.Create(&category).Error; err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
	}, nil
}

// GetCategories return all categories
func (s *ProductService) GetCategories() ([]dto.CategoryResponse, error) {
	var categories []models.Category
	if err := s.db.Where("is_active = ?", true).Find(&categories).Error; err != nil {
		return nil, err
	}

	response := make([]dto.CategoryResponse, len(categories))

	for i := range categories {
		response[i] = dto.CategoryResponse{
			ID:          categories[i].ID,
			Name:        categories[i].Name,
			Description: categories[i].Description,
			IsActive:    categories[i].IsActive,
		}
	}
	return response, nil
}

// UpdateCategory for update category
func (s *ProductService) UpdateCategory(id uint, req *dto.UpdateCategoryRequest) (*dto.ProductResponse, error) {

	var category models.Category
	if err := s.db.First(&category, id).Error; err != nil {
		return nil, err
	}

	category.Name = req.Name
	category.Description = req.Description
	if req.IsActive != nil {
		category.IsActive = *req.IsActive
	}

	if err := s.db.Save(&category).Error; err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
	}, nil
}

// DeleteCategory for  delete category by id and this will not delete the category, but it will set is_active to false
func (s *ProductService) DeleteCategory(id uint) error {
	return s.db.Delete(&models.Category{}, id).Error
}

// CreateProduct for create product
func (s *ProductService) CreateProduct(req *dto.CreateProductRequest) (*dto.ProductResponse, error) {

	product := models.Product{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		SKU:         req.SKU,
	}

	if err := s.db.Create(&product).Error; err != nil {
		return nil, err
	}

	return s.GetProduct(product.ID)

}

// GetProducts for get all products with pagination and this will return a list of products and pagination meta data such as total, page, limit and total pages
func (s *ProductService) GetProducts(limit, page int) ([]dto.ProductResponse, *utils.PaginationMeta, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var products []models.Product
	var total int64
	s.db.Model(&models.Product{}).Where("is_active = ? ", true).Count(&total)

	if err := s.db.Preload("Category").Preload("Images").
		Where("is_active = ? ", true).
		Offset(offset).Limit(limit).
		Find(&products).Error; err != nil {
		return nil, nil, err
	}

	response := make([]dto.ProductResponse, len(products))
	for i := range products {
		response[i] = s.convertToProductResponse(&products[i])
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	meta := &utils.PaginationMeta{
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}

	return response, meta, nil
}

// GetProduct for get one product
func (s *ProductService) GetProduct(id uint) (*dto.ProductResponse, error) {
	var product models.Product
	if err := s.db.Preload("Category").Preload("Images").
		First(&product, id).Error; err != nil {
		return nil, err
	}
	return new(s.convertToProductResponse(&product)), nil
}

// UpdateProduct for update product
func (s *ProductService) UpdateProduct(id uint, req *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	var product models.Product
	if err := s.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	product.CategoryID = req.CategoryID
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	if req.IsActive != nil {
		product.IsActive = *req.IsActive
	}
	if err := s.db.Save(&product).Error; err != nil {
		return nil, err
	}

	return s.GetProduct(product.ID)
}

// DeleteProduct for delete product
func (s *ProductService) DeleteProduct(id uint) error {
	return s.db.Delete(&models.Product{}, id).Error
}

// AddProductImage for add image to product
func (s *ProductService) AddProductImage(productID uint, url, altText string) error {

	var count int64

	s.db.Model(&models.ProductImage{}).Where("product_id = ?", productID).Count(&count)

	image := models.ProductImage{
		ProductID: productID,
		URL:       url,
		AltText:   altText,
		IsPrimary: count == 0, // Set as primary if it's the first image}
	}

	return s.db.Create(&image).Error
}

// converToProductResponse this is private function ant it helps return product easily
func (s *ProductService) convertToProductResponse(product *models.Product) dto.ProductResponse {

	images := make([]dto.ProductImagesResponse, len(product.Images))
	for i := range product.Images {
		images[i] = dto.ProductImagesResponse{
			ID:        product.Images[i].ID,
			URL:       product.Images[i].URL,
			AltText:   product.Images[i].AltText,
			IsPrimary: product.Images[i].IsPrimary,
		}
	}

	return dto.ProductResponse{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		SKU:         product.SKU,
		IsActive:    product.IsActive,
		Category: dto.CategoryResponse{
			ID:          product.Category.ID,
			Name:        product.Category.Name,
			Description: product.Category.Description,
			IsActive:    product.Category.IsActive,
		},
		Images: images,
	}
}
