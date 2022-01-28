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
	var id int64

	if r.TX != nil {
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

	result, err := r.DB.Exec("INSERT INTO suppliers(name, type, address, image, opening, closing) VALUES(?, ?, ?, ?, ?, ?)",
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

func (r SupplierDBRepository) GetAll() (suppliers []models.Supplier, err error) {
	var sup models.Supplier
	rows, err := r.DB.Query("SELECT id, name, type, address, image, opening, closing FROM suppliers")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&sup.Id, &sup.Name, &sup.Type, &sup.Address,
			&sup.Image, &sup.WorkingHours.Opening, &sup.WorkingHours.Closing)
		if err != nil {
			panic(err)
		}
		suppliers = append(suppliers, sup)
	}
	return suppliers, nil
}

func (r SupplierDBRepository) GetById(id int) (supplier models.Supplier, err error) {
	err = r.DB.QueryRow("SELECT id, name, type, address, image, opening, closing FROM suppliers "+
		"WHERE id = ?", id).Scan(&supplier.Id, &supplier.Name, &supplier.Type, &supplier.Address,
		&supplier.Image, &supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing)
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (r SupplierDBRepository) GetByName(name string) (suppliers []models.Supplier, err error) {
	var sup models.Supplier
	rows, err := r.DB.Query("SELECT id, name, type, address, image, opening, closing FROM suppliers "+
		"WHERE name = ?", name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&sup.Id, &sup.Name, &sup.Type, &sup.Address,
			&sup.Image, &sup.WorkingHours.Opening, &sup.WorkingHours.Closing)
		if err != nil {
			panic(err)
		}
		suppliers = append(suppliers, sup)
	}
	return suppliers, nil
}

func (r SupplierDBRepository) GetByType(t string) (suppliers []models.Supplier, err error) {
	var sup models.Supplier
	rows, err := r.DB.Query("SELECT id, name, type, address, image, opening, closing FROM suppliers "+
		"WHERE type = ?", t)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&sup.Id, &sup.Name, &sup.Type, &sup.Address,
			&sup.Image, &sup.WorkingHours.Opening, &sup.WorkingHours.Closing)
		if err != nil {
			panic(err)
		}
		suppliers = append(suppliers, sup)
	}
	return suppliers, nil
}

func (r SupplierDBRepository) GetByAddress(address string) (suppliers []models.Supplier, err error) {
	var sup models.Supplier
	rows, err := r.DB.Query("SELECT id, name, type, address, image, opening, closing FROM suppliers "+
		"WHERE address LIKE '%(?)%", address)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&sup.Id, &sup.Name, &sup.Type, &sup.Address,
			&sup.Image, &sup.WorkingHours.Opening, &sup.WorkingHours.Closing)
		if err != nil {
			panic(err)
		}
		suppliers = append(suppliers, sup)
	}
	return suppliers, nil
}

func (r SupplierDBRepository) GetByWorkingTime(workingTime string) (suppliers []models.Supplier, err error) {
	var sup models.Supplier
	rows, err := r.DB.Query("SELECT id, name, type, address, image, opening, closing FROM suppliers "+
		"WHERE opening <= (?) AND closing >= (?)", workingTime, workingTime)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&sup.Id, &sup.Name, &sup.Type, &sup.Address,
			&sup.Image, &sup.WorkingHours.Opening, &sup.WorkingHours.Closing)
		if err != nil {
			panic(err)
		}
		suppliers = append(suppliers, sup)
	}
	return suppliers, nil
}

func (r SupplierDBRepository) DeleteAll() error {
	_, err := r.DB.Exec("DELETE FROM suppliers")
	if err != nil {
		return err
	}
	return nil
}

func (r SupplierDBRepository) DeleteByName(name string) error {
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
