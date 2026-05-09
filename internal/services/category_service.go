package services

import (
	"cw/internal/inputs"
	"cw/internal/models"
	"cw/internal/repositories"

	"github.com/google/uuid"
)

type CategoryService interface {
	CreateCategory(tenantId uuid.UUID, request inputs.CreateCategoryInput) (*models.Categories, error)
	ListTenantCategories(tenantId uuid.UUID) ([]models.Categories, error)
}

type categoryService struct {
	repository repositories.CategoryRepository
}

func NewCategoryRepository(repository repositories.CategoryRepository) CategoryService{
	return &categoryService{repository: repository}
}

func (service *categoryService) CreateCategory(tenantID uuid.UUID, request inputs.CreateCategoryInput) (*models.Categories, error){
	category := &models.Categories{
		TenantID: tenantID,
		Name: request.Name,
		Position: request.Position,
		IsActive: true,
      }

	if err := service.repository.CreateCategory(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (service *categoryService) ListTenantCategories(tenantID uuid.UUID) ([]models.Categories, error) {
	return service.repository.FindAllCategoriesByTenant(tenantID)
}