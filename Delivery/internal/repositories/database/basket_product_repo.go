package database

import (
	"database/sql"
	"log"
)

type BasketProductDBRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewBasketProductRepository(conn *sql.DB) BasketProductDBRepository {
	return BasketProductDBRepository{
		DB: conn,
	}
}

func (r BasketProductDBRepository) Insert(basketId int, productId int, price float32) (int, error) {
	var basketProductId int64

	result, err := r.DB.Exec("INSERT INTO basket_product(basket_id, product_id, price)"+
		"VALUES (?, ?, ?)", basketId, productId, price)
	if err != nil {
		log.Println(err)
	}
	basketProductId, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return int(basketProductId), err
}

func (r BasketProductDBRepository) DeleteAll() error {
	_, err := r.DB.Exec("DELETE FROM basket_product")
	if err != nil {
		return err
	}
	return nil
}
