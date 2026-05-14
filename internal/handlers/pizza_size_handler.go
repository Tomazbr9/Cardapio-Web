package handlers

import (
	"cw/internal/inputs"
	"cw/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type PizzaSizeHandler struct {
	service services.PizzaSizeService
}

func NewPizzaSizeHandler(service services.PizzaSizeService) *PizzaSizeHandler {
	return &PizzaSizeHandler{service: service}
} 

func (handler *PizzaSizeHandler) CreatePizzaSize(c *gin.Context) {

	tenantId := c.MustGet("tenantID").(uuid.UUID)

	var input inputs.CreatePizzaSizeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),})
		return 
	}

	pizzaSize, err := handler.service.CreatePizzaSize(tenantId, input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao criar tamanho",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": pizzaSize})
} 

func (handler *PizzaSizeHandler) ListPizzaSizes(c *gin.Context) {
	tenantId := c.MustGet("tenantID").(uuid.UUID)

	pizzaSizes, err := handler.service.ListPizzaSizes(tenantId)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar tamanhos",
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": pizzaSizes})
}

func (handler *PizzaSizeHandler) GetPizzaFlavor(c *gin.Context) {

	pizzaSizeIdStr := c.Param("productId")
	tenantId := c.MustGet("id").(uuid.UUID)

	pizzaSizeId, err := uuid.Parse(pizzaSizeIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do produto inválido",
		})
		return 
	}

	pizzaSize, err := handler.service.GetPizzaSize(pizzaSizeId, tenantId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Produto não encontrado",
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": pizzaSize})
}