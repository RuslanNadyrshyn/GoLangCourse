package repositories

import (
	"database/sql"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories/models"
	"time"
)

type SupplierRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewSupplierRepo(conn *sql.DB) *SupplierRepo {
	return &SupplierRepo{
		DB: conn,
	}
}

func (r *SupplierRepo) Insert(s *models.Supplier) (id int64, err error) {
	result, err := r.DB.Exec("INSERT INTO suppliers(id, name, type, image, opening, closing) "+
		"VALUES(?, ?, ?, ?, ?, ?)", s.Id, s.Name, s.Type, s.Image, s.WorkingHours.Opening, s.WorkingHours.Closing)
	if err != nil {
		return 0, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *SupplierRepo) GetAll() (*[]models.Supplier, error) {
	var suppliers []models.Supplier
	rows, err := r.DB.Query("SELECT id, name, type, image, opening, closing FROM suppliers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s models.Supplier
		err = rows.Scan(&s.Id, &s.Name, &s.Type, &s.Image,
			&s.WorkingHours.Opening, &s.WorkingHours.Closing)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, s)
	}
	return &suppliers, nil
}

func (r *SupplierRepo) GetTypes() ([]string, error) {
	var types []string
	rows, err := r.DB.Query("SELECT DISTINCT type FROM suppliers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t string
		err = rows.Scan(&t)
		if err != nil {
			return nil, err
		}
		types = append(types, t)
	}

	return types, nil
}

func (r *SupplierRepo) GetById(id int64) (*models.Supplier, error) {
	var supplier models.Supplier
	err := r.DB.QueryRow("SELECT id, name, type, image, opening, closing FROM suppliers "+
		"WHERE id = ?", id).Scan(&supplier.Id, &supplier.Name, &supplier.Type, &supplier.Image,
		&supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing)
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (r *SupplierRepo) GetByType(t string) (*[]models.Supplier, error) {
	var suppliers []models.Supplier
	rows, err := r.DB.Query("SELECT id, name, type, image, opening, closing FROM suppliers WHERE type = (?)", t)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s models.Supplier
		err = rows.Scan(&s.Id, &s.Name, &s.Type, &s.Image,
			&s.WorkingHours.Opening, &s.WorkingHours.Closing)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, s)
	}
	return &suppliers, nil
}

func (r *SupplierRepo) GetByWorkingHours() (*[]models.Supplier, error) {
	var suppliers []models.Supplier
	time := time.Now().Format("15:04")

	rows, err := r.DB.Query("SELECT id, name, type, image, opening, closing "+
		"FROM suppliers WHERE (SELECT CONVERT(closing, TIME)) >= (?) "+
		"AND (SELECT CONVERT(opening, TIME)) <= (?)", time, time)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s models.Supplier
		err = rows.Scan(&s.Id, &s.Name, &s.Type, &s.Image,
			&s.WorkingHours.Opening, &s.WorkingHours.Closing)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, s)
	}
	return &suppliers, nil
}

func (r *SupplierRepo) GetByProductId(prodId int64) (*models.Supplier, error) {
	var supplier models.Supplier
	err := r.DB.QueryRow("SELECT id, name, type, image, opening, closing FROM suppliers WHERE id = "+
		"(SELECT supplier_id FROM menus WHERE id = (SELECT menu_id FROM products WHERE id = ?))",
		prodId).Scan(&supplier.Id, &supplier.Name, &supplier.Type, &supplier.Image,
		&supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing)
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (r *SupplierRepo) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *SupplierRepo) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Commit()
	}
	return nil
}

func (r *SupplierRepo) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Rollback()
	}
	return nil
}

func (r *SupplierRepo) GetTx() *sql.Tx {
	return r.TX
}

func (r *SupplierRepo) SetTx(tx *sql.Tx) {
	r.TX = tx
}
