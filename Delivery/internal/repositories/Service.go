package repositories

import (
	"awesomeProject/internal/repositories/database"
	"awesomeProject/internal/repositories/models"
	"database/sql"
	"fmt"
	"log"
)

type SupplierService struct {
	SupplierRepo          database.SupplierDBRepository
	MenuRepo              database.MenuDBRepository
	ProductRepo           database.ProductDBRepository
	IngredientRepo        database.IngredientDBRepository
	ProductIngredientRepo database.ProductIngredientDBRepository
}

func NewSupplierService(conn *sql.DB) SupplierService {
	return SupplierService{
		SupplierRepo:          database.NewSupplierRepository(conn),
		MenuRepo:              database.NewMenuRepository(conn),
		ProductRepo:           database.NewProductRepository(conn),
		IngredientRepo:        database.NewIngredientRepository(conn),
		ProductIngredientRepo: database.NewProductIngredientRepository(conn),
	}
}

func (ss SupplierService) CreateSupplier(sup models.Supplier) {

	//Supplier TX
	err := ss.SupplierRepo.BeginTx()
	supId, err := ss.SupplierRepo.Insert(sup)
	if err != nil {
		log.Println("SupplierRepo error:", err)
		ss.SupplierRepo.TX.Rollback()
	}
	if err != nil {
		fmt.Println(err)
		ss.SupplierRepo.TX.Rollback()
	}
	ss.SupplierRepo.TX.Commit()

	//Menu TX
	err = ss.MenuRepo.BeginTx()
	menuId, err := ss.MenuRepo.Insert(sup, supId)
	if err != nil {
		log.Println("MenuRepo error:", err)
		ss.SupplierRepo.TX.Rollback()
	}
	if err != nil {
		fmt.Println(err)
		ss.MenuRepo.TX.Rollback()
	}
	ss.MenuRepo.TX.Commit()

	//Products TX
	for i := range sup.Menu {
		err = ss.ProductRepo.BeginTx()
		sup.Menu[i].MenuId = int(menuId)
		productId, err := ss.ProductRepo.Insert(sup.Menu[i])
		if err != nil {
			fmt.Println(err)
			ss.ProductRepo.TX.Rollback()
		}
		ss.ProductRepo.TX.Commit()

		//Ingredients TX
		for j := range sup.Menu[i].Ingredients {
			err = ss.IngredientRepo.BeginTx()
			ingredientId, err := ss.IngredientRepo.Insert(sup.Menu[i].Ingredients[j])
			if err != nil {
				log.Println(err)
				ss.IngredientRepo.TX.Rollback()
			}
			ss.IngredientRepo.TX.Commit()

			//ProductIngredients TX
			err = ss.ProductIngredientRepo.BeginTx()
			_, err = ss.ProductIngredientRepo.Insert(int(productId), ingredientId)
			if err != nil {
				log.Println(err)
				ss.ProductIngredientRepo.TX.Rollback()
			}
			ss.ProductIngredientRepo.TX.Commit()
		}
	}
}
