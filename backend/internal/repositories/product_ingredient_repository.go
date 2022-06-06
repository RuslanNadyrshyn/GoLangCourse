package repositories

import (
	"database/sql"
)

type ProductIngredientRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewProductIngredientRepo(conn *sql.DB) *ProductIngredientRepo {
	return &ProductIngredientRepo{
		DB: conn,
	}
}

func (r *ProductIngredientRepo) Insert(productId int64, ingredientId int64) (int64, error) {
	var ProductIngredientId int64
	if r.TX != nil {
		result, err := r.TX.Exec("INSERT INTO product_ingredients(product_id, ingredient_id)"+
			" VALUES(?, ?)", productId, ingredientId)
		if err != nil {
			return 0, err
		}
		ProductIngredientId, err = result.LastInsertId()
		if err != nil {
			return 0, err
		}
		return ProductIngredientId, nil
	}
	result, err := r.DB.Exec("INSERT INTO product_ingredients(product_id, ingredient_id) "+
		"VALUES(?, ?)", productId, ingredientId)
	if err != nil {
		return 0, err
	}
	ProductIngredientId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return ProductIngredientId, nil
}

func (r *ProductIngredientRepo) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *ProductIngredientRepo) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Commit()
	}
	return nil
}

func (r *ProductIngredientRepo) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Rollback()
	}
	return nil
}

func (r *ProductIngredientRepo) GetTx() *sql.Tx {
	return r.TX
}

func (r *ProductIngredientRepo) SetTx(tx *sql.Tx) {
	r.TX = tx
}
