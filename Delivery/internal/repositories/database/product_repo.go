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
	var productId int
	var ingredientId int
	var productIngredientId int

	if r.TX != nil {
		err := r.TX.QueryRow("INSERT products(id, menu_id, name, type, price, image) VALUES(?, ?, ?, ?, ?, ?) RETURNING id",
			p.Id, p.MenuId, p.Name, p.Type, p.Price, p.Image).Scan(&productId)
		if err != nil {
			_ = r.TX.Rollback()
		}

		for i := range p.Ingredients {
			err = r.DB.QueryRow("INSERT IGNORE INTO ingredients(name) values (?) RETURNING id",
				p.Ingredients[i]).Scan(&ingredientId)
			if err != nil {
				_ = r.TX.Rollback()
			}

			err = r.DB.QueryRow("INSERT INTO product_ingredients(product_id, ingredient_id)"+
				"VALUES (?, (SELECT id FROM ingredients WHERE name=(?)))", p.Id, p.Ingredients[i]).Scan(&productIngredientId)
			if err != nil {
				_ = r.TX.Rollback()
			}
		}

		err = r.TX.Commit()
		if err != nil {
			_ = r.TX.Rollback()
		}
		return productId, err
	}

	_, err := r.DB.Exec("INSERT products(id, menu_id, name, type, price, image) VALUES(?, ?, ?, ?, ?, ?)",
		p.Id, p.MenuId, p.Name, p.Type, p.Price, p.Image)
	if err != nil {
		log.Panic(err)
		return productId, err
	}

	for i := range p.Ingredients {
		_, err = r.DB.Exec("INSERT IGNORE INTO ingredients(name) values (?)",
			p.Ingredients[i])
		if err != nil {
			log.Panic(err)
			return productId, err
		}

		_, err = r.DB.Exec("INSERT INTO product_ingredients(product_id, ingredient_id)"+
			"VALUES (?, (SELECT id FROM ingredients WHERE name=(?)))", p.Id, p.Ingredients[i])
		if err != nil {
			log.Panic(err)
			return productId, err
		}
	}
	if err != nil {
		log.Panic(err)
		return productId, err
	}

	return productId, err
}
