package server

import (
	"Delivery/Delivery/cfg"
	"Delivery/Delivery/internal/handlers"
	"Delivery/Delivery/internal/repositories/database/Connection"
	"Delivery/Delivery/internal/services"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func Start(cfg *cfg.Config) {
	conn, err := Connection.Connect()
	if err != nil {
		log.Println(err)
	}
	rs := services.NewRepositoryService(conn)
	tokenService := services.NewTokenService(cfg)
	//mw := middleware.NewMiddleware(tokenService)

	authHandler := handlers.NewAuthHandler(cfg, rs.UserRepo)
	userHandler := handlers.NewUserHandler(tokenService, rs.UserRepo)
	supplierHandler := handlers.NewSupplierHandler(
		rs.SupplierRepo, rs.MenuRepo, rs.ProductRepo,
		rs.ProductIngredientRepo, rs.IngredientRepo,
	)
	orderHandler := handlers.NewOrderHandler(tokenService, rs.UserRepo, rs.OrderRepo, rs.OrderProductRepo, rs.ProductRepo)
	productHandler := handlers.NewProductHandler(rs.ProductRepo, rs.MenuRepo, rs.SupplierRepo)

	mux := http.NewServeMux()
	//auth
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/refresh", authHandler.Refresh)
	//user
	mux.HandleFunc("/sign_in", userHandler.SignIn)
	mux.HandleFunc("/profile", userHandler.GetProfile)
	//mux.Handle("/profile", mw.IsAuth(userHandler.GetProfile))
	//supplier
	mux.HandleFunc("/get_suppliers", supplierHandler.GetAll)
	//products
	mux.HandleFunc("/get_products", productHandler.GetAll)
	mux.HandleFunc("/prod", productHandler.GetById)
	//order
	mux.HandleFunc("/get_order", orderHandler.GetById)
	mux.HandleFunc("/post_order", orderHandler.Add)
	mux.HandleFunc("/orders", orderHandler.GetByUserId)

	log.Fatal(http.ListenAndServe(cfg.Port, mux))
}
