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
	if r.TX != nil {
		var id int64
		result, err := r.TX.Exec("INSERT INTO suppliers(name, type, address, image, opening, closing) VALUES(?, ?, ?, ?, ?, ?)",
			s.Name, s.Type, s.Address, s.Image, s.WorkingHours.Opening, s.WorkingHours.Closing)
		if err != nil {
			log.Println(err)
			return int(id), err
		}
		id, err = result.LastInsertId()
		if err != nil {
			log.Println(err)
			return int(id), err
		}
		return int(id), err
	}
	return 0, nil
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

func (r *SupplierDBRepository) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *SupplierDBRepository) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return nil
		//return r.CommitTx()
	}
	return nil
}

func (r *SupplierDBRepository) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return nil
		//return r.RollbackTx()
	}
	return nil
}

func (r *SupplierDBRepository) GetTx() *sql.Tx {
	return r.TX
}

func (r *SupplierDBRepository) SetTx(tx *sql.Tx) {
	r.TX = tx
}
