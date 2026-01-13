type InventoryCheckCommand struct {
	OrderID string `json:"order_id"`
	Items []OrderItem `json:"items"`
}