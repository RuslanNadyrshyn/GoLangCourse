package handlers

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/services"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type SupplierHandler struct {
	supplierRepository repositories.ISupplierRepository
	conn               *sql.DB
}

func NewSupplierHandler(
	supplierRepository repositories.ISupplierRepository,
	conn *sql.DB,
) *SupplierHandler {
	return &SupplierHandler{
		supplierRepository: supplierRepository,
		conn:               conn,
	}
}

func (h *SupplierHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		dbService := services.NewDBService(h.conn)
		suppliers, err := dbService.GetAll()
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(suppliers)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
