package server

import (
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/cfg"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/handlers"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/middleware"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories/database/Connection"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/services"
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
	mux.HandleFunc("/get_sup_by_type", h.Supplier.GetByType)
	//products
	mux.HandleFunc("/get_products", h.Product.GetAll)
	mux.HandleFunc("/get_prod_by_id", h.Product.GetById)
	mux.HandleFunc("/get_prod_by_params", h.Product.GetByParams)
	//auth
	mux.HandleFunc("/login", h.Auth.Login)
	mux.HandleFunc("/refresh", h.Auth.Refresh)
	//order
	mux.HandleFunc("/get_order", h.Order.GetById)
	mux.HandleFunc("/post_order", h.Order.Add)
	mux.Handle("/orders", mw.Validation(http.HandlerFunc(h.Order.GetByUserId)))

	log.Fatal(http.ListenAndServe(cfg.Port, mux))
}
