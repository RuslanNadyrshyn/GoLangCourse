package services

import (
	"Delivery/Delivery/internal/repositories/database"
	"Delivery/Delivery/internal/repositories/models"
	"Delivery/Delivery/internal/repositories/requests"
	"Delivery/Delivery/internal/repositories/responses"
	"database/sql"
	"fmt"
	"log"
)

type DBService struct {
	UserRepo              database.UserDBRepository
	SupplierRepo          database.SupplierDBRepository
	MenuRepo              database.MenuDBRepository
	ProductRepo           database.ProductDBRepository
	IngredientRepo        database.IngredientDBRepository
	ProductIngredientRepo database.ProductIngredientDBRepository
	OrderRepo             database.OrderDBRepository
	OrderProductRepo      database.OrderProductDBRepository
}

func NewDBService(conn *sql.DB) DBService {
	return DBService{
		UserRepo:              database.NewUserRepository(conn),
		SupplierRepo:          database.NewSupplierRepository(conn),
		MenuRepo:              database.NewMenuRepository(conn),
		ProductRepo:           database.NewProductRepository(conn),
		IngredientRepo:        database.NewIngredientRepository(conn),
		ProductIngredientRepo: database.NewProductIngredientRepository(conn),
		OrderRepo:             database.NewOrderRepository(conn),
		OrderProductRepo:      database.NewOrderProductRepository(conn),
	}
}

func (dbs *DBService) GetAll() (suppliers []requests.SupplierRequest, err error) {
	var sup requests.SupplierRequest
	rows, err := dbs.SupplierRepo.DB.Query("SELECT id, name, type, image, opening, closing FROM suppliers")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&sup.Id, &sup.Name, &sup.Type, &sup.Image,
			&sup.WorkingHours.Opening, &sup.WorkingHours.Closing)
		if err != nil {
			panic(err)
		}
		sup.Menu, err = dbs.ProductRepo.GetBySupplierId(sup.Id)
		suppliers = append(suppliers, sup)
	}

	return suppliers, nil
}

func (dbs *DBService) CreateSupplier(sup *requests.SupplierRequest) int {
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
	sup.Id, err = dbs.SupplierRepo.Insert(&supplier)
	if err != nil {
		log.Println(err)
		return 0
	}

	//Menu
	menuId, err := dbs.MenuRepo.Insert(sup.Id)
	if err != nil {
		log.Println(err)
		return 0
	}

	//Products
	for i := range sup.Menu {
		sup.Menu[i].MenuId = menuId
		sup.Menu[i].Id, err = dbs.ProductRepo.Insert(sup.Menu[i])
		if err != nil {
			fmt.Println(err)
			return 0
		}

		//Ingredients
		for j := range sup.Menu[i].Ingredients {
			ingredientId, err := dbs.IngredientRepo.Insert(sup.Menu[i].Ingredients[j])
			if err != nil {
				log.Println(err)
				return 0
			}

			//ProductIngredients
			_, err = dbs.ProductIngredientRepo.Insert(sup.Menu[i].Id, ingredientId)
			if err != nil {
				log.Println(err)
				return 0
			}
		}
	}
	return supplier.Id
}

func (dbs *DBService) DeleteAll() error {
	err := dbs.ProductIngredientRepo.DeleteAll()
	if err != nil {
		log.Println(err)
		return err
	}
	err = dbs.IngredientRepo.DeleteAll()
	if err != nil {
		log.Println(err)
		return err
	}
	err = dbs.ProductRepo.DeleteAll()
	if err != nil {
		log.Println(err)
		return err
	}
	err = dbs.MenuRepo.DeleteAll()
	if err != nil {
		log.Println(err)
		return err
	}
	err = dbs.SupplierRepo.DeleteAll()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (dbs *DBService) AddOrder(req *requests.OrderRequest) (orderId int, err error) {
	userId, err := dbs.UserRepo.Insert(req.User)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	orderId, err = dbs.OrderRepo.Insert(
		&models.Order{
			Price:   req.TotalPrice,
			UserId:  userId,
			Address: req.Address,
		})
	if err != nil {
		log.Println(err)
		return 0, err
	}
	for _, product := range req.Products {
		_, err := dbs.OrderProductRepo.Insert(orderId, product.ProductId, product.Counter, product.ProductPrice)
		if err != nil {
			log.Println(err)
			return 0, err
		}
	}
	return orderId, nil
}

func (dbs *DBService) GetOrderById(orderId int) (order *responses.OrderResponse, err error) {
	o, err := dbs.OrderRepo.GetById(orderId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	orderProducts, err := dbs.OrderProductRepo.GetByOrderId(o.Id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var respProducts []responses.Product
	for _, orderProd := range orderProducts {
		product, err := dbs.ProductRepo.GetById(orderProd.ProductId)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		prod := responses.Product{
			Id:          product.Id,
			MenuId:      product.MenuId,
			Name:        product.Name,
			Price:       product.Price,
			Image:       product.Image,
			Type:        product.Type,
			Ingredients: product.Ingredients,
			Counter:     orderProd.Count,
			OldPrice:    orderProd.Price,
		}
		respProducts = append(respProducts, prod)
	}

	order = &responses.OrderResponse{
		Id:       o.Id,
		UserId:   o.UserId,
		Address:  o.Address,
		Price:    o.Price,
		Products: respProducts,
	}
	return order, nil
}

func (dbs *DBService) GetProductById(productId int) (product *responses.ProductResponse, err error) {
	prod, err := dbs.ProductRepo.GetById(productId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	supplierId, err := dbs.MenuRepo.GetById(prod.MenuId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	supplier, err := dbs.SupplierRepo.GetById(supplierId)
	product = &responses.ProductResponse{
		Product:  prod,
		Supplier: supplier,
	}

	return product, nil
}
