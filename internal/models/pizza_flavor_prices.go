package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PizzaFlavorPrices struct {
	ID uuid.UUID `json:"id"`
	FlavorID uuid.UUID `json:"flavor_id"`
	SizeID uuid.UUID `json:"size_id"`
	Price decimal.Decimal `json:"price"`
}