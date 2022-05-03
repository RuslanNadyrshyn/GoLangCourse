package handlers

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/requests"
	"Delivery/Delivery/internal/repositories/responses"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductRepository  repositories.IProductRepository
	MenuRepository     repositories.IMenuRepository
	SupplierRepository repositories.ISupplierRepository
}

func NewProductHandler(
	ProductRepository repositories.IProductRepository,
	MenuRepository repositories.IMenuRepository,
	SupplierRepository repositories.ISupplierRepository,
) *ProductHandler {
	return &ProductHandler{
		ProductRepository:  ProductRepository,
		MenuRepository:     MenuRepository,
		SupplierRepository: SupplierRepository,
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

		products, err := h.ProductRepository.GetAll()
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(products)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

// GetById ?
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

		prod, err := h.ProductRepository.GetById(id)
		if err != nil {
			fmt.Println(err)
		}

		supplierId, err := h.MenuRepository.GetById(prod.MenuId)
		if err != nil {
			fmt.Println(err)
		}

		supplier, err := h.SupplierRepository.GetById(supplierId)
		resp := &responses.ProductResponse{
			Product:  prod,
			Supplier: supplier,
		}
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
