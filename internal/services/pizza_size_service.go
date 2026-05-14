package services

import (
	"cw/internal/inputs"
	"cw/internal/models"
	"cw/internal/repositories"

	"github.com/google/uuid"
)


type PizzaSizeService interface {
	CreatePizzaSize(tenantId uuid.UUID, input inputs.CreatePizzaSizeInput) (*models.PizzaSizes, error)
	ListPizzaSizes(tenantId uuid.UUID) ([]models.PizzaSizes, error)
	GetPizzaSize(pizzaSizeId uuid.UUID, tenantId uuid.UUID) (*models.PizzaSizes, error)
}

type pizzaSizeService struct {
	repository repositories.PizzaSizeRepository
}

func NewPizzaSizeService(repository repositories.PizzaSizeRepository) PizzaSizeService {
	return &pizzaSizeService{repository: repository}
}

func (service *pizzaSizeService) CreatePizzaSize(tenantId uuid.UUID, input inputs.CreatePizzaSizeInput) (*models.PizzaSizes, error) {
	
	pizzaSize := &models.PizzaSizes{
		TenantID: tenantId,
		Name: input.Name,
		Slices: input.Slices,
		MaxFlavors: input.MaxFlavors,
		BasePrice: input.BasePrice,
      }

	if err := service.repository.CreatePizzaSize(pizzaSize); err != nil {
		return nil, err
	}

	return pizzaSize, nil 
}

func (service *pizzaSizeService) ListPizzaSizes(tenantId uuid.UUID) ([]models.PizzaSizes, error) {
	return service.repository.FindAllPizzaSizeByTenant(tenantId)
}

func (service *pizzaSizeService) GetPizzaSize(pizzaSizeId uuid.UUID, tenantId uuid.UUID) (*models.PizzaSizes, error) {
	
	pizzaSize, err := service.repository.FindById(pizzaSizeId, tenantId)

	if err != nil {
		return nil, err
	}

	return pizzaSize, nil
}