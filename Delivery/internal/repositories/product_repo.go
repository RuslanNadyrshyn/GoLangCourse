package repositories

import (
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
	"fmt"
	"log"
)

type IProductRepository interface {
	Insert(p models.Product) (int64, error)
	UpdatePrice(p models.Product) (int64, error)
	GetAll() (products []models.Product, err error)
	GetById(id int64) (models.Product, error)
	GetBySupplierId(id int64) (products []models.Product, err error)
}
type ProductRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewProductRepository(conn *sql.DB) IProductRepository {
	return &ProductRepository{
		DB: conn,
	}
}

func (r *ProductRepository) Insert(p models.Product) (productId int64, err error) {
	if r.TX != nil {
		result, err := r.TX.Exec("INSERT products(id, menu_id, name, type, price, image) VALUES(?, ?, ?, ?, ?, ?)",
			p.Id, p.MenuId, p.Name, p.Type, p.Price, p.Image)
		if err != nil {
			log.Println(err)
			return productId, err
		}
		productId, err = result.LastInsertId()
		if err != nil {
			log.Println(err)
			return productId, err
		}

		return productId, nil
	}
	result, err := r.DB.Exec("INSERT products(id, menu_id, name, type, price, image) VALUES(?, ?, ?, ?, ?, ?)",
		p.Id, p.MenuId, p.Name, p.Type, p.Price, p.Image)
	if err != nil {
		log.Println(err)
		return productId, err
	}
	productId, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
		return productId, err
	}

	return productId, nil
}

func (r *ProductRepository) UpdatePrice(p models.Product) (productId int64, err error) {
	if r.TX != nil {
		result, err := r.TX.Exec("UPDATE products SET price=? WHERE id=?", p.Price, p.Id)
		if err != nil {
			fmt.Println(err)
			return productId, err
		}
		productId, err = result.LastInsertId()
		if err != nil {
			fmt.Println(err)
			return productId, err
		}
		return productId, nil
	}
	result, err := r.DB.Exec("UPDATE products SET price=? WHERE id=?", p.Price, p.Id)
	if err != nil {
		fmt.Println(err)
		return productId, err
	}
	productId, err = result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return productId, err
	}

	return productId, nil
}

func (r *ProductRepository) GetAll() (products []models.Product, err error) {
	var prod models.Product

	rows, err := r.DB.Query("SELECT id, menu_id, name, price, image, type FROM products")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	IngredientRepo := NewIngredientRepository(r.DB)
	for rows.Next() {
		err = rows.Scan(&prod.Id, &prod.MenuId, &prod.Name, &prod.Price, &prod.Image, &prod.Type)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		prod.Ingredients, err = IngredientRepo.GetByProductId(prod.Id)
		products = append(products, prod)
	}

	return products, nil
}

func (r *ProductRepository) GetById(id int64) (product models.Product, err error) {
	err = r.DB.QueryRow("SELECT id, menu_id, name, price, image, type FROM products WHERE id = (?)", id).
		Scan(&product.Id, &product.MenuId, &product.Name, &product.Price, &product.Image, &product.Type)
	if err != nil {
		log.Println(err)
		return product, err
	}

	IngredientRepo := NewIngredientRepository(r.DB)
	product.Ingredients, err = IngredientRepo.GetByProductId(product.Id)

	return product, nil
}

func (r *ProductRepository) GetBySupplierId(id int64) (products []models.Product, err error) {
	var prod models.Product

	rows, err := r.DB.Query("SELECT id, menu_id, name, price, image, type FROM products "+
		"WHERE menu_id = (SELECT id FROM menus WHERE supplier_id =(?))", id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	IngredientRepo := NewIngredientRepository(r.DB)
	for rows.Next() {
		err = rows.Scan(&prod.Id, &prod.MenuId, &prod.Name, &prod.Price, &prod.Image, &prod.Type)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		prod.Ingredients, err = IngredientRepo.GetByProductId(prod.Id)
		products = append(products, prod)
	}

	return products, nil
}
