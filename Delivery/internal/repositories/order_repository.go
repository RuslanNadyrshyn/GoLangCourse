package repositories

import (
	"Delivery/Delivery/internal/repositories/models"
	"errors"
)

type IOrderRepository interface {
	GetAll() (order []*models.Order, err error)
	Add(order *models.Order) (int, error)
	GetById(orderId int) (*models.Order, error)
}

type OrderRepository struct {
	orders []*models.Order
}

func NewOrderRepository() IOrderRepository {
	orders := []*models.Order{
		{
			Id:      1,
			Price:   1.0,
			UserId:  1,
			Address: "dlksaj",
		},
	}

	return &OrderRepository{orders: orders}
}

func (o *OrderRepository) GetAll() ([]*models.Order, error) {
	return o.orders, nil
}

func (o *OrderRepository) Add(order *models.Order) (int, error) {
	o.orders = append(o.orders, order)
	return order.Id, nil
}

func (o *OrderRepository) GetById(orderId int) (*models.Order, error) {
	for _, order := range o.orders {
		if order.Id == orderId {
			return order, nil
		}
	}
	return nil, errors.New("order not found")
}
