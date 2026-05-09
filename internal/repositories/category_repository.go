package repositories

import (
	"cw/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type CategoryRepository interface {
	CreateCategory(category *models.Categories) error
	FindAllCategoriesByTenant(tenantId uuid.UUID) ([]models.Categories, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (repository *categoryRepository) CreateCategory(category *models.Categories) error {
	return repository.db.Create(category).Error
}

func (repository *categoryRepository) FindAllCategoriesByTenant(tenantId uuid.UUID) ([]models.Categories, error){
	var categories []models.Categories
	
	err := repository.db.Where(
		"tenant_id = ? AND is_active = ?", tenantId, true).
		Order("position ASC").
		Find(&categories).Error

	return categories, err
}