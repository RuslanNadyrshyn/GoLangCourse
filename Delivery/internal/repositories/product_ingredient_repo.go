package repositories

import (
	"database/sql"
	"log"
)

type IProductIngredientRepository interface {
	Insert(productId int64, ingredientId int64) (int64, error)
}
type ProductIngredientRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewProductIngredientRepository(conn *sql.DB) IProductIngredientRepository {
	return &ProductIngredientRepository{
		DB: conn,
	}
}

func (r *ProductIngredientRepository) Insert(productId int64, ingredientId int64) (int64, error) {
	var ProductIngredientId int64
	if r.TX != nil {
		result, err := r.TX.Exec("INSERT INTO product_ingredients(product_id, ingredient_id)"+
			"VALUES (?, ?)", productId, ingredientId)
		if err != nil {
			log.Println(err)
			return 0, err
		}
		ProductIngredientId, err = result.LastInsertId()
		if err != nil {
			log.Println(err)
			return 0, err
		}
		return ProductIngredientId, nil
	}
	result, err := r.DB.Exec("INSERT INTO product_ingredients(product_id, ingredient_id)"+
		"VALUES (?, ?)", productId, ingredientId)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	ProductIngredientId, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return ProductIngredientId, nil
}
