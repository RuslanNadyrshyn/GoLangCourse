package repositories

import (
	"database/sql"
	"log"
)

type IMenuRepository interface {
	Insert(supId int64) (int64, error)
	GetById(id int64) (int64, error)
}

type MenuRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewMenuRepository(conn *sql.DB) IMenuRepository {
	return &MenuRepository{
		DB: conn,
	}
}

func (r *MenuRepository) Insert(supId int64) (int64, error) {
	var menuId int64

	if r.TX != nil {
		result, err := r.TX.Exec("INSERT INTO menus(supplier_id) VALUES(?)", supId)
		if err != nil {
			log.Println(err)
			return menuId, err
		}
		menuId, err = result.LastInsertId()
		if err != nil {
			log.Println(err)
			return menuId, err
		}
		return menuId, nil
	}
	result, err := r.DB.Exec("INSERT INTO menus(supplier_id) VALUES(?)", supId)
	if err != nil {
		log.Println(err)
		return menuId, err
	}
	menuId, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
		return menuId, err
	}
	return menuId, nil
}

func (r *MenuRepository) GetById(id int64) (supplierId int64, err error) {
	err = r.DB.QueryRow("SELECT supplier_id FROM menus WHERE id = ?", id).Scan(&supplierId)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return supplierId, nil
}
