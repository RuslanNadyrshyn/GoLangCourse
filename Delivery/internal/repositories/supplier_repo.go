package repositories

import (
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
	"log"
)

type ISupplierRepository interface {
	Insert(s *models.Supplier) (int64, error)
	GetAll() ([]models.Supplier, error)
	GetById(id int64) (models.Supplier, error)
	GetByProductId(id int64) (models.Supplier, error)
}

type SupplierRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewSupplierRepository(conn *sql.DB) ISupplierRepository {
	return &SupplierRepository{
		DB: conn,
	}
}

func (r *SupplierRepository) Insert(s *models.Supplier) (id int64, err error) {
	result, err := r.DB.Exec("INSERT INTO suppliers(id, name, type, image, opening, closing) "+
		"VALUES(?, ?, ?, ?, ?, ?)", s.Id, s.Name, s.Type, s.Image, s.WorkingHours.Opening, s.WorkingHours.Closing)
	if err != nil {
		log.Println(err)
		return id, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
		return id, err
	}
	return id, nil
}

func (r *SupplierRepository) GetAll() (suppliers []models.Supplier, err error) {
	var sup models.Supplier
	rows, err := r.DB.Query("SELECT id, name, type, image, opening, closing FROM suppliers")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&sup.Id, &sup.Name, &sup.Type, &sup.Image,
			&sup.WorkingHours.Opening, &sup.WorkingHours.Closing)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		suppliers = append(suppliers, sup)
	}
	return suppliers, nil
}

func (r *SupplierRepository) GetById(id int64) (supplier models.Supplier, err error) {
	err = r.DB.QueryRow("SELECT id, name, type, image, opening, closing FROM suppliers "+
		"WHERE id = ?", id).Scan(&supplier.Id, &supplier.Name, &supplier.Type, &supplier.Image,
		&supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing)
	if err != nil {
		log.Println(err)
		return supplier, err
	}
	return supplier, nil
}

func (r *SupplierRepository) GetProductById(id int) (supplier models.Supplier, err error) {
	err = r.DB.QueryRow("SELECT id, name, type, image, opening, closing FROM suppliers "+
		"WHERE id = ?", id).Scan(&supplier.Id, &supplier.Name, &supplier.Type, &supplier.Image,
		&supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing)
	if err != nil {
		log.Println(err)
		return supplier, err
	}
	return supplier, nil
}

func (r *SupplierRepository) GetByProductId(id int64) (supplier models.Supplier, err error) {
	err = r.DB.QueryRow("SELECT id, name, type, image, opening, closing FROM suppliers "+
		"WHERE id = (SELECT supplier_id FROM menus WHERE id = (?))", id).
		Scan(&supplier.Id, &supplier.Name, &supplier.Type, &supplier.Image,
			&supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing)
	if err != nil {
		log.Println(err)
		return supplier, err
	}
	return supplier, nil
}
