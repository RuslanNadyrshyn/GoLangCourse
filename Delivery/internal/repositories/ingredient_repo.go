package repositories

import (
	"database/sql"
	"log"
)

type IIngredientRepository interface {
	Insert(p string) (int, error)
	GetByProductId(ProductId int) ([]string, error)
}

type IngredientRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewIngredientRepository(conn *sql.DB) IIngredientRepository {
	return &IngredientRepository{
		DB: conn,
	}
}

func (r *IngredientRepository) Insert(p string) (int, error) {
	var ingredientId int64
	if r.TX != nil {
		result, err := r.TX.Exec("INSERT IGNORE INTO ingredients(name) VALUES (?);", p)
		if err != nil {
			log.Println(err)
			return 0, err
		}
		ingredientId, err = result.LastInsertId()
		if err != nil {
			log.Println(err)
			return 0, err
		}

		if ingredientId == 0 {
			err = r.DB.QueryRow("SELECT id FROM ingredients WHERE ingredients.name = (?)", p).Scan(&ingredientId)
			if err != nil {
				log.Println(err)
				return 0, err
			}
		}
		return int(ingredientId), err
	}

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

	if ingredientId == 0 {
		err = r.DB.QueryRow("SELECT id FROM ingredients WHERE ingredients.name = (?)", p).Scan(&ingredientId)
		if err != nil {
			log.Println(err)
			return 0, err
		}
	}
	return int(ingredientId), err
}

func (r *IngredientRepository) GetByProductId(ProductId int) (ingredients []string, err error) {
	var ingr string
	rows, err := r.DB.Query("SELECT name FROM ingredients WHERE id IN "+
		"(SELECT ingredient_id FROM product_ingredients WHERE product_id = (?))", ProductId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ingr)
		if err != nil {
			panic(err)
		}
		ingredients = append(ingredients, ingr)
	}
	return ingredients, nil
}
