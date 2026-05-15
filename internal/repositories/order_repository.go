package repositories

import (
	"cw/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) error 
	FindAllOrdersByTenant(tenantId uuid.UUID) ([]models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (repository *orderRepository) CreateOrder(order *models.Order) error {
	return repository.db.Create(order).Error
}

func (repository *orderRepository) FindAllOrdersByTenant(tenantId uuid.UUID) ([]models.Order, error) {
	
	var orders []models.Order

	err := repository.db.Preload("Items").
	Preload("Items.Flavors").
	Where("tenant_id = ?", tenantId).
	Order("created_at desc").
	Find(&orders).Error

	return orders, err
}