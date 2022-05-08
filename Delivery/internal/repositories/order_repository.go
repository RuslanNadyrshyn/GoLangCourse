package repositories

import (
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
	"log"
)

type OrderRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewOrderRepo(conn *sql.DB) *OrderRepo {
	return &OrderRepo{
		DB: conn,
	}
}

func (r *OrderRepo) Insert(o *models.Order) (orderId int64, err error) {
	if r.TX != nil {
		result, err := r.TX.Exec("INSERT orders(price, user_id, address) VALUES(?, ?, ?)",
			o.Price, o.UserId, o.Address)
		if err != nil {
			return 0, err
		}
		orderId, err = result.LastInsertId()
		if err != nil {
			return 0, err
		}

		return orderId, nil
	}
	result, err := r.DB.Exec("INSERT orders(price, user_id, address) VALUES(?, ?, ?)",
		o.Price, o.UserId, o.Address)
	if err != nil {
		return 0, err
	}
	orderId, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return orderId, nil
}

func (r *OrderRepo) GetById(id int64) (*models.Order, error) {
	var o models.Order
	err := r.DB.QueryRow("SELECT id, price, user_id, address, created_at FROM orders WHERE id = (?)", id).
		Scan(&o.Id, &o.Price, &o.UserId, &o.Address, &o.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *OrderRepo) GetByUserId(userId int64) (*[]models.Order, error) {
	var orders []models.Order
	rows, err := r.DB.Query("SELECT id, price, user_id, address, created_at "+
		"FROM orders WHERE user_id = (?)", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o models.Order
		err = rows.Scan(&o.Id, &o.Price, &o.UserId, &o.Address, &o.CreatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return &orders, nil
}

func (r *OrderRepo) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *OrderRepo) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Commit()
	}
	return nil
}

func (r *OrderRepo) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Rollback()
	}
	return nil
}

func (r *OrderRepo) GetTx() *sql.Tx {
	return r.TX
}

func (r *OrderRepo) SetTx(tx *sql.Tx) {
	r.TX = tx
}
