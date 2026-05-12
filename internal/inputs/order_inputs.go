package inputs

import "encoding/json"

type CreateOrderInput struct {
	CustomerName     string          `json:"customer_name" binding:"required"`
	CustomerWhatsapp string          `json:"customer_whatsapp" binding:"required"`
	PaymentMethod    string          `json:"payment_method"`
	AddressJSON      json.RawMessage `json:"address_json"`
	DeliveryFee      float64         `json:"delivery_fee"`
	Items            []OrderItemInput `json:"items" binding:"required,min=1"`
}