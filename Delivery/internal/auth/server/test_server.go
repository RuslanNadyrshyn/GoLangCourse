package server

import (
	"Delivery/Delivery/internal/auth/config"
	"Delivery/Delivery/internal/auth/repositories"
	"Delivery/Delivery/internal/auth/server/handlers"
	"Delivery/Delivery/internal/auth/services"
	repos "Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/database/Connection"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var (
	conn, err       = Connection.Connect()
	supplierService = repos.NewSupplierService(conn)
)

func Start(cfg *config.Config) {
	userRepository := repositories.NewUserRepository()
	tokenService := services.NewTokenService(cfg)
	if err != nil {
		log.Println(err)
	}

	authHandler := handlers.NewAuthHandler(cfg, conn)
	userHandler := handlers.NewUserHandler(tokenService, userRepository, conn)

	mux := http.NewServeMux()
	mux.HandleFunc("/sign_in", userHandler.SignIn)
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/profile", userHandler.GetProfile)
	mux.HandleFunc("/refresh", authHandler.Refresh)

	mux.HandleFunc("/get_products", GetProducts)
	mux.HandleFunc("/get_suppliers", GetSuppliers)

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

	data, _ := json.Marshal(suppliers)
	resp.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(resp, string(data))
}

func GetProducts(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resp, "not allowed", http.StatusMethodNotAllowed)
	}

	products, err := supplierService.ProductRepo.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	data, _ := json.Marshal(products)
	resp.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(resp, string(data))
}

func GetBasket(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resp, "not allowed", http.StatusMethodNotAllowed)
	}
	//
	//basket, err := supplierService.BasketRepo.GetById(id)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//data, _ := json.Marshal(basket)
	//resp.Header().Add("Access-Control-Allow-Origin", "*")
	//fmt.Fprintln(resp, string(data))
}
