package repositories

import (
	"cw/internal/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) error 
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