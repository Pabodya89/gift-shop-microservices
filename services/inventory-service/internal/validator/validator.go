package validator

import (
	"errors"
	"services/inventory-service/internal/model"
)

func ValidateItem (item model.Item) error {
	if item.Quantitu <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	if item.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}