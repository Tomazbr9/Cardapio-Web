// internal/services/order_service.go
package services

import (
	"cw/internal/inputs"
	"cw/internal/models"
	"cw/internal/repositories"
	"errors"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderService interface {
	CreateOrder(tenantID uuid.UUID, input inputs.CreateOrderInput) (*models.Order, error)
}

type orderService struct {
	orderRepository   repositories.OrderRepository
	productRepository repositories.ProductRepository
	pizzaSizeRepository    repositories.PizzaSizeRepository   
	pizzaFlavorRepository  repositories.PizzaFlavorRepository 
}

func NewOrderService(
	orderRepository repositories.OrderRepository,
	productRepository repositories.ProductRepository,
	pizzaSizeRepository repositories.PizzaSizeRepository,
	pizzaFlavorRepository repositories.PizzaFlavorRepository,
) OrderService {
	return &orderService{
		orderRepository:   orderRepository,
		productRepository: productRepository,
		pizzaSizeRepository:    pizzaSizeRepository,
		pizzaFlavorRepository:  pizzaFlavorRepository,
	}
}

func (s *orderService) CreateOrder(tenantID uuid.UUID, input inputs.CreateOrderInput) (*models.Order, error) {
	// 1. Inicializa o Cabeçalho do Pedido
	order := &models.Order{
		TenantID:         tenantID,
		CustomerName:     input.CustomerName,
		CustomerWhatsapp: input.CustomerWhatsapp,
		DeliveryFee:      input.DeliveryFee,
		PaymentMethod:    input.PaymentMethod,
		AddressJSON:      input.AddressJSON,
		Status:           "PENDING",         // Todo pedido começa como pendente
		TotalAmount:      input.DeliveryFee, // O total começa valendo apenas a taxa de entrega
	}

	for _, itemInput := range input.Items {
		
		
		product, err := s.productRepository.FindById(itemInput.ProductID, tenantID)
		if err != nil {
			return nil, errors.New("produto não encontrado: " + itemInput.ProductID.String())
		}

		orderItem := models.OrderItem{
			ProductID:   product.ID,
			ProductName: product.Name,
			Quantity:    itemInput.Quantity,
			Notes:       itemInput.Notes,
		}

		var unitPrice decimal.Decimal

		if product.IsPizza {
			if itemInput.SizeID == nil || len(itemInput.FlavorIDs) == 0 {
				return nil, errors.New("Uma pizza exige tamanho e pelo menos um sabor")
			}

			orderItem.SizeID = itemInput.SizeID

			size, err := s.pizzaSizeRepository.FindById(*itemInput.SizeID, tenantID)

			if err != nil {
				return nil, errors.New("tamanho de pizza não encontrado")
			}

			var maxFlavorModifier decimal.Decimal = decimal.Zero

			for _, flavorID := range itemInput.FlavorIDs {
				
				flavor, err := s.pizzaFlavorRepository.FindById(flavorID, tenantID)
				
				if err != nil {
					return nil, errors.New("Sabor de pizza não encontrado")
				}

				if flavor.PriceModifier.GreaterThan(maxFlavorModifier) {
					maxFlavorModifier = flavor.PriceModifier
				}

				orderItem.Flavors = append(orderItem.Flavors, models.OrderItemFlavor{
					FlavorID:   flavor.ID,
					FlavorName: flavor.Name,
				})
			}

			unitPrice = size.BasePrice.Add(maxFlavorModifier)

		} else {
			unitPrice = product.BasePrice
		}

		orderItem.UnitPrice = unitPrice
		orderItem.Subtotal = unitPrice.Mul(decimal.NewFromInt(int64(orderItem.Quantity)))

		order.Items = append(order.Items, orderItem)
		order.TotalAmount = order.TotalAmount.Add(orderItem.Subtotal)
	}

	if err := s.orderRepository.CreateOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}