package repositories

import (
	"database/sql"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories/models"
)

type OrderProductRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewOrderProductRepo(conn *sql.DB) *OrderProductRepo {
	return &OrderProductRepo{
		DB: conn,
	}
}

func (r *OrderProductRepo) Insert(orderId int64, productId int64, count int64, price float64) (int64, error) {
	var orderProductId int64
	if r.TX != nil {
		result, err := r.TX.Exec("INSERT order_product(order_id, product_id, count, price) VALUES(?, ?, ?, ?)",
			orderId, productId, count, price)
		if err != nil {
			return orderProductId, err
		}
		orderProductId, err = result.LastInsertId()
		if err != nil {
			return orderProductId, err
		}
		return orderProductId, nil
	}
	result, err := r.DB.Exec("INSERT order_product(order_id, product_id, count, price) VALUES(?, ?, ?, ?)",
		orderId, productId, count, price)
	if err != nil {
		return orderProductId, err
	}
	orderProductId, err = result.LastInsertId()
	if err != nil {
		return orderProductId, err
	}
	return orderProductId, nil
}

func (r *OrderProductRepo) GetByOrderId(orderId int64) (*[]models.OrderProduct, error) {
	var orderProducts []models.OrderProduct
	rows, err := r.DB.Query("SELECT id, order_id, product_id, count, price FROM order_product WHERE order_id = (?)", orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var op models.OrderProduct
		err = rows.Scan(&op.Id, &op.OrderId, &op.ProductId, &op.Count, &op.Price)
		if err != nil {
			return nil, err
		}
		orderProducts = append(orderProducts, op)
	}

	return &orderProducts, nil
}

func (r *OrderProductRepo) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *OrderProductRepo) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Commit()
	}
	return nil
}

func (r *OrderProductRepo) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Rollback()
	}
	return nil
}

func (r *OrderProductRepo) GetTx() *sql.Tx {
	return r.TX
}

func (r *OrderProductRepo) SetTx(tx *sql.Tx) {
	r.TX = tx
}
