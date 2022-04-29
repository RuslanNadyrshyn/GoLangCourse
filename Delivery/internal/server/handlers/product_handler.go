package handlers

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/requests"
	"Delivery/Delivery/internal/services"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	productRepository repositories.IProductRepository
	conn              *sql.DB
}

func NewProductHandler(
	productRepository repositories.IProductRepository,
	conn *sql.DB,
) *ProductHandler {
	return &ProductHandler{
		productRepository: productRepository,
		conn:              conn,
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	case "POST":
		if r.Method != http.MethodGet {
			http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		}

		dbService := services.NewDBService(h.conn)
		products, err := dbService.ProductRepo.GetAll()
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(products)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
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

		dbService := services.NewDBService(h.conn)
		resp, err := dbService.GetProductById(id)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
