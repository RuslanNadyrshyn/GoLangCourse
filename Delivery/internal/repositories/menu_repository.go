package repositories

import (
	"database/sql"
)

type MenuRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewMenuRepo(conn *sql.DB) *MenuRepo {
	return &MenuRepo{
		DB: conn,
	}
}

func (r *MenuRepo) Insert(supId int64) (menuId int64, err error) {
	if r.TX != nil {
		result, err := r.TX.Exec("INSERT INTO menus(supplier_id) VALUES(?)", supId)
		if err != nil {
			return 0, err
		}
		menuId, err = result.LastInsertId()
		if err != nil {
			return 0, err
		}
		return menuId, nil
	}
	result, err := r.DB.Exec("INSERT INTO menus(supplier_id) VALUES(?)", supId)
	if err != nil {
		return 0, err
	}
	menuId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return menuId, nil
}

func (r *MenuRepo) GetById(id int64) (supplierId int64, err error) {
	err = r.DB.QueryRow("SELECT supplier_id FROM menus WHERE id = ?", id).Scan(&supplierId)
	if err != nil {
		return 0, err
	}
	return supplierId, nil
}

func (r *MenuRepo) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *MenuRepo) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.CommitTx()
	}
	return nil
}

func (r *MenuRepo) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.RollbackTx()
	}
	return nil
}

func (r *MenuRepo) GetTx() *sql.Tx {
	return r.TX
}

func (r *MenuRepo) SetTx(tx *sql.Tx) {
	r.TX = tx
}
