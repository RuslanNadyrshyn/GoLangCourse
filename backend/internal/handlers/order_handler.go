package handlers

import (
	"Delivery/backend/internal/repositories/requests"
	"Delivery/backend/internal/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type OrderH struct {
	services *services.ServiceManager
}

func NewOrderH(services *services.ServiceManager) *OrderH {
	return &OrderH{
		services: services,
	}
}

func (h *OrderH) Add(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	case "POST":
		req := new(requests.OrderRequest)
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		resp, err := h.services.Order.AddOrder(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *OrderH) GetById(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "GET":
		req := r.URL.Query().Get("id")
		id, err := strconv.ParseInt(req, 0, 64)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		resp, err := h.services.Order.GetById(id)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
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

func (h *OrderH) GetByUserId(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	case "GET":
		requestToken := h.services.Token.GetTokenFromBearerString(r.Header.Get("Authorization"))
		claims, err := h.services.Token.ValidateAccessToken(requestToken)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		resp, err := h.services.Order.GetByUserId(claims.ID)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
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
