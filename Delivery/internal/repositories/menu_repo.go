package repositories

import (
	"database/sql"
	"log"
)

type IMenuRepository interface {
	Insert(supId int) (int, error)
	GetById(id int) (int, error)
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

func (r *MenuRepository) Insert(supId int) (int, error) {
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

func (r *MenuRepository) GetById(id int) (supplierId int, err error) {
	err = r.DB.QueryRow("SELECT supplier_id FROM menus WHERE id = ?", id).Scan(&supplierId)
	if err != nil {
		return 0, err
	}
	return supplierId, nil
}
