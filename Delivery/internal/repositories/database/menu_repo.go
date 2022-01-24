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

func (r MenuDBRepository) Insert(s models.Supplier, supId int) (int, error) {
	var menuId int64

	if r.TX != nil {
		productRepo := NewProductRepository(r.DB)

		productRepo.TX = r.TX

		result, err := r.TX.Exec("INSERT INTO menus(supplier_id) VALUES(?)", supId)
		if err != nil {
			log.Println(err)
			return int(menuId), err
		}

		menuId, err = result.LastInsertId()
		for i := range s.Menu {
			s.Menu[i].MenuId = int(menuId)
			productRepo.Insert(s.Menu[i])
			if err != nil {
				log.Println(err)
				return int(menuId), err
			}
		}

		return int(menuId), err
	}
	return 0, nil
}
