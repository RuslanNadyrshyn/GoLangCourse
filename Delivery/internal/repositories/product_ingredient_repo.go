package repositories

import (
	"database/sql"
	"log"
)

type IProductIngredientRepository interface {
	Insert(productId int, ingredientId int) (int, error)
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

func (r *ProductIngredientRepository) Insert(productId int, ingredientId int) (int, error) {
	var ProductIngredientId int64
	if r.TX != nil {
		result, err := r.TX.Exec("INSERT INTO product_ingredients(product_id, ingredient_id)"+
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
