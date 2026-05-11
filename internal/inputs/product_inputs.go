package inputs

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)


type CreateProductInput struct {
	CategoryID uuid.UUID `json:"category_id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Description string `json:"description"`
	BasePrice decimal.Decimal `json:"base_price" binding:"required"`
	ImageUrl string `json:"image_url"`
	IsPizza bool `json:"is_pizza"`
}