package server

import (
	"Delivery/Delivery/cfg"
	"Delivery/Delivery/internal/repositories/database/Connection"
	h "Delivery/Delivery/internal/server/handlers"
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

	authHandler := h.NewAuthHandler(cfg, rs.UserRepo)
	userHandler := h.NewUserHandler(tokenService, rs.UserRepo)
	supplierHandler := h.NewSupplierHandler(
		rs.SupplierRepo, rs.MenuRepo, rs.ProductRepo,
		rs.ProductIngredientRepo, rs.IngredientRepo,
	)
	orderHandler := h.NewOrderHandler(rs.UserRepo, rs.OrderRepo, rs.OrderProductRepo, rs.ProductRepo)
	productHandler := h.NewProductHandler(rs.ProductRepo, rs.MenuRepo, rs.SupplierRepo)

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
