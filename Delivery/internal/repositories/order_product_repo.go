package repositories

import (
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
	"log"
)

type IOrderProductRepository interface {
	Insert(orderId int64, productId int64, count int64, price float64) (int64, error)
	GetByOrderId(orderId int64) ([]models.OrderProduct, error)
}
type OrderProductRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewOrderProductRepository(conn *sql.DB) IOrderProductRepository {
	return &OrderProductRepository{
		DB: conn,
	}
}

func (r *OrderProductRepository) Insert(orderId int64, productId int64, count int64, price float64) (int64, error) {
	var orderProductId int64
	result, err := r.DB.Exec("INSERT order_product(order_id, product_id, count, price) VALUES(?, ?, ?, ?)",
		orderId, productId, count, price)
	if err != nil {
		log.Println(err)
		return orderProductId, err
	}
	orderProductId, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
		return orderProductId, err
	}
	return orderProductId, nil
}

func (r *OrderProductRepository) GetByOrderId(orderId int64) (orderProducts []models.OrderProduct, err error) {
	var orderProd models.OrderProduct
	rows, err := r.DB.Query("SELECT id, order_id, product_id, count, price FROM order_product WHERE order_id = (?)", orderId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&orderProd.Id, &orderProd.OrderId, &orderProd.ProductId, &orderProd.Count, &orderProd.Price)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		orderProducts = append(orderProducts, orderProd)
	}

	return orderProducts, nil
}
