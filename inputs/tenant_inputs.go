package inputs

import "encoding/json"

type CreateTenantInput struct {
	Name           string          `json:"name" binding:"required"`
	Slug           string          `json:"slug" binding:"required"`
	WhatsappNumber string          `json:"whatsapp_number"`
	ConfigJSON     json.RawMessage `json:"config_json"` 
}

type UpdateTenantInput struct {
	Name           *string          `json:"name"`
	Slug           *string          `json:"slug"`
	WhatsappNumber *string          `json:"whatsapp_number"`
	ConfigJSON     *json.RawMessage `json:"config_json"`
}