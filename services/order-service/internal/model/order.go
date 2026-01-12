package model

package model

type OrderItem struct {
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

type Order struct {
	ID         int         `json:"id"`
	CustomerID int         `json:"customer_id"`
	Items      []OrderItem `json:"items"`
	Total      int         `json:"total"`
	Status     string      `json:"status"` 
}

