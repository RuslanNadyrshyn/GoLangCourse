package handlers

import (
	"Delivery/backend/internal/repositories/requests"
	"Delivery/backend/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ProductH struct {
	services *services.ServiceManager
}

func NewProductH(services *services.ServiceManager) *ProductH {
	return &ProductH{
		services: services,
	}
}

func (h *ProductH) GetAll(w http.ResponseWriter, r *http.Request) {
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

func (h *ProductH) GetById(w http.ResponseWriter, r *http.Request) {
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

		prod, err := h.services.Product.GetById(int64(id))
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(prod)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
