package repositories

import (
	"database/sql"
	"log"
)

type IIngredientRepository interface {
	Insert(p string) (int64, error)
	GetByProductId(ProductId int64) ([]string, error)
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

func (r *IngredientRepository) Insert(p string) (ingredientId int64, err error) {
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
		return ingredientId, nil
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
	return ingredientId, nil
}

func (r *IngredientRepository) GetByProductId(ProductId int64) (ingredients []string, err error) {
	var ingredient string
	rows, err := r.DB.Query("SELECT name FROM ingredients WHERE id IN "+
		"(SELECT ingredient_id FROM product_ingredients WHERE product_id = (?))", ProductId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ingredient)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}
