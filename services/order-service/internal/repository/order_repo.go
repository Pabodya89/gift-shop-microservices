package repository

import (
	"services/order-service/internal/database"
	"services/order-service/internal/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order model.Order) (model.Order, error)
	GetOrderByID(id int) (model.Order, error)
	UpdateOrder(order model.Order) (model.Order, error)
	DeleteOrder(id int) error
}

type orderRepo struct {
	db *gorm.DB
} 

func NewOrderRepository() OrderRepository {
	return &orderRepo{
		db: database.DB,
	}
}

func (r *orderRepo) CreateOrder(order model.Order) (model.Order, error) {
	if err := r.db.Create(&order).Error; err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (r *orderRepo) GetOrderByID(id int) (model.Order, error) {
	var order model.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (r *orderRepo) UpdateOrder(order model.Order) (model.Order, error) {
	if err := r.db.Save(&order).Error; err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (r *orderRepo) DeleteOrder(id int) error {
	if err := r.db.Delete(&model.Order{}, id).Error; err != nil {
		return err
	}
	return nil
}

