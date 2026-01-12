package repository

import (
	"services/customer-service/internal/database"
	"services/customer-service/internal/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(customer model.Customer) (model.Customer, error)
	GetCustomerByID(id int) (model.Customer, error)
	UpdateCustomer(customer model.Customer) (model.Customer, error)
	DeleteCustomer(id int) error
}

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository() CustomerRepository {
	return &customerRepo{
		db: database.DB,
	}
}

func (r *customerRepo) CreateCustomer(customer model.Customer) (model.Customer, error) {
	if err := r.db.Create(&customer).Error; err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (r *customerRepo) GetCustomerByID(id int) (model.Customer, error) {
	var customer model.Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (r *customerRepo) UpdateCustomer(customer model.Customer) (model.Customer, error) {
	if err := r.db.Save(&customer).Error; err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (r *customerRepo) DeleteCustomer(id int) error {
	if err := r.db.Delete(&model.Customer{}, id).Error; err != nil {
		return err
	}
	return nil
}



