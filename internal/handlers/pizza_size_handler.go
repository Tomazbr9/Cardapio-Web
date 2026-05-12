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
	tenantID := c.MustGet("tenantID").(uuid.UUID)

	pizzaSizes, err := handler.service.ListPizzaSizes(tenantID)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar tamanhos",
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": pizzaSizes})
}