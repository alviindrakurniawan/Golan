package repository

import "assignment2/model"


type IOrderRepository interface {
	CreateOrder(newOrder model.Orders)(model.Orders, error)
	GetOrder( []model.Orders)([]model.Orders, error)
	UpdateOrder(updateOrder model.Orders,uuid string)(model.Orders, error)
	DeleteOrder(uuid string)(model.Orders, error)
}