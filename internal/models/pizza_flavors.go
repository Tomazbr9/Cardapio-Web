package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PizzaFlavors struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	TenantID uuid.UUID `json:"tenant_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	PriceModifier decimal.Decimal `json:"price_modifier"`
}

