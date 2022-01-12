package database

import (
	"awesomeProject/internal/repositories/models"
	"database/sql"
	"log"
)

type SupplierDBRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewSupplierRepository(conn *sql.DB) SupplierDBRepository {
	return SupplierDBRepository{
		DB: conn,
	}
}

func (r SupplierDBRepository) Insert(s models.Supplier) (int, error) {
	var id int

	if r.TX != nil {
		err := r.TX.QueryRow("INSERT suppliers(name, address) VALUES(?, ?) RETURNING id", s.Name, s.Address).Scan(&id)
		if err != nil {
			_ = r.TX.Rollback()
		}
		err = r.TX.Commit()
		if err != nil {
			_ = r.TX.Rollback()
		}
		return id, err
	}
	_, err := r.DB.Exec("INSERT suppliers(name, address) VALUES(?, ?)", s.Name, s.Address)

	if err != nil {
		log.Panic(err)
		return 0, err
	}

	return id, err
}

func (r SupplierDBRepository) GetById(id int) (models.Supplier, error) {
	var supplier models.Supplier

	err := r.DB.QueryRow("SELECT id, name, address FROM suppliers WHERE id = ?", id).Scan(&supplier.Id, &supplier.Name, &supplier.Address)
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (r SupplierDBRepository) Delete(name string) error {
	_, err := r.DB.Exec("DELETE FROM suppliers WHERE name = ?", name)
	if err != nil {
		return err
	}

	return nil
}
