package repositories

import (
	models2 "awesomeProject/internal/repositories/models"
)

type UserRepositoryInterface interface {
	GetByEmail(email string) models2.User
	Insert(user *models2.User) error
	Delete(user *models2.User) error
}

type SupplierRepositoryInterface interface {
	GetAll() (suppliers []*models2.Supplier, err error)
	Add(id int, name string) (supplier models2.Supplier, err error)
	Delete(supplier *models2.Supplier) error
}

type ProductRepositoryInterface interface {
	GetAll() (products []*models2.Product, err error)
	Add(id int, price float32) (product models2.Product, err error)
	Delete(product *models2.Product) error
}

type OrdersRepositoryInterface interface {
	GetAll(orderId int) (order []*models2.Orders, err error)
	Add(ProductId int) (orders models2.Orders, err error)
}

type OrderProductsRepositoryInterface interface {
	GetAll() (orderList models2.OrderProducts, err error)
	GetByOrderId(orderProductsId int) (orderProducts models2.OrderProducts, err error)
	GetByUserId(userId int) (orderProducts []*models2.OrderProducts, err error)
}
