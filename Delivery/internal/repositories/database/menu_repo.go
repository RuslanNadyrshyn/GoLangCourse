package database

import (
	"awesomeProject/internal/repositories/models"
	"database/sql"
	"log"
)

type MenuDBRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewMenuRepository(conn *sql.DB) MenuDBRepository {
	return MenuDBRepository{
		DB: conn,
	}
}

func (r MenuDBRepository) Insert(menu []models.Product, supId int) (int, error) {
	var menuId int64

	if r.TX != nil {
		result, err := r.TX.Exec("INSERT INTO menus(supplier_id) VALUES(?)", supId)
		if err != nil {
			log.Println(err)
			return int(menuId), err
		}
		menuId, err = result.LastInsertId()
		if err != nil {
			log.Println(err)
			return int(menuId), err
		}
		return int(menuId), err
	}
	result, err := r.DB.Exec("INSERT INTO menus(supplier_id) VALUES(?)", supId)
	if err != nil {
		log.Println(err)
		return int(menuId), err
	}
	menuId, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
		return int(menuId), err
	}
	return int(menuId), err
}

func (r *MenuDBRepository) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *MenuDBRepository) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.CommitTx()
	}
	return nil
}

func (r *MenuDBRepository) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.RollbackTx()
	}
	return nil
}

func (r *MenuDBRepository) GetTx() *sql.Tx {
	return r.TX
}

func (r *MenuDBRepository) SetTx(tx *sql.Tx) {
	r.TX = tx
}