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
	if ingredientId != 0 {
		result, err := r.DB.Exec("INSERT INTO product_ingredients(product_id, ingredient_id)"+
			"VALUES (?, (SELECT id FROM ingredients WHERE id=(?)))", productId, ingredientId)
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
