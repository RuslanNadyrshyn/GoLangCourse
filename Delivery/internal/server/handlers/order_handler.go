package handlers

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/requests"
	"Delivery/Delivery/internal/services"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	orderRepository repositories.IOrderRepository
	conn            *sql.DB
}

func NewOrderHandler(
	orderRepository repositories.IOrderRepository,
	conn *sql.DB,
) *OrderHandler {
	return &OrderHandler{
		orderRepository: orderRepository,
		conn:            conn,
	}
}

func (h *OrderHandler) Add(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode("OKAY")
		if err != nil {
			fmt.Println(err)
			return
		}
	case "POST":
		req := new(requests.OrderRequest)
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		dbService := services.NewDBService(h.conn)
		orderId, err := dbService.AddOrder(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(orderId)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *OrderHandler) GetById(w http.ResponseWriter, r *http.Request) {
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
		resp, err := dbService.GetOrderById(id)
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
