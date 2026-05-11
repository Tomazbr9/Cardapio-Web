package services

import (
	"cw/internal/inputs"
	"cw/internal/models"
	"cw/internal/repositories"

	"github.com/google/uuid"
)


type ProductService interface {
	CreateProduct(tenantId uuid.UUID, input inputs.CreateProductInput) (*models.Products, error)
	ListProductsForCategory(tenantId uuid.UUID, categoryId uuid.UUID) ([]models.Products, error)
}

type productService struct {
	repository repositories.ProductRepository
}

func NewProductService(repository repositories.ProductRepository) ProductService {
	return &productService{repository: repository}
}

func (service *productService) CreateProduct(tenantId uuid.UUID, input inputs.CreateProductInput) (*models.Products, error) {

	product := &models.Products{
		TenantID: tenantId,
		CategoryID: input.CategoryID,
		Name: input.Name,
		Description: input.Description,
		BasePrice: input.BasePrice,
		ImageUrl: input.ImageUrl,
		IsPizza: input.IsPizza,
		IsActive: true,
	}

	if err := service.repository.CreateProduct(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (service productService) ListProductsForCategory(tenantId uuid.UUID, categoryId uuid.UUID) ([]models.Products, error) {
	return service.repository.FindByTenantAndCategory(tenantId, categoryId)
}

