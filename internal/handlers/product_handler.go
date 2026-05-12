package handlers

import (
	"cw/internal/inputs"
	"cw/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (handler *ProductHandler) CreateProduct(c *gin.Context) {

	tenantID := c.MustGet("tenantID").(uuid.UUID)

	var input inputs.CreateProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),})
		return 
	}

	product, err := handler.service.CreateProduct(tenantID, input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao criar produto",
		})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{"data": product})

}

func (handler *ProductHandler) ListProductsForCategory(c *gin.Context) {

	tenantId := c.MustGet("tenantID").(uuid.UUID)
	categoryId := c.Param("category_id")

	categoryIdStr, err := uuid.Parse(categoryId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da categoria inválido",
		})
	}

	products, err := handler.service.ListProductsForCategory(tenantId, categoryIdStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar produtos da categoria",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})

}

func (handler *ProductHandler) GetProduct(c *gin.Context) {
	
	productIdStr := c.Param("id")
	tenantId := c.MustGet("tenantID").(uuid.UUID)

	productId, err := uuid.Parse(productIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do produto inválido",
		})
		return 
	}

	product, err := handler.service.GetProduct(productId, tenantId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Produto não encontrado",
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": product})

}