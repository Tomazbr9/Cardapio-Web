package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Order struct {
	ID               uuid.UUID       `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	TenantID         uuid.UUID       `json:"tenant_id"`
	CustomerName     string          `json:"customer_name"`
	CustomerWhatsapp string          `json:"customer_whatsapp"` 
	TotalAmount      decimal.Decimal `json:"total_amount"`      
	DeliveryFee      decimal.Decimal `json:"delivery_fee"`      
	Status           string          `json:"status"`            
	PaymentMethod    string          `json:"payment_method"`
	AddressJSON      json.RawMessage `gorm:"type:jsonb" json:"address_json"` 
	Items            []OrderItem     `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt        time.Time       `json:"created_at"`
}
