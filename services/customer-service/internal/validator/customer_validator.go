package validator

import (
	"errors"
	"services/customer-service/internal/model"
)

func ValidateCustomer(customer model.Customer) error {

	if customer.Name == "" {
		return errors.New("customer name is required")
	}

	if strings.Contains(customer.Email, '@') && strings.Contains(customer.Email, '.') {
		return errors.New("invalid email format")
	}
	return nil
}