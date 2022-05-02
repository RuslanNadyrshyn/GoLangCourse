package repositories

import (
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
	"log"
)

type ISupplierRepository interface {
	Insert(s *models.Supplier) (int, error)
	GetAll() ([]models.Supplier, error)
	GetById(id int) (models.Supplier, error)
	GetByWorkingTime(workingTime string) ([]models.Supplier, error)
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

func (r *SupplierRepository) Insert(s *models.Supplier) (int, error) {
	var id int64

	if r.TX != nil {
		result, err := r.TX.Exec("INSERT INTO suppliers(id, name, type, image, opening, closing) "+
			"VALUES(?, ?, ?, ?, ?, ?)", s.Id, s.Name, s.Type, s.Image, s.WorkingHours.Opening, s.WorkingHours.Closing)
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

	result, err := r.DB.Exec("INSERT INTO suppliers(id, name, type, image, opening, closing) "+
		"VALUES(?, ?, ?, ?, ?, ?)", s.Id, s.Name, s.Type, s.Image, s.WorkingHours.Opening, s.WorkingHours.Closing)
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

func (r *SupplierRepository) GetAll() (suppliers []models.Supplier, err error) {
	var sup models.Supplier
	rows, err := r.DB.Query("SELECT id, name, type, image, opening, closing FROM suppliers")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&sup.Id, &sup.Name, &sup.Type, &sup.Image,
			&sup.WorkingHours.Opening, &sup.WorkingHours.Closing)
		if err != nil {
			panic(err)
		}
		suppliers = append(suppliers, sup)
	}
	return suppliers, nil
}

func (r *SupplierRepository) GetById(id int) (models.Supplier, error) {
	var supplier models.Supplier
	err := r.DB.QueryRow("SELECT id, name, type, image, opening, closing FROM suppliers "+
		"WHERE id = ?", id).Scan(&supplier.Id, &supplier.Name, &supplier.Type, &supplier.Image,
		&supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing)
	if err != nil {
		return supplier, err
	}
	return supplier, nil
}

func (r *SupplierRepository) GetProductById(id int) (models.Supplier, error) {
	var supplier models.Supplier
	err := r.DB.QueryRow("SELECT id, name, type, image, opening, closing FROM suppliers "+
		"WHERE id = ?", id).Scan(&supplier.Id, &supplier.Name, &supplier.Type, &supplier.Image,
		&supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing)
	if err != nil {
		return supplier, err
	}
	return supplier, nil
}

func (r *SupplierRepository) GetByWorkingTime(workingTime string) (suppliers []models.Supplier, err error) {
	var sup models.Supplier
	rows, err := r.DB.Query("SELECT id, name, type, address, image, opening, closing FROM suppliers "+
		"WHERE opening <= (?) AND closing >= (?)", workingTime, workingTime)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&sup.Id, &sup.Name, &sup.Type, &sup.Image,
			&sup.WorkingHours.Opening, &sup.WorkingHours.Closing)
		if err != nil {
			panic(err)
		}
		suppliers = append(suppliers, sup)
	}
	return suppliers, nil
}
