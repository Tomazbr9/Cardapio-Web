package repositories

import (
	"cw/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PizzaFlavorRepository interface {
	CreatePizzaFlavor(pizzaFlavor *models.PizzaFlavors) error
	FindAllPizzaFlavorByTenant(tenantId uuid.UUID) ([]models.PizzaFlavors, error)
	FindById(pizzaFlavorId uuid.UUID, tenantId uuid.UUID) (*models.PizzaFlavors, error)
}

type pizzaFlavorRepository struct {
	db *gorm.DB
}

func NewPizzaFlavorRepository(db *gorm.DB) PizzaFlavorRepository {
	return &pizzaFlavorRepository{db: db}
}

func (repository *pizzaFlavorRepository) CreatePizzaFlavor(pizzaFlavor *models.PizzaFlavors) error {
	return repository.db.Create(pizzaFlavor).Error
}

func (repository *pizzaFlavorRepository) FindAllPizzaFlavorByTenant(tenantId uuid.UUID) ([]models.PizzaFlavors, error) {
	var pizzaFlavors []models.PizzaFlavors

	err := repository.db.Where("tenant_id = ?", tenantId).First(&pizzaFlavors).Error

	return pizzaFlavors, err
}

func (repository *pizzaFlavorRepository) FindById(pizzaFlavorId uuid.UUID, tenantId uuid.UUID) (*models.PizzaFlavors, error) {
	var pizzaFlavor models.PizzaFlavors

	err := repository.db.Where("id = ? AND tenant_id = ?", pizzaFlavorId, tenantId).First(&pizzaFlavor).Error

	if err != nil {
		return nil, err
	}

	return &pizzaFlavor, nil
}