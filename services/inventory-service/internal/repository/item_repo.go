package repository

import "inventory-service/services/inventory-service/internal/model"

var items []model.Item

var lastID int = 0

func GetAllItems() []model.Item {
	return items
}

func GetItemByID(id int) *model.Item {
	for _,item := range items {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func CreateItem(item model.Item) model.Item {
	lastID++
	item.ID = lastID
	items = append(items,item)
	return item
}

func UpdateItem(id int, updated model.Item) *model.Item{
	for i,item := range items {
		if item.ID == id {
			items[i].ID = updated.ID
			items[i].Name = updated.Name
			items[i].Price = updated.Price
			items[i].Stock = updated.Stock
			return &items[i]
		}
	}
	return nil
}

func DeleteItem(id int) bool {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[:i+1]...)
			return true
		}
	}
	return false
}