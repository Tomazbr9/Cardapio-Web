package inputs

import "github.com/google/uuid"

type OrderItemInput struct {
	ProductID uuid.UUID   `json:"product_id" binding:"required"`
	Quantity  int         `json:"quantity" binding:"required,min=1"`
	Notes     string      `json:"notes"`
	SizeID    *uuid.UUID  `json:"size_id"`
	FlavorIDs []uuid.UUID `json:"flavor_ids"`
}