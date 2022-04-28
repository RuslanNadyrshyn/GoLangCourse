package repositories

import "Delivery/Delivery/internal/repositories/models"

type UserRepositoryInterface interface {
	GetByEmail(email string) models.User
	Insert(user *models.User) error
	Delete(user *models.User) error
}

type SupplierRepositoryInterface interface {
	GetAll() (suppliers []models.Supplier, err error)
	Insert(id int, name string) (supplier models.Supplier, err error)
	Delete(supplier *models.Supplier) error
}

type ProductRepositoryInterface interface {
	GetAll() (products []models.Product, err error)
	Add(id int, price float32) (product models.Product, err error)
	Delete(product *models.Product) error
}

type OrderRepositoryInterface interface {
	GetAll(orderId int) (order []*models.Order, err error)
	Insert(order *models.Order) (id int, err error)
}

//type OrderProductsRepositoryInterface interface {
//	GetAll() (orderList models2.OrderProducts, err error)
//	GetByOrderId(orderProductsId int) (orderProducts models2.OrderProducts, err error)
//	GetByUserId(userId int) (orderProducts []*models2.OrderProducts, err error)
//}
