package handlers

import (
	"cw/internal/inputs"
	"cw/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PizzaFlavorHandler struct {
	service services.PizzaFlavorService
}

func NewPizzaFlavorHandler(service services.PizzaFlavorService) *PizzaFlavorHandler {
	return &PizzaFlavorHandler{service: service}
}

func (handler *PizzaFlavorHandler) CreatePizzaFlavor(c *gin.Context) {

	tenantID := c.MustGet("tenantID").(uuid.UUID)

	var input inputs.CreatePizzaFlavorInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),})
		return 
	}

	pizzaFlavor, err := handler.service.CreatePizzaFlavor(tenantID, input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": pizzaFlavor,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": pizzaFlavor})

}

func (handler *PizzaFlavorHandler) ListFlavors(c *gin.Context) {
	tenantID := c.MustGet("tenantID").(uuid.UUID)

	pizzaFlavors, err := handler.service.ListFlavors(tenantID)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar sabores",
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": pizzaFlavors})
}

func (handler *PizzaSizeHandler) GetPizzaSize(c *gin.Context) {

	pizzaFlavorIdStr := c.Param("productId")
	tenantId := c.MustGet("id").(uuid.UUID)

	pizzaFlavorId, err := uuid.Parse(pizzaFlavorIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do produto inválido",
		})
		return 
	}

	pizzaFlavor, err := handler.service.GetPizzaSize(pizzaFlavorId, tenantId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Produto não encontrado",
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": pizzaFlavor})
}