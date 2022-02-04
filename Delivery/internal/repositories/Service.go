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

func (ss SupplierService) CreateSupplierTX(sup models.Supplier) {
	//Supplier TX
	err := ss.SupplierRepo.BeginTx()
	sup.Id, err = ss.SupplierRepo.Insert(sup)
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
	menuId, err := ss.MenuRepo.Insert(sup.Id)
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
		sup.Menu[i].MenuId = menuId
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

func (ss SupplierService) CreateSupplier(sup models.Supplier) models.Supplier {
	//Supplier
	var err error
	sup.Id, err = ss.SupplierRepo.Insert(sup)
	if err != nil {
		log.Println(err)
		return sup
	}

	//Menu
	menuId, err := ss.MenuRepo.Insert(sup.Id)
	if err != nil {
		log.Println(err)
		return sup
	}

	//Products
	for i := range sup.Menu {
		sup.Menu[i].MenuId = menuId
		sup.Menu[i].Id, err = ss.ProductRepo.Insert(sup.Menu[i])
		if err != nil {
			fmt.Println(err)
			return sup
		}

		//Ingredients
		for j := range sup.Menu[i].Ingredients {
			ingredientId, err := ss.IngredientRepo.Insert(sup.Menu[i].Ingredients[j])
			if err != nil {
				log.Println(err)
				return sup
			}

			//ProductIngredients
			_, err = ss.ProductIngredientRepo.Insert(sup.Menu[i].Id, ingredientId)
			if err != nil {
				log.Println(err)
				return sup
			}
		}
	}
	return sup
}

func (ss SupplierService) DeleteAll() error {
	err := ss.ProductIngredientRepo.DeleteAll()
	if err != nil {
		log.Println(err)
		return err
	}
	err = ss.IngredientRepo.DeleteAll()
	if err != nil {
		log.Println(err)
		return err
	}
	err = ss.ProductRepo.DeleteAll()
	if err != nil {
		log.Println(err)
		return err
	}
	err = ss.MenuRepo.DeleteAll()
	if err != nil {
		log.Println(err)
		return err
	}
	err = ss.SupplierRepo.DeleteAll()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ss SupplierService) CreateMenu(sup models.Supplier) {
	//Menu
	menuId, err := ss.MenuRepo.Insert(sup.Id)
	if err != nil {
		log.Println("MenuRepo error:", err)
		return
	}

	//Products TX
	for i := range sup.Menu {
		sup.Menu[i].MenuId = menuId
		productId, err := ss.ProductRepo.Insert(sup.Menu[i])
		if err != nil {
			fmt.Println(err)
			return
		}

		//Ingredients TX
		for j := range sup.Menu[i].Ingredients {
			ingredientId, err := ss.IngredientRepo.Insert(sup.Menu[i].Ingredients[j])
			if err != nil {
				log.Println(err)
				return
			}

			//ProductIngredients TX
			_, err = ss.ProductIngredientRepo.Insert(int(productId), ingredientId)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
