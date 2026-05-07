package models

import (

	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Products struct {
	ID uuid.UUID `json:"id"`
	TenantID uuid.UUID `json:"tenant_id"`
	CategoryID uuid.UUID `json:"category_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	BasePrice decimal.Decimal `json:"base_price"`
	ImageUrl string `json:"image_url"`
	IsPizza bool `json:"is_pizza"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}