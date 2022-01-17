package database

import (
	"awesomeProject/internal/repositories/models"
	"database/sql"
	"fmt"
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
	productRepo := NewProductRepository(r.DB)
	var newProduct models.Product

	if r.TX != nil {
		err := r.TX.QueryRow("INSERT suppliers(id, name, type, address, image, opening, closing) VALUES(?, ?, ?,?, ?,?, ?) RETURNING id",
			s.Id, s.Name, s.Type, s.Address, s.Image, s.WorkingHours.Opening, s.WorkingHours.Closing).Scan(&id)
		if err != nil {
			_ = r.TX.Rollback()
		}

		//
		for i := range s.Menu {
			newProduct.Name, newProduct.Type, newProduct.Price, newProduct.Image =
				s.Menu[i].Name, s.Menu[i].Type, s.Menu[i].Price, s.Menu[i].Image
			productRepo.Insert(newProduct)
			fmt.Println("Added menu")

		}

		err = r.TX.Commit()
		if err != nil {
			_ = r.TX.Rollback()
		}
		return id, err
	}

	_, err := r.DB.Exec("INSERT suppliers(id, name, type, address, image, opening, closing) VALUES(?, ?, ?, ?, ?, ?, ?)",
		s.Id, s.Name, s.Type, s.Address, s.Image, s.WorkingHours.Opening, s.WorkingHours.Closing)

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
