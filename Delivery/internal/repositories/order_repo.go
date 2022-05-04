package repositories

import (
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
	"log"
)

type IOrderRepository interface {
	Insert(o *models.Order) (int64, error)
	GetById(id int64) (models.Order, error)
	GetByUserId(userId int64) ([]models.Order, error)
}

type OrderRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewOrderRepository(conn *sql.DB) IOrderRepository {
	return &OrderRepository{
		DB: conn,
	}
}

func (r *OrderRepository) Insert(o *models.Order) (orderId int64, err error) {
	result, err := r.DB.Exec("INSERT orders(price, user_id, address) VALUES(?, ?, ?)",
		o.Price, o.UserId, o.Address)
	if err != nil {
		log.Println(err)
		return orderId, err
	}
	orderId, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
		return orderId, err
	}

	return orderId, nil
}

func (r *OrderRepository) GetById(id int64) (order models.Order, err error) {
	err = r.DB.QueryRow("SELECT id, price, user_id, address, created_at FROM orders WHERE id = (?)", id).
		Scan(&order.Id, &order.Price, &order.UserId, &order.Address, &order.CreatedAt)
	if err != nil {
		log.Println(err)
		return order, err
	}
	return order, nil
}

func (r *OrderRepository) GetByUserId(userId int64) (orders []models.Order, err error) {
	var order models.Order
	rows, err := r.DB.Query("SELECT id, price, user_id, address, created_at "+
		"FROM orders WHERE user_id = (?)", userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&order.Id, &order.Price, &order.UserId, &order.Address, &order.CreatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
