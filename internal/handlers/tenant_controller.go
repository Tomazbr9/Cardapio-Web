package handlers

import (
	"cw/internal/inputs"
	"cw/internal/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type TenantHandler struct {
	service services.TenantService
}

func NewTenantHandler(service services.TenantService) *TenantHandler {
	return &TenantHandler{service: service}
}

func (handler *TenantHandler) CreateTenant(c *gin.Context) {

	var input inputs.CreateTenantInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return 
	}

	tenant, err := handler.service.CreateTenant(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao criar cliente",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": tenant,
	})
}

func (handler *TenantHandler) GetTenant(c *gin.Context) {
	
	id := c.Param("id")

	tenant, err := handler.service.GetTenant(id)
	if err != nil {
		
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Pizzaria não encontrada"})
			return
		}
		
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido ou erro interno"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tenant})
}

func (handler *TenantHandler) UpdateTenant(c *gin.Context) {

	id := c.Param("id")

	var input inputs.UpdateTenantInput
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenant, err := handler.service.UpdateTenant(id, input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Pizzaria não encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tenant})
}