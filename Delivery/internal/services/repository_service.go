package services

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/models"
	"Delivery/Delivery/internal/repositories/requests"
	"database/sql"
	"fmt"
	"log"
)

type RepositoryService struct {
	UserRepo              repositories.IUserRepository
	SupplierRepo          repositories.ISupplierRepository
	MenuRepo              repositories.IMenuRepository
	ProductRepo           repositories.IProductRepository
	IngredientRepo        repositories.IIngredientRepository
	ProductIngredientRepo repositories.IProductIngredientRepository
	OrderRepo             repositories.IOrderRepository
	OrderProductRepo      repositories.IOrderProductRepository
}

func NewRepositoryService(conn *sql.DB) RepositoryService {
	return RepositoryService{
		UserRepo:              repositories.NewUserRepository(conn),
		SupplierRepo:          repositories.NewSupplierRepository(conn),
		MenuRepo:              repositories.NewMenuRepository(conn),
		ProductRepo:           repositories.NewProductRepository(conn),
		IngredientRepo:        repositories.NewIngredientRepository(conn),
		ProductIngredientRepo: repositories.NewProductIngredientRepository(conn),
		OrderRepo:             repositories.NewOrderRepository(conn),
		OrderProductRepo:      repositories.NewOrderProductRepository(conn),
	}
}

func (rs *RepositoryService) CreateSupplier(sup *requests.SupplierRequest) int {
	//Supplier

	supplier := models.Supplier{
		Id:    sup.Id,
		Name:  sup.Name,
		Type:  sup.Type,
		Image: sup.Image,
		WorkingHours: struct {
			Opening string `json:"opening"`
			Closing string `json:"closing"`
		}{
			Opening: sup.WorkingHours.Opening,
			Closing: sup.WorkingHours.Closing,
		},
	}
	var err error
	sup.Id, err = rs.SupplierRepo.Insert(&supplier)
	if err != nil {
		log.Println(err)
		return 0
	}

	//Menu
	menuId, err := rs.MenuRepo.Insert(sup.Id)
	if err != nil {
		log.Println(err)
		return 0
	}

	//Products
	for i := range sup.Menu {
		sup.Menu[i].MenuId = menuId
		sup.Menu[i].Id, err = rs.ProductRepo.Insert(sup.Menu[i])
		if err != nil {
			fmt.Println(err)
			return 0
		}

		//Ingredients
		for j := range sup.Menu[i].Ingredients {
			ingredientId, err := rs.IngredientRepo.Insert(sup.Menu[i].Ingredients[j])
			if err != nil {
				log.Println(err)
				return 0
			}

			//ProductIngredients
			_, err = rs.ProductIngredientRepo.Insert(sup.Menu[i].Id, ingredientId)
			if err != nil {
				log.Println(err)
				return 0
			}
		}
	}
	return supplier.Id
}
