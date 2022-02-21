package server

import (
	"awesomeProject/internal/auth/config"
	"awesomeProject/internal/auth/repositories"
	"awesomeProject/internal/auth/server/handlers"
	"awesomeProject/internal/auth/services"
	repos "awesomeProject/internal/repositories"
	"awesomeProject/internal/repositories/database/Connection"
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

	suppliers, err := supplierService.SupplierRepo.GetAll()
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
