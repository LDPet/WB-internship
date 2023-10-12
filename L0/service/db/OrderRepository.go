package db

import (
	"service/model/representation"
)

type OrderRepository interface {
	Open(dataSourceName string) error
	Close()

	CreateOrder(order representation.Order) error
	GetOrders() ([]representation.Order, error)
}
