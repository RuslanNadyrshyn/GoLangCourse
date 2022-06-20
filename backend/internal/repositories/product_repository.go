package repositories

import (
	"database/sql"
	"fmt"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories/models"
	"time"
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

func (r *ProductRepo) GetAll() ([]models.Product, error) {
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

	return products, nil
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

func (r *ProductRepo) GetByType(t string) ([]models.Product, error) {
	var products []models.Product
	rows, err := r.DB.Query("SELECT id, menu_id, name, price, image, type FROM products WHERE type = (?)", t)
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

func (r *ProductRepo) GetByWorkingHours(prodType string) ([]models.Product, error) {
	var products []models.Product
	location, _ := time.LoadLocation("EET")
	timeNow := time.Now().In(location).Format("15:04")
	rows, err := r.DB.Query(fmt.Sprintf("SELECT id, menu_id, name, price, image, type "+
		"FROM products WHERE menu_id IN "+
		"(SELECT id FROM menus WHERE supplier_id IN "+
		"(SELECT id FROM suppliers WHERE (SELECT CONVERT(closing, TIME)) >= '%s' "+
		"AND (SELECT CONVERT(opening, TIME)) <= '%s'))%s", timeNow, timeNow, prodType))
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

func (r *ProductRepo) GetBySupId(supId int64, prodType string) ([]models.Product, error) {
	var products []models.Product
	rows, err := r.DB.Query(fmt.Sprintf("SELECT id, menu_id, name, price, image, type "+
		"FROM products WHERE menu_id IN "+
		"(SELECT id FROM menus WHERE supplier_id IN "+
		"(SELECT id FROM suppliers WHERE id = %d))%s", supId, prodType))
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

func (r *ProductRepo) GetBySupType(supType string, prodType string) ([]models.Product, error) {
	var products []models.Product
	rows, err := r.DB.Query(fmt.Sprintf("SELECT id, menu_id, name, price, image, type "+
		"FROM products WHERE menu_id IN "+
		"(SELECT id FROM menus WHERE supplier_id IN "+
		"(SELECT id FROM suppliers WHERE type = '%s'))%s", supType, prodType))
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

func (r *ProductRepo) GetAllTypes() ([]string, error) {
	var types []string
	rows, err := r.DB.Query("SELECT DISTINCT type FROM products")
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

func (r *ProductRepo) GetTypesBySupType(supType string) ([]string, error) {
	var types []string
	var query string
	if supType == "workingHours" {
		timeNow := time.Now().Format("15:04")
		query = fmt.Sprintf("SELECT DISTINCT type FROM products WHERE menu_id IN "+
			"(SELECT id FROM menus WHERE supplier_id IN "+
			"(SELECT id FROM suppliers WHERE "+
			"(SELECT CONVERT(closing, TIME)) >= '%s' AND "+
			"(SELECT CONVERT(opening, TIME)) <= '%s'))", timeNow, timeNow)
	} else {
		query = fmt.Sprintf("SELECT DISTINCT type FROM products WHERE menu_id IN "+
			"(SELECT id FROM menus WHERE supplier_id IN "+
			"(SELECT id FROM suppliers WHERE type = '%s'))", supType)
	}
	rows, err := r.DB.Query(query)
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

func (r *ProductRepo) GetTypesBySupId(supId int64) ([]string, error) {
	var types []string
	rows, err := r.DB.Query(fmt.Sprintf("SELECT DISTINCT type FROM products WHERE menu_id IN "+
		"(SELECT id FROM menus WHERE supplier_id IN "+
		"(SELECT id FROM suppliers WHERE id = %d))", supId))
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
