package services

import (
	"cw/inputs"
	"cw/internal/models"
	"cw/internal/repositories"

	"github.com/google/uuid"
)

type TenantService struct {
	repository repositories.TenantRepository
}

func NewTenantService(repository repositories.TenantRepository)  *TenantService {
	return &TenantService{repository: repository}
}

func (service *TenantService) CreateTenant(input inputs.CreateTenantInput) (models.Tenant, error) {
	
	tenant := models.Tenant{
		Name: input.Name,
		Slug: input.Slug,
		WhatsappNumber: input.WhatsappNumber,
		ConfigJson: input.ConfigJSON,
	}

	err := service.repository.CreateTenant(&tenant)

	return tenant, err
}

func (service *TenantService) GetTenant(tenantId string) (*models.Tenant, error) {
	
	id, err := uuid.Parse(tenantId)

	if err != nil {
		return nil, err
	}
	
	return service.repository.FindByID(id)
}

func (service *TenantService) UpdateTenant(tenantId string, input inputs.UpdateTenantInput) (*models.Tenant, error) {

	tenant, err := service.GetTenant(tenantId)

	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		tenant.Name = *input.Name
	}

	if input.Slug != nil {
		tenant.Slug = *input.Slug
	}

	if input.WhatsappNumber != nil {
		tenant.WhatsappNumber = *input.WhatsappNumber
	}

	if input.ConfigJSON != nil {
		tenant.ConfigJson = *input.ConfigJSON
	}

	if err := service.repository.UpdateTenant(tenant); err != nil {
		return nil, err
	}

	return tenant, nil

}


