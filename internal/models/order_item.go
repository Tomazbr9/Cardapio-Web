package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderItem struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	OrderID     uuid.UUID `json:"order_id"`
	ProductID   uuid.UUID `json:"product_id"`
	SizeID      *uuid.UUID `json:"size_id"` 
	ProductName string    `json:"product_name"` 
	UnitPrice   decimal.Decimal   `json:"unit_price"` 
	Quantity    int       `json:"quantity"`
	Subtotal    decimal.Decimal   `json:"subtotal"` 
	Notes       string    `json:"notes"`    
	Flavors     []OrderItemFlavor `gorm:"foreignKey:OrderItemID" json:"flavors,omitempty"`
}