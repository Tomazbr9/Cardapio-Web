package repositories

import (
	"cw/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type PizzaSizeRepository interface {
	CreatePizzaSize(pizzaSize *models.PizzaSizes) error
	FindAllPizzaSizeByTenant(tenantId uuid.UUID) ([]models.PizzaSizes, error)
	FindById(pizzaSizeId uuid.UUID, tenantId uuid.UUID) (*models.PizzaSizes, error)
}

type pizzaSizeRepository struct {
	db *gorm.DB
}

func NewPizzaSizeRepository(db *gorm.DB) PizzaSizeRepository {
	return &pizzaSizeRepository{db: db}
}

func (repository *pizzaSizeRepository) CreatePizzaSize(pizzaSize *models.PizzaSizes) error {
	return repository.db.Create(pizzaSize).Error
}

func (repository *pizzaSizeRepository) FindAllPizzaSizeByTenant(tenantId uuid.UUID) ([]models.PizzaSizes, error) {
	var pizzaSizes []models.PizzaSizes

	err := repository.db.Where("tenant_id = ?", tenantId).First(&pizzaSizes).Error

	return pizzaSizes, err
}

func (repository *pizzaSizeRepository) FindById(pizzaSizeId uuid.UUID, tenantId uuid.UUID) (*models.PizzaSizes, error) {
	var pizzaSize models.PizzaSizes

	err := repository.db.Where("id = ? AND tenant_id = ?", pizzaSizeId, tenantId).First(&pizzaSize).Error

	if err != nil {
		return nil, err
	}

	return &pizzaSize, nil
}