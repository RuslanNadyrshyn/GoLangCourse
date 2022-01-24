package database

import (
	"awesomeProject/internal/repositories/models"
	"database/sql"
	"log"
)

type ProductDBRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewProductRepository(conn *sql.DB) ProductDBRepository {
	return ProductDBRepository{
		DB: conn,
	}
}

func (r ProductDBRepository) Insert(p models.Product) (int, error) {
	var productId int64

	if r.TX != nil {
		ingredientRepo := NewIngredientRepository(r.DB)
		productIngredientRepo := NewProductIngredientRepository(r.DB)
		ingredientRepo.TX = r.TX
		productIngredientRepo.TX = r.TX

		result, err := r.TX.Exec("INSERT products(menu_id, name, type, price, image) VALUES(?, ?, ?, ?, ?)",
			p.MenuId, p.Name, p.Type, p.Price, p.Image)
		if err != nil {
			log.Println(err)
			return int(productId), err
		}
		productId, err = result.LastInsertId()

		for i := range p.Ingredients {
			ingredientId, err := ingredientRepo.Insert(p.Ingredients[i])
			if err != nil {
				log.Println(err)
				return int(productId), err
			}
			if ingredientId != 0 {
				_, err = productIngredientRepo.Insert(int(productId), ingredientId)
				if err != nil {
					log.Println(err)
					return int(productId), err
				}
			}
		}
		return int(productId), err
	}
	return 0, nil
}
