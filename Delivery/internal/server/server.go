package server

import (
	"Delivery/Delivery/cfg"
	"Delivery/Delivery/internal/handlers"
	"Delivery/Delivery/internal/repositories"
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

	store := repositories.NewStore(conn)
	service, err := services.NewServiceManager(store, cfg)
	h := handlers.NewController(service, cfg)
	//mw := middleware.NewMiddleware(service)

	mux := http.NewServeMux()
	//user
	mux.HandleFunc("/sign_in", h.User.SignIn)
	mux.HandleFunc("/profile", h.User.GetProfile)
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
	mux.HandleFunc("/orders", h.Order.GetByUserId)

	//mux.Handle("/profile", mw.IsAuth(h.User.GetProfile))
	//mux.Handle("/profile", mw.AuthCheck(http.HandlerFunc(h.User.GetProfile)))

	log.Fatal(http.ListenAndServe(cfg.Port, mux))
}
