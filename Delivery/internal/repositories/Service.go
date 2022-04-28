package repositories

import (
	"Delivery/Delivery/internal/auth/requests"
	"Delivery/Delivery/internal/auth/responses"
	"Delivery/Delivery/internal/repositories/database"
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type SupplierService struct {
	UserRepo              database.UserDBRepository
	SupplierRepo          database.SupplierDBRepository
	MenuRepo              database.MenuDBRepository
	ProductRepo           database.ProductDBRepository
	IngredientRepo        database.IngredientDBRepository
	ProductIngredientRepo database.ProductIngredientDBRepository
	OrderRepo             database.OrderDBRepository
	OrderProductRepo      database.OrderProductDBRepository
}

func NewSupplierService(conn *sql.DB) SupplierService {
	return SupplierService{
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

func (ss SupplierService) GetAll() (suppliers []models.Supplier, err error) {
	var sup models.Supplier
	rows, err := ss.SupplierRepo.DB.Query("SELECT id, name, type, image, opening, closing FROM suppliers")
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
		sup.Menu, err = ss.ProductRepo.GetBySupplier(sup.Id)
		suppliers = append(suppliers, sup)
	}
	fmt.Println(suppliers)
	return suppliers, nil
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

	//Products
	for i := range sup.Menu {
		sup.Menu[i].MenuId = menuId
		productId, err := ss.ProductRepo.Insert(sup.Menu[i])
		if err != nil {
			fmt.Println(err)
			return
		}

		//Ingredients
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

func (ss SupplierService) CreateOrder(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode("OKAY")
		if err != nil {
			fmt.Println(err)
			return
		}
	case "POST":
		var userId, orderId int
		req := new(requests.OrderRequest)
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		userId, err = ss.UserRepo.Insert(req.User)
		if err != nil {
			log.Println(err)
			return
		}
		order := models.Order{
			Price:   req.TotalPrice,
			UserId:  userId,
			Address: req.Address,
		}
		orderId, err = ss.OrderRepo.Insert(&order)
		if err != nil {
			log.Println(err)
			return
		}
		for _, product := range req.Products {
			_, err := ss.OrderProductRepo.Insert(orderId, product.ProductId, product.Counter, product.ProductPrice)
			if err != nil {
				log.Println(err)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(orderId)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func (ss *SupplierService) GetOrder(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "GET":
		req := r.URL.Query().Get("id")
		id, err := strconv.Atoi(req)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		order, err := ss.OrderRepo.GetById(id)
		if err != nil {
			http.Error(w, "Order does not exist", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		fmt.Printf("Id: %d\nPrice: %f\nAddress: %s\n", order.Id, order.Price, order.Address)

		orderProducts, err := ss.OrderProductRepo.GetByOrderId(order.Id)
		if err != nil {
			http.Error(w, "OrderProduct error", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		var respProducts []responses.Product
		for _, orderProd := range orderProducts {
			product, err := ss.ProductRepo.GetById(orderProd.ProductId)
			if err != nil {
				http.Error(w, "Product error", http.StatusBadRequest)
				fmt.Println(err)
				return
			}
			p := responses.Product{
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
			respProducts = append(respProducts, p)
		}

		resp := responses.OrderResponse{
			Id:       order.Id,
			UserId:   order.UserId,
			Address:  order.Address,
			Price:    order.Price,
			Products: respProducts,
		}

		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
