package database

import (
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
	"log"
)

type BasketDBRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewBasketRepository(conn *sql.DB) BasketDBRepository {
	return BasketDBRepository{
		DB: conn,
	}
}

func (r BasketDBRepository) Insert(b models.Basket) (int, error) {
	var basketId int64

	result, err := r.DB.Exec("INSERT baskets(user_id, price) "+
		"VALUES(?, ?)", b.UserId, b.Price)
	if err != nil {
		log.Println(err)
	}
	basketId, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
	}

	return int(basketId), err
}
