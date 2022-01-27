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
	var productId int64

	if r.TX != nil {
		result, err := r.TX.Exec("INSERT products(menu_id, name, type, price, image) VALUES(?, ?, ?, ?, ?)",
			p.MenuId, p.Name, p.Type, p.Price, p.Image)
		if err != nil {
			log.Println(err)
			return int(productId), err
		}
		productId, err = result.LastInsertId()
		if err != nil {
			log.Println(err)
			return int(productId), err
		}

		return int(productId), err
	}
	result, err := r.DB.Exec("INSERT products(menu_id, name, type, price, image) VALUES(?, ?, ?, ?, ?)",
		p.MenuId, p.Name, p.Type, p.Price, p.Image)
	if err != nil {
		log.Println(err)
		return int(productId), err
	}
	productId, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
		return int(productId), err
	}

	return int(productId), err
}

func (r *ProductDBRepository) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *ProductDBRepository) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.CommitTx()
	}
	return nil
}

func (r *ProductDBRepository) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.RollbackTx()
	}
	return nil
}

func (r *ProductDBRepository) GetTx() *sql.Tx {
	return r.TX
}

func (r *ProductDBRepository) SetTx(tx *sql.Tx) {
	r.TX = tx
}
