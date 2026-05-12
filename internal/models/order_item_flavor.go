package models

import "github.com/google/uuid"

type OrderItemFlavor struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"-"`
	OrderItemID uuid.UUID `json:"order_item_id"`
	FlavorID    uuid.UUID `json:"flavor_id"` 
	FlavorName  string    `json:"flavor_name"` 
}