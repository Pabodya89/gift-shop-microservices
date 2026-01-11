package repository

import (
	"services/inventory-service/internal/database"
	"services/inventory-service/internal/model"
	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(item model.Item) (model.Item, error)
	GetItemByID(id int) (model.Item, error)
	UpdateItem(item model.Item) (model.Item, error)
	DeleteItem(id int) error
}
type itemRepo struct {
	db *gorm.DB
}
func NewItemRepository() ItemRepository {
	return &itemRepo{
		db: database.DB,
	}

}
func (r *itemRepo) CreateItem(item model.Item) (model.Item, error) {
	if err := r.db.Create(&item).Error; err != nil {
		return model.Item{}, err
	}
	return item, nil
}
func (r *itemRepo) GetItemByID(id int) (model.Item, error) {
	var item model.Item
	if err := r.db.First(&item, id).Error; err != nil {
		return model.Item{}, err
	}
	return item, nil
}
func (r *itemRepo) UpdateItem(item model.Item) (model.Item, error) {
	if err := r.db.Save(&item).Error; err != nil {
		return model.Item{}, err
	}
	return item, nil
}
func (r *itemRepo) DeleteItem(id int) error {
	if err := r.db.Delete(&model.Item{}, id).Error; err != nil {
		return err
	}
	return nil
}

