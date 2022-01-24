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
	return 0, nil
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
