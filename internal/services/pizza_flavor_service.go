package services

import (
	"cw/internal/inputs"
	"cw/internal/models"
	"cw/internal/repositories"

	"github.com/google/uuid"
)


type PizzaFlavorService interface {
	CreatePizzaFlavor(tenantId uuid.UUID, input inputs.CreatePizzaFlavorInput) (*models.PizzaFlavors, error)
	ListFlavors(tenantId uuid.UUID) ([]models.PizzaFlavors, error)
	GetPizzaFlavor(pizzaFlavorId uuid.UUID, tenantId uuid.UUID) (*models.PizzaFlavors, error)
}

type pizzaFlavorService struct {
	repository repositories.PizzaFlavorRepository
}

func NewPizzaFlavorService(repository repositories.PizzaFlavorRepository) PizzaFlavorService {
	return  &pizzaFlavorService{repository: repository}
}

func (service  pizzaFlavorService) CreatePizzaFlavor(tenantID uuid.UUID, input inputs.CreatePizzaFlavorInput) (*models.PizzaFlavors, error){
	
	pizzaFlavor := &models.PizzaFlavors{
		TenantID: tenantID,
		Name: input.Name,
		Description: input.Description,
		PriceModifier: input.PriceModifier,
	}
      
	if err := service.repository.CreatePizzaFlavor(pizzaFlavor); err != nil {
		return nil, err
	}

	return pizzaFlavor, nil
}

func (service pizzaFlavorService) ListFlavors(tenantID uuid.UUID) ([]models.PizzaFlavors, error) {
	return service.repository.FindAllPizzaFlavorByTenant(tenantID)
}

func (service pizzaFlavorService) GetPizzaFlavor(pizzaFlavorId uuid.UUID, tenantId uuid.UUID) (*models.PizzaFlavors, error) {
	
	pizzaFlavor, err := service.repository.FindById(pizzaFlavorId, tenantId)

	if err != nil {
		return nil, err
	}

	return pizzaFlavor, nil
}