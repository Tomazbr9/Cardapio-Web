package models

import "github.com/google/uuid"

type PizzaSizes struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	TenantID uuid.UUID `json:"tenant_id"`
	Name string `json:"name"`
	Slices int `json:"slices"`
	MaxFlavors int `json:"max_flavors"` 
}