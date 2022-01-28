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

func (r IngredientDBRepository) GetByProductId(ProductId int) (ingredients []string, err error) {
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

func (r IngredientDBRepository) DeleteAll() error {
	_, err := r.DB.Exec("DELETE FROM ingredients")
	if err != nil {
		return err
	}
	return nil
}

func (r IngredientDBRepository) DeleteByName(name string) error {
	_, err := r.DB.Exec("DELETE FROM ingredients WHERE name = ?", name)
	if err != nil {
		return err
	}
	return nil
}

func (r *IngredientDBRepository) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *IngredientDBRepository) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.CommitTx()
	}
	return nil
}

func (r *IngredientDBRepository) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.RollbackTx()
	}
	return nil
}

func (r *IngredientDBRepository) GetTx() *sql.Tx {
	return r.TX
}

func (r *IngredientDBRepository) SetTx(tx *sql.Tx) {
	r.TX = tx
}
