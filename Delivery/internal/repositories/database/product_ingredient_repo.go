package database

import (
	"database/sql"
	"log"
)

type ProductIngredientDBRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewProductIngredientRepository(conn *sql.DB) ProductIngredientDBRepository {
	return ProductIngredientDBRepository{
		DB: conn,
	}
}

func (r ProductIngredientDBRepository) Insert(productId int, ingredientId int) (int, error) {
	var ProductIngredientId int64
	if r.TX != nil {
		result, err := r.DB.Exec("INSERT INTO product_ingredients(product_id, ingredient_id)"+
			"VALUES (?, ?)", productId, ingredientId)
		if err != nil {
			log.Println(err)
			return int(ProductIngredientId), err
		}
		ProductIngredientId, err = result.LastInsertId()
		if err != nil {
			log.Println(err)
			return int(ProductIngredientId), err
		}
		return int(ProductIngredientId), err
	}
	return 0, nil
}
func (r *ProductIngredientDBRepository) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *ProductIngredientDBRepository) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.CommitTx()
	}
	return nil
}

func (r *ProductIngredientDBRepository) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.RollbackTx()
	}
	return nil
}

func (r *ProductIngredientDBRepository) GetTx() *sql.Tx {
	return r.TX
}

func (r *ProductIngredientDBRepository) SetTx(tx *sql.Tx) {
	r.TX = tx
}
