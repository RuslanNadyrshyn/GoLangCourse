package handlers

import (
	"Delivery/backend/internal/repositories/requests"
	"Delivery/backend/internal/repositories/responses"
	"Delivery/backend/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	services *services.ServiceManager
}

func NewProductHandler(services *services.ServiceManager) *ProductHandler {
	return &ProductHandler{
		services: services,
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	case "GET":
		products, err := h.services.Product.GetAll()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		types, err := h.services.Product.GetTypes(requests.SortRequest{})
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		resp := responses.ProductsTypesResponse{
			Products: *products,
			Types:    *types,
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			fmt.Println(err)
			return
		}
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

		product, err := h.services.Product.GetById(int64(id))
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		supplier, err := h.services.Supplier.GetByProductId(int64(id))
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		resp := responses.ProductResponse{
			Product:  *product,
			Supplier: *supplier,
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

func (h *ProductHandler) GetByParams(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "GET":
		var supId int
		var err error

		supIdReq := r.URL.Query().Get("sup_id")
		supType := r.URL.Query().Get("sup_type")
		prodType := r.URL.Query().Get("prod_type")

		if supIdReq != "" {
			supId, err = strconv.Atoi(supIdReq)
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				fmt.Println(err)
				return
			}
		}

		params := requests.SortRequest{
			SupplierId:   int64(supId),
			SupplierType: supType,
			ProductType:  prodType,
		}

		prod, err := h.services.Product.GetByParams(params)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		types, err := h.services.Product.GetTypes(params)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		resp := responses.ProductsTypesResponse{
			Products: *prod,
			Types:    *types,
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
