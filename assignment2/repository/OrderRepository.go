package repository

import (
	"assignment2/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

// FindOrder implements IOrderRepository.
func (or *OrderRepository) FindOrder(uuid string) (model.Orders, error) {
	panic("unimplemented")
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (or *OrderRepository) CreateOrder(newOrder model.Orders) (model.Orders, error) {
	tx := or.db.Create(&newOrder)

	return newOrder, tx.Error
}

func (or *OrderRepository) GetOrder([]model.Orders) ([]model.Orders, error) {
	var orders = []model.Orders{}
	if err := or.db.Preload("Items").Find(&orders).Error; err != nil {

		return nil, err
	}

	return orders, nil
}

// func (or *OrderRepository) UpdateOrder(updateOrder model.Orders,uuid string)(model.Orders, error){
// 	tx:= or.db.Model(&model.Orders{}).Where("id = ?",uuid).Updates(&updateOrder)

// 	return updateOrder, tx.Error
// }

func (or *OrderRepository) UpdateOrder(updateOrder model.Orders, uuid string) (model.Orders, error) {
	var lastOrder model.Orders

	tx := or.db.Where("id = ?", uuid).First(&lastOrder)
	if tx.Error != nil {
		return lastOrder, tx.Error
	}

	if err := or.db.Model(&lastOrder).Updates(&updateOrder).Error; err != nil {
		return lastOrder, err
	}

	if err := or.db.Model(&lastOrder).Association("Items").Replace(updateOrder.Items); err != nil {
		return lastOrder, err
	}
	return lastOrder, nil

}

func (or *OrderRepository) DeleteOrder(uuid string)(model.Orders, error ){
	var deletedOrder model.Orders
	tx:= or.db.Where("id = ?",uuid).First(&deletedOrder)
	if tx.Error != nil {
		return deletedOrder, tx.Error
	}
	tx = or.db.Where("id = ?", uuid).Delete(&deletedOrder)
	if tx.Error != nil {
		return deletedOrder, tx.Error
	}

	return deletedOrder,tx.Error
}
