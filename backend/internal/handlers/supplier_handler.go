package handlers

import (
	"Delivery/backend/internal/repositories/requests"
	"Delivery/backend/internal/repositories/responses"
	"Delivery/backend/internal/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SupplierHandler struct {
	services *services.ServiceManager
}

func NewSupplierHandler(services *services.ServiceManager) *SupplierHandler {
	return &SupplierHandler{
		services: services,
	}
}

func (h *SupplierHandler) Insert(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	case "GET":
		req := new(requests.SupplierRequest)
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}

		id, err := h.services.Supplier.Insert(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(id)
		if err != nil {
			log.Println(err)
			return
		}
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *SupplierHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "GET":
		suppliers, err := h.services.Supplier.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}

		types, err := h.services.Supplier.GetTypes()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
		resp := responses.SupplierResponse{
			Suppliers: *suppliers,
			Types:     *types,
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Println(err)
			return
		}
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *SupplierHandler) GetByType(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "GET":
		req := r.URL.Query().Get("type")
		resp, err := h.services.Supplier.GetByType(req)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
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
