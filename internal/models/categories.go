package models

import (
	"time"

	"github.com/google/uuid"
)

type Categories struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	TenantID uuid.UUID `json:"tenant_id"`
	Name string `json:"name"`
	Position int `json:"position"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
} 