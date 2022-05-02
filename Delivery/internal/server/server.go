package server

import (
	"Delivery/Delivery/cfg"
	repos "Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/database/Connection"
	handlers2 "Delivery/Delivery/internal/server/handlers"
	"Delivery/Delivery/internal/services"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var (
	conn, err = Connection.Connect()
)

func Start(cfg *cfg.Config) {
	userRepository := repos.NewUserRepository()
	supplierRepository := repos.NewSupplierRepository()
	orderRepository := repos.NewOrderRepository()
	productRepository := repos.NewProductRepository()
	tokenService := services.NewTokenService(cfg)
	if err != nil {
		log.Println(err)
	}

	authHandler := handlers2.NewAuthHandler(cfg, conn)
	userHandler := handlers2.NewUserHandler(tokenService, userRepository, conn)
	supplierHandler := handlers2.NewSupplierHandler(supplierRepository, conn)
	orderHandler := handlers2.NewOrderHandler(orderRepository, conn)
	productHandler := handlers2.NewProductHandler(productRepository, conn)

	mux := http.NewServeMux()
	mux.HandleFunc("/sign_in", userHandler.SignIn)
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/profile", userHandler.GetProfile)
	mux.HandleFunc("/refresh", authHandler.Refresh)

	mux.HandleFunc("/get_suppliers", supplierHandler.GetAll)

	mux.HandleFunc("/get_products", productHandler.GetAll)
	mux.HandleFunc("/prod", productHandler.GetById)

	mux.HandleFunc("/get_order", orderHandler.GetById)
	mux.HandleFunc("/post_order", orderHandler.Add)

	mux.HandleFunc("/orders", orderHandler.GetByUserId)

	log.Fatal(http.ListenAndServe(cfg.Port, mux))
}

//func GetBasket(resp parser.ResponseWriter, req *parser.Request) {
//	if req.Method != parser.MethodGet {
//		parser.Error(resp, "not allowed", parser.StatusMethodNotAllowed)
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
