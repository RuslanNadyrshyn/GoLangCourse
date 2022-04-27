package database

import (
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
	"fmt"
	"log"
)

type ProductDBRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewProductRepository(conn *sql.DB) ProductDBRepository {
	return ProductDBRepository{
		DB: conn,
	}
}

func (r ProductDBRepository) Insert(p models.Product) (int, error) {
	var productId int64

	if r.TX != nil {
		result, err := r.TX.Exec("INSERT products(id, menu_id, name, type, price, image) VALUES(?, ?, ?, ?, ?, ?)",
			p.Id, p.MenuId, p.Name, p.Type, p.Price, p.Image)
		if err != nil {
			log.Println(err)
			return int(productId), err
		}
		productId, err = result.LastInsertId()
		if err != nil {
			log.Println(err)
			return int(productId), err
		}

		return int(productId), err
	}
	result, err := r.DB.Exec("INSERT products(id, menu_id, name, type, price, image) VALUES(?, ?, ?, ?, ?, ?)",
		p.Id, p.MenuId, p.Name, p.Type, p.Price, p.Image)
	if err != nil {
		log.Println(err)
		return int(productId), err
	}
	productId, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
		return int(productId), err
	}

	return int(productId), err
}

func (r ProductDBRepository) UpdatePrice(p models.Product) (int, error) {
	var productId int64

	if r.TX != nil {
		result, err := r.TX.Exec("UPDATE products SET price=? WHERE id=?", p.Price, p.Id)
		if err != nil {
			fmt.Println(err)
			return int(productId), err
		}
		productId, err = result.LastInsertId()
		if err != nil {
			fmt.Println(err)
			return int(productId), err
		}
		return int(productId), err
	}
	result, err := r.DB.Exec("UPDATE products SET price=? WHERE id=?", p.Price, p.Id)
	if err != nil {
		fmt.Println(err)
		return int(productId), err
	}
	productId, err = result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return int(productId), err
	}

	return int(productId), err
}

func (r ProductDBRepository) GetAll() (products []models.Product, err error) {
	var prod models.Product

	rows, err := r.DB.Query("SELECT id, menu_id, name, price, image, type FROM products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&prod.Id, &prod.MenuId, &prod.Name, &prod.Price, &prod.Image, &prod.Type)
		if err != nil {
			panic(err)
		}
		products = append(products, prod)
	}
	IngredientRepo := NewIngredientRepository(r.DB)

	for i := range products {
		products[i].Ingredients, err = IngredientRepo.GetByProductId(products[i].Id)
	}
	return products, nil
}

func (r ProductDBRepository) GetById(id int) (models.Product, error) {
	var prod models.Product

	err := r.DB.QueryRow("SELECT id, menu_id, name, price, image, type FROM products WHERE id = (?)", id).
		Scan(&prod.Id, &prod.MenuId, &prod.Name, &prod.Price, &prod.Image, &prod.Type)
	if err != nil {
		panic(err)
	}

	IngredientRepo := NewIngredientRepository(r.DB)
	prod.Ingredients, err = IngredientRepo.GetByProductId(prod.Id)

	return prod, nil
}

func (r ProductDBRepository) GetByName(n string) (products []models.Product, err error) {
	var prod models.Product

	rows, err := r.DB.Query("SELECT id, menu_id, name, price, image, type FROM products WHERE name = (?)", n)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&prod.Id, &prod.MenuId, &prod.Name, &prod.Price, &prod.Image, &prod.Type)
		if err != nil {
			panic(err)
		}
		products = append(products, prod)
	}
	IngredientRepo := NewIngredientRepository(r.DB)

	for i := range products {
		products[i].Ingredients, err = IngredientRepo.GetByProductId(products[i].Id)
	}
	return products, nil
}

func (r ProductDBRepository) GetBySupplier(id int) (products []models.Product, err error) {
	var prod models.Product

	rows, err := r.DB.Query("SELECT id, menu_id, name, price, image, type FROM products "+
		"WHERE menu_id = (SELECT id FROM menus WHERE supplier_id =(?))", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&prod.Id, &prod.MenuId, &prod.Name, &prod.Price, &prod.Image, &prod.Type)
		if err != nil {
			panic(err)
		}
		products = append(products, prod)
	}
	IngredientRepo := NewIngredientRepository(r.DB)

	for i := range products {
		products[i].Ingredients, err = IngredientRepo.GetByProductId(products[i].Id)
	}
	return products, nil
}

func (r ProductDBRepository) GetByType(t string) (products []models.Product, err error) {
	var prod models.Product

	rows, err := r.DB.Query("SELECT id, menu_id, name, price, image, type FROM products WHERE type = (?)", t)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&prod.Id, &prod.MenuId, &prod.Name, &prod.Price, &prod.Image, &prod.Type)
		if err != nil {
			panic(err)
		}
		products = append(products, prod)
	}
	IngredientRepo := NewIngredientRepository(r.DB)

	for i := range products {
		products[i].Ingredients, err = IngredientRepo.GetByProductId(products[i].Id)
	}
	return products, nil
}

func (r ProductDBRepository) DeleteAll() error {
	_, err := r.DB.Exec("DELETE FROM products")
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductDBRepository) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *ProductDBRepository) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.CommitTx()
	}
	return nil
}

func (r *ProductDBRepository) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.RollbackTx()
	}
	return nil
}

func (r *ProductDBRepository) GetTx() *sql.Tx {
	return r.TX
}

func (r *ProductDBRepository) SetTx(tx *sql.Tx) {
	r.TX = tx
}
