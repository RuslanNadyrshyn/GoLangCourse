package database

import (
	"database/sql"
	"log"
)

type IngredientDBRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewIngredientRepository(conn *sql.DB) IngredientDBRepository {
	return IngredientDBRepository{
		DB: conn,
	}
}

func (r IngredientDBRepository) Insert(p string) (int, error) {
	var ingredientId int64

	result, err := r.DB.Exec("INSERT IGNORE INTO ingredients(name) VALUES (?);", p)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	ingredientId, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return int(ingredientId), err
}
