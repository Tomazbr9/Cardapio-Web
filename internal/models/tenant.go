package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	WhatsappNumber string `json:"whatsapp_number"`
	ConfigJson json.RawMessage `json:"config_json"`
	CreatedAt time.Time `json:"created_at"`
}