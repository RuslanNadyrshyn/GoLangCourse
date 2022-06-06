package server

import (
	"Delivery/backend/cfg"
	"Delivery/backend/internal/handlers"
	"Delivery/backend/internal/middleware"
	"Delivery/backend/internal/repositories"
	"Delivery/backend/internal/repositories/database/Connection"
	"Delivery/backend/internal/services"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func Start(cfg *cfg.Config) {
	conn, err := Connection.Connect()
	if err != nil {
		log.Println(err)
	}

	store := repositories.NewStore(conn)
	service, err := services.NewServiceManager(store, cfg)
	h := handlers.NewController(service, cfg)
	mw := middleware.NewMiddleware(service)

	mux := http.NewServeMux()
	//user
	mux.HandleFunc("/sign_up", h.User.SignUp)
	mux.Handle("/profile", mw.Validation(http.HandlerFunc(h.User.GetProfile)))
	//supplier
	mux.HandleFunc("/get_suppliers", h.Supplier.GetAll)
	//products
	mux.HandleFunc("/get_products", h.Product.GetAll)
	mux.HandleFunc("/prod", h.Product.GetById)
	//auth
	mux.HandleFunc("/login", h.Auth.Login)
	mux.HandleFunc("/refresh", h.Auth.Refresh)
	//order
	mux.HandleFunc("/get_order", h.Order.GetById)
	mux.HandleFunc("/post_order", h.Order.Add)
	mux.Handle("/orders", mw.Validation(http.HandlerFunc(h.Order.GetByUserId)))

	log.Fatal(http.ListenAndServe(cfg.Port, mux))
}
