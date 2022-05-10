package repositories

import (
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
)

type ProductRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewProductRepo(conn *sql.DB) *ProductRepo {
	return &ProductRepo{
		DB: conn,
	}
}

func (r *ProductRepo) Insert(p *models.Product) (productId int64, err error) {
	if r.TX != nil {
		result, err := r.TX.Exec("INSERT products(id, menu_id, name, type, price, image) VALUES(?, ?, ?, ?, ?, ?)",
			p.Id, p.MenuId, p.Name, p.Type, p.Price, p.Image)
		if err != nil {
			return 0, err
		}
		productId, err = result.LastInsertId()
		if err != nil {
			return 0, err
		}

		return productId, nil
	}
	result, err := r.DB.Exec("INSERT products(id, menu_id, name, type, price, image) VALUES(?, ?, ?, ?, ?, ?)",
		p.Id, p.MenuId, p.Name, p.Type, p.Price, p.Image)
	if err != nil {
		return 0, err
	}
	productId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return productId, nil
}

func (r *ProductRepo) UpdatePrice(p *models.Product) (productId int64, err error) {
	_, err = r.DB.Exec("UPDATE products SET price=? WHERE id=?", p.Price, p.Id)
	if err != nil {
		return 0, err
	}

	return p.Id, nil
}

func (r *ProductRepo) GetAll() (*[]models.Product, error) {
	var products []models.Product
	rows, err := r.DB.Query("SELECT id, menu_id, name, price, image, type FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err = rows.Scan(&p.Id, &p.MenuId, &p.Name, &p.Price, &p.Image, &p.Type)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return &products, nil
}

func (r *ProductRepo) GetById(id int64) (*models.Product, error) {
	var product models.Product
	err := r.DB.QueryRow("SELECT id, menu_id, name, price, image, type FROM products WHERE id = (?)", id).
		Scan(&product.Id, &product.MenuId, &product.Name, &product.Price, &product.Image, &product.Type)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepo) GetByMenuId(menuId int64) ([]models.Product, error) {
	var products []models.Product

	rows, err := r.DB.Query("SELECT id, menu_id, name, price, image, type FROM products "+
		"WHERE menu_id = (?)", menuId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err = rows.Scan(&p.Id, &p.MenuId, &p.Name, &p.Price, &p.Image, &p.Type)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepo) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *ProductRepo) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Commit()
	}
	return nil
}

func (r *ProductRepo) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Rollback()
	}
	return nil
}

func (r *ProductRepo) GetTx() *sql.Tx {
	return r.TX
}

func (r *ProductRepo) SetTx(tx *sql.Tx) {
	r.TX = tx
}
