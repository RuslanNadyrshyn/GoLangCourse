package handlers

import (
	"Delivery/backend/internal/repositories/requests"
	"Delivery/backend/internal/services"
	"encoding/json"
	"log"
	"net/http"
)

type SupplierH struct {
	services *services.ServiceManager
}

func NewSupplierH(services *services.ServiceManager) *SupplierH {
	return &SupplierH{
		services: services,
	}
}

func (h *SupplierH) AddSupplier(w http.ResponseWriter, r *http.Request) {
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

func (h *SupplierH) GetAll(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		resp, err := h.services.Supplier.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
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
