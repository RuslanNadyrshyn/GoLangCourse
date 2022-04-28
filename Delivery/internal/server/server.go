package server

import (
	"Delivery/Delivery/cfg"
	repos "Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/database/Connection"
	handlers2 "Delivery/Delivery/internal/server/handlers"
	"Delivery/Delivery/internal/services"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var (
	conn, err       = Connection.Connect()
	supplierService = services.NewSupplierService(conn)
)

func Start(cfg *cfg.Config) {
	userRepository := repos.NewUserRepository()
	supplierRepository := repos.NewSupplierRepository()
	tokenService := services.NewTokenService(cfg)
	if err != nil {
		log.Println(err)
	}

	authHandler := handlers2.NewAuthHandler(cfg, conn)
	userHandler := handlers2.NewUserHandler(tokenService, userRepository, conn)
	supplierHandler := handlers2.NewSupplierHandler(&supplierService, supplierRepository, conn)

	mux := http.NewServeMux()
	mux.HandleFunc("/sign_in", userHandler.SignIn)
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/profile", userHandler.GetProfile)
	mux.HandleFunc("/refresh", authHandler.Refresh)

	mux.HandleFunc("/get_suppliers", supplierHandler.GetAll)

	// Надо перенести в хендлеры
	mux.HandleFunc("/get_products", GetProducts)

	mux.HandleFunc("/get_order", supplierService.GetOrder)
	mux.HandleFunc("/post_order", supplierService.CreateOrder)

	log.Fatal(http.ListenAndServe(cfg.Port, mux))
}

func GetSuppliers(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resp, "not allowed", http.StatusMethodNotAllowed)
	}

	suppliers, err := supplierService.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	resp.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(resp).Encode(suppliers)
}

func GetProducts(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resp, "not allowed", http.StatusMethodNotAllowed)
	}

	products, err := supplierService.ProductRepo.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	resp.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(resp).Encode(products)
}

//func GetBasket(resp http.ResponseWriter, req *http.Request) {
//	if req.Method != http.MethodGet {
//		http.Error(resp, "not allowed", http.StatusMethodNotAllowed)
//	}
//	//
//	//basket, err := supplierService.BasketRepo.GetById(id)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//
//	//data, _ := json.Marshal(basket)
//	//resp.Header().Add("Access-Control-Allow-Origin", "*")
//	//fmt.Fprintln(resp, string(data))
//}
