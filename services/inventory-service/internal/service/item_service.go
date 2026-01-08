package service

import (
	"inventory-service/services/inventory-service/internal/model"
	"inventory-service/services/inventory-service/internal/repository"
)

func GetAllItems() []model.Item {
	return repository.GetAllItems()
}

func GetItemByID(id int) *model.Item {
	return repository.GetItemByID(id)
}

func CreateItem(item model.Item) model.Item {
	return repository.CreateItem(item)
}

func UpdateItem(id int ,item model.Item) *model.Item {
	return repository.UpdateItem(id, item)
}

func DeleteItem(id int) bool {
	return repository.DeleteItem(id)
}