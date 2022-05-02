package repositories

import (
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
	"fmt"
	"log"
)

type IOrderProductRepository interface {
	Insert(orderId int, productId int, count int, price float64) (int, error)
	GetByOrderId(orderId int) ([]models.OrderProduct, error)
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

func (r *OrderProductRepository) Insert(orderId int, productId int, count int, price float64) (int, error) {
	var orderProductId int64
	result, err := r.DB.Exec("INSERT order_product(order_id, product_id, count, price) VALUES(?, ?, ?, ?)",
		orderId, productId, count, price)
	if err != nil {
		log.Println(err)
		return int(orderProductId), err
	}
	orderProductId, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
		return int(orderProductId), err
	}
	return int(orderProductId), err
}

func (r *OrderProductRepository) GetByOrderId(orderId int) (orderProducts []models.OrderProduct, err error) {
	var orderProd models.OrderProduct
	rows, err := r.DB.Query("SELECT id, order_id, product_id, count, price FROM order_product WHERE order_id = (?)", orderId)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&orderProd.Id, &orderProd.OrderId, &orderProd.ProductId, &orderProd.Count, &orderProd.Price)
		if err != nil {
			panic(err)
		}
		orderProducts = append(orderProducts, orderProd)
	}

	return orderProducts, nil
}
