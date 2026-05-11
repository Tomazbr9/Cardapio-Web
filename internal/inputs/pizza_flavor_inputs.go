package inputs

import "github.com/shopspring/decimal"

type CreatePizzaFlavorInput struct {
	Name string `json:"name" binding:"required"`
	Description string `json:"description"`
	PriceModifier decimal.Decimal `json:"price_modifier"`
}