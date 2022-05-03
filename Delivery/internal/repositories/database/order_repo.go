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

func (odbr *OrderDBRepository) Insert(o *models.Order) (int, error) {
	var orderId int64

	result, err := odbr.DB.Exec("INSERT orders(price, user_id, address) VALUES(?, ?, ?)",
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

func (odbr *OrderDBRepository) GetById(id int) (models.Order, error) {
	var order models.Order
	err := odbr.DB.QueryRow("SELECT id, price, user_id, address, created_at FROM orders WHERE id = (?)", id).
		Scan(&order.Id, &order.Price, &order.UserId, &order.Address, &order.CreatedAt)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (odbr *OrderDBRepository) GetByUserId(userId int) (orders []models.Order, err error) {
	var order models.Order
	rows, err := odbr.DB.Query("SELECT id, price, user_id, address, created_at "+
		"FROM orders WHERE user_id = (?)", userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&order.Id, &order.Price, &order.UserId, &order.Address, &order.CreatedAt)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}

	return orders, nil
}
