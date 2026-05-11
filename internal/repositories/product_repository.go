package repositories

import (
	"cw/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type ProductRepository interface {
	CreateProduct(product *models.Products) error
	FindByTenantAndCategory(tenantId uuid.UUID, categoryId uuid.UUID) ([]models.Products, error)
} 


type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (repository *productRepository) CreateProduct(product *models.Products) error {
	return repository.db.Create(product).Error
}

func (repository *productRepository) FindByTenantAndCategory(tenantId uuid.UUID, categoryId uuid.UUID) ([]models.Products, error) {
	var products []models.Products

	err := repository.db.Where("tenant_id = ? AND category_id = ? AND is_active = ?", tenantId, categoryId, true).Find(&products).Error

	return products, err
	
}