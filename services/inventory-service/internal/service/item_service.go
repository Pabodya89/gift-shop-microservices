package services

import (
	"services/inventory-service/internal/model"
	"services/inventory-service/internal/repository"
	"services/inventory-service/internal/validator"
)
type ItemService interface {
	CreateItem(item model.Item) (model.Item, error)
	GetItemByID(id int) (model.Item, error)
	UpdateItem(item model.Item) (model.Item, error)
	DeleteItem(id int) error
}
type itemService struct {
	repo repository.ItemRepository
}
func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{
		repo: repo,
	}
}

func (s *itemService) CreateItem(item model.Item) (model.Item, error) {
	if err := validator.ValidateItem(item); err != nil {
		return model.Item{}, err
	}
	return s.repo.CreateItem(item)
}
func (s *itemService) GetItemByID(id int) (model.Item, error) {
	return s.repo.GetItemByID(id)
}

func (s *itemService) UpdateItem(item model.Item) (model.Item, error) {
	if err := validator.ValidateItem(item); err != nil {
		return model.Item{}, err
	}
	return s.repo.UpdateItem(item)
}

func (s *itemService) DeleteItem(id int) error {
	return s.repo.DeleteItem(id)
}

