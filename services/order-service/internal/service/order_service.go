package service

import (
	"github.com/gin-gonic/gin"
	"services/order-service/internal/handlers"
	"services/order-service/internal/routes"
)

type OrderService interface {
	CreateOrder(order model.Order) (model.Order, error)
	GetOrderByID(id int) (model.Order, error)
	UpdateOrder(order model.Order) (model.Order, error)
	DeleteOrder(id int) error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{
		repo: repo,
	}
}

func (s *orderService) CreateOrder(order model.Order) (model.Order, error) {
	if err := validator.ValidateOrder(order); err != nil {
		return model.Order{}, err
	}

	order.Status = "PENDING"
	saveOrder, err := s.repo.CreateOrder(order)

	if err != nil {
		return model.Order{}, err 
	}

	err = s.rabbitProducer.PublishInventoryCheck(saveOrder)
	if err != nil {
		return model.Order{}, err 
	}

	s.kafkaProducer.PublishOrderCreated(saveOrder)

	return saveOrder, nil
}

func (s *orderService) GetOrderByID(id int) (model.Order, error) {
	return s.repo.GetOrderByID(id)
}

func (s *orderService) UpdateOrder(order model.Order) (model.Order, error) {
	if err := validator.ValidateOrder(order); err != nil {
		return model.Order{}, err
	}
	return s.repo.UpdateOrder(order)
}

func (s *orderService) DeleteOrder(id int) error {
	return s.repo.DeleteOrder(id)
}






