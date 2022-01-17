package database

import (
	"awesomeProject/internal/repositories/models"
	"database/sql"
	"log"
)

type ProductDBRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewProductRepository(conn *sql.DB) ProductDBRepository {
	return ProductDBRepository{
		DB: conn,
	}
}

func (r ProductDBRepository) Insert(p models.Product) (int, error) {
	var productId int

	if r.TX != nil {
		err := r.TX.QueryRow("INSERT products(name, price, image, type) VALUES(?, ?, ?, ?) RETURNING product_id",
			p.Name, p.Price, p.Image, p.Type).Scan(&productId)
		if err != nil {
			_ = r.TX.Rollback()
		}

		err = r.TX.Commit()
		if err != nil {
			_ = r.TX.Rollback()
		}
		return productId, err
	}
	_, err := r.DB.Exec("INSERT products(name, price, image, type) VALUES(?, ?, ?, ?)",
		p.Name, p.Price, p.Image, p.Type)

	if err != nil {
		log.Panic(err)
		return 0, err
	}

	return productId, err
}
