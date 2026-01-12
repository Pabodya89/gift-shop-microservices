package services

import (
	"services/customer-service/internal/model"
	"services/customer-service/internal/repository"
	"services/customer-service/internal/validator"
)

type CustomerService interface {
	CreateCustomer(customer model.Customer) (model.Customer, error)
	GetCustomerByID(id int) (model.Customer, error)
	UpdateCustomer(customer model.Customer) (model.Customer, error)
	DeleteCustomer(id int) error
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{
		repo: repo,
	}
}

func (s *customerService) CreateCustomer(customer model.Customer) (model.Customer, error) {
	if err := validator.ValidateCustomer(customer); err != nil {
		return model.Customer{}, err
	}
	return s.repo.CreateCustomer(customer)
}

func (s *customerService) GetCustomerByID(id int) (model.Customer, error) {
	return s.repo.GetCustomerByID(id)
}

func (s *customerService) UpdateCustomer(customer model.Customer) (model.Customer, error) {
	if err := validator.ValidateCustomer(customer); err != nil {
		return model.Customer{}, err
	}
	return s.repo.UpdateCustomer(customer)
}

func (s *customerService) DeleteCustomer(id int) error {
	return s.repo.DeleteCustomer(id)
}

