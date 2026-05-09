package repositories

import (
	"cw/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) *TenantRepository {
	return &TenantRepository{db: db}
}

func (repository *TenantRepository) CreateTenant(tenant *models.Tenant) error {
	return repository.db.Create(tenant).Error
}

func (repository *TenantRepository) FindByID(id uuid.UUID) (*models.Tenant, error) {
	var tenant models.Tenant
	
	err := repository.db.First(&tenant, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (repository *TenantRepository) FindBySlug(slug string) (*models.Tenant, error) {

	var tenant models.Tenant

	err := repository.db.First(&tenant, "slug = ?", slug).Error

	if err != nil {
		return nil, err
	}

	return &tenant, nil 
}

func (repository *TenantRepository) UpdateTenant(tenant *models.Tenant) (error) {
	return repository.db.Save(tenant).Error
}