package handlers

import (
	"cw/internal/inputs"
	"cw/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	service services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (handler *CategoryHandler) CreateCategory(c *gin.Context) {
	tenantID := c.MustGet("tenantID").(uuid.UUID)

	var input inputs.CreateCategoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),})
		return 
	}

	category, err := handler.service.CreateCategory(tenantID, input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao criar categoria",
		})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{"data": category})

}

func (handler *CategoryHandler) ListTenantCategories(c *gin.Context) {

	tenantID := c.MustGet("tenantID").(uuid.UUID)

	categories, err := handler.service.ListTenantCategories(tenantID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar categorias",
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}