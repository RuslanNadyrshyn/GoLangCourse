package repositories

import (
	"database/sql"
)

type IngredientRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewIngredientRepo(conn *sql.DB) *IngredientRepo {
	return &IngredientRepo{
		DB: conn,
	}
}

func (r *IngredientRepo) Insert(name string) (ingredientId int64, err error) {
	if r.TX != nil {
		result, err := r.TX.Exec("INSERT IGNORE INTO ingredients(name) VALUES (?);", name)
		if err != nil {
			return 0, err
		}
		ingredientId, err = result.LastInsertId()
		if err != nil {
			return 0, err
		}

		if ingredientId == 0 {
			err = r.TX.QueryRow("SELECT id FROM ingredients WHERE ingredients.name = (?)", name).Scan(&ingredientId)
			if err != nil {
				return 0, err
			}
		}
		return ingredientId, nil
	}
	result, err := r.DB.Exec("INSERT IGNORE INTO ingredients(name) VALUES (?);", name)
	if err != nil {
		return 0, err
	}
	ingredientId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	if ingredientId == 0 {
		err = r.DB.QueryRow("SELECT id FROM ingredients WHERE ingredients.name = (?)", name).Scan(&ingredientId)
		if err != nil {
			return 0, err
		}
	}
	return ingredientId, nil
}

func (r *IngredientRepo) GetById(id int64) (ingredient string, err error) {
	err = r.DB.QueryRow("SELECT name FROM ingredients WHERE id = (?)", id).Scan(ingredient)
	if err != nil {
		return "", err
	}
	return ingredient, nil
}

func (r *IngredientRepo) GetByProductId(ProductId int64) ([]string, error) {
	var ingredients []string
	var ingredient string
	rows, err := r.DB.Query("SELECT name FROM ingredients WHERE id IN (SELECT ingredient_id FROM product_ingredients WHERE product_id = (?))", ProductId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ingredient)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}

func (r *IngredientRepo) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *IngredientRepo) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.CommitTx()
	}
	return nil
}

func (r *IngredientRepo) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.RollbackTx()
	}
	return nil
}

func (r *IngredientRepo) GetTx() *sql.Tx {
	return r.TX
}

func (r *IngredientRepo) SetTx(tx *sql.Tx) {
	r.TX = tx
}
