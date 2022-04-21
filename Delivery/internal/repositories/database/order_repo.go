package database

import (
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
	"log"
)

type OrderDBRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewOrderRepository(conn *sql.DB) OrderDBRepository {
	return OrderDBRepository{
		DB: conn,
	}
}

func (odbr OrderDBRepository) Insert(o *models.Order) (int, error) {
	var orderId int64

	result, err := odbr.DB.Exec("INSERT orders(price, user_id, address) VALUES(?, ?, ?, ?)",
		o.Price, o.UserId, o.Address)
	if err != nil {
		log.Println(err)
		return int(orderId), err
	}
	orderId, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
		return int(orderId), err
	}

	return int(orderId), err
}

func (odbr OrderDBRepository) InsertOrderProduct(orderId int, productId int, count int, price float64) (int, error) {
	var orderProductId int64

	result, err := odbr.DB.Exec("INSERT order_product(order_id, product_id, count, price) VALUES(?, ?, ?, ?)",
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
