package server

import (
	"strconv"

	"github.com/JavascriptDev347/learning-go-shop/internal/dto"
	"github.com/JavascriptDev347/learning-go-shop/internal/services"
	"github.com/JavascriptDev347/learning-go-shop/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) createCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
		return
	}

	productService := services.NewProductService(s.db)
	category, err := productService.CreateCategory(req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to create category", err)
		return
	}
	utils.SuccessResponse(c, "Category created successfully", category)
}

func (s *Server) getCategories(c *gin.Context) {
	productService := services.NewProductService(s.db)
	categories, err := productService.GetCategories()
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to get categories", err)
		return
	}
	utils.SuccessResponse(c, "Category retrieved successfully", categories)
}

func (s *Server) updateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid id", err)
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
		return
	}

	productService := services.NewProductService(s.db)
	category, err := productService.UpdateCategory(uint(id), &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to update category", err)
		return
	}

	utils.SuccessResponse(c, "Category updated successfully", category)
}

func (s *Server) deleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid id", err)
		return
	}

	productService := services.NewProductService(s.db)
	err = productService.DeleteCategory(uint(id))
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to delete category", err)
		return
	}
	utils.SuccessResponse(c, "Category deleted successfully", nil)
}

func (s *Server) createProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
		return
	}

	productService := services.NewProductService(s.db)
	product, err := productService.CreateProduct(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to create product", err)
		return
	}
	utils.SuccessResponse(c, "Product created successfully", product)

}

func (s *Server) getProducts(c *gin.Context) {
	productService := services.NewProductService(s.db)
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	products, meta, err := productService.GetProducts(limit, page)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to get products", err)
		return
	}
	utils.PaginatedSuccessResponse(c, "Products retrieved successfully", products, *meta)

}

func (s *Server) getProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid id", err)
		return
	}
	productService := services.NewProductService(s.db)
	product, err := productService.GetProduct(uint(id))
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to get product", err)
		return
	}
	utils.SuccessResponse(c, "Product retrieved successfully", product)
}

func (s *Server) updateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid id", err)
		return
	}
	var req dto.UpdateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
	}

	productService := services.NewProductService(s.db)
	product, err := productService.UpdateProduct(uint(id), &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to update category", err)
		return
	}
	utils.SuccessResponse(c, "Product updated successfully", product)
}

func (s *Server) deleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid id", err)
		return
	}
	productService := services.NewProductService(s.db)
	err = productService.DeleteProduct(uint(id))
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to delete category", err)
		return
	}
	utils.SuccessResponse(c, "Category deleted successfully", nil)
}
