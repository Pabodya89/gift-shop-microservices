package validator

import (
	"errors"
	"services/order-service/internal/model"
)

func ValidateOrder(order model.Order) error {
	if order.CustomerID <= 0 {
		return errors.New("invalid customer ID")
	}
	if len(order.Items) == 0 {
		return errors.New("order must contain at least one item")
	}
	for _, item := range order.Items {
		if item.ItemID <= 0 {
			return errors.New("invalid item ID")
		}
		if item.Quantity <= 0 {
			return errors.New("item quantity must be greater than zero")
		}
	}
	if order.Total < 0 {
		return errors.New("total amount cannot be negative")
	}
	validStatuses := map[string]bool{
		"pending":   true,
		"shipped":   true,
		"delivered": true,
		"canceled":  true,
	}
	if !validStatuses[order.Status] {
		return errors.New("invalid order status")
	}
	return nil
}