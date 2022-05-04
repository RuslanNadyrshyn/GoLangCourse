package handlers

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/requests"
	"Delivery/Delivery/internal/repositories/responses"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductRepository  repositories.IProductRepository
	SupplierRepository repositories.ISupplierRepository
}

func NewProductHandler(
	ProductRepository repositories.IProductRepository,
	SupplierRepository repositories.ISupplierRepository,
) *ProductHandler {
	return &ProductHandler{
		ProductRepository:  ProductRepository,
		SupplierRepository: SupplierRepository,
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	case "GET":
		products, err := h.ProductRepository.GetAll()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(products)
		if err != nil {
			fmt.Println(err)
			return
		}
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

		prod, err := h.ProductRepository.GetById(int64(id))
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		supplier, err := h.SupplierRepository.GetByProductId(prod.MenuId)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		resp := &responses.ProductResponse{
			Product:  prod,
			Supplier: supplier,
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
