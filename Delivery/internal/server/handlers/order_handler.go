package handlers

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/models"
	"Delivery/Delivery/internal/repositories/requests"
	"Delivery/Delivery/internal/repositories/responses"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	UserRepository         repositories.IUserRepository
	OrderRepository        repositories.IOrderRepository
	OrderProductRepository repositories.IOrderProductRepository
	ProductRepository      repositories.IProductRepository
}

func NewOrderHandler(
	UserRepository repositories.IUserRepository,
	OrderRepository repositories.IOrderRepository,
	OrderProductRepository repositories.IOrderProductRepository,
	ProductRepository repositories.IProductRepository,
) *OrderHandler {
	return &OrderHandler{
		UserRepository:         UserRepository,
		OrderRepository:        OrderRepository,
		OrderProductRepository: OrderProductRepository,
		ProductRepository:      ProductRepository,
	}
}

func (h *OrderHandler) Add(w http.ResponseWriter, r *http.Request) {
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

		var userId int
		if req.User.Id == 0 {
			userId, err = h.UserRepository.Insert(req.User)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Println(err)
				return
			}
		} else {
			userId = req.User.Id
		}

		orderId, err := h.OrderRepository.Insert(
			&models.Order{
				Price:   req.TotalPrice,
				UserId:  userId,
				Address: req.Address,
			})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		for _, product := range req.Products {
			_, err := h.OrderProductRepository.Insert(orderId, product.ProductId, product.Counter, product.ProductPrice)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Println(err)
				return
			}
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

		order, err := h.OrderRepository.GetById(id)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		orderProducts, err := h.OrderProductRepository.GetByOrderId(order.Id)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		var respProducts []responses.Product
		for _, orderProd := range orderProducts {
			product, err := h.ProductRepository.GetById(orderProd.ProductId)
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				fmt.Println(err)
				return
			}
			prod := responses.Product{
				Id:          product.Id,
				MenuId:      product.MenuId,
				Name:        product.Name,
				Price:       product.Price,
				Image:       product.Image,
				Type:        product.Type,
				Ingredients: product.Ingredients,
				Counter:     orderProd.Count,
				OldPrice:    orderProd.Price,
			}
			respProducts = append(respProducts, prod)
		}

		resp := &responses.OrderResponse{
			Id:        order.Id,
			UserId:    order.UserId,
			Address:   order.Address,
			Price:     order.Price,
			Products:  respProducts,
			CreatedAt: order.CreatedAt,
		}
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *OrderHandler) GetByUserId(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "GET":
		req := r.URL.Query().Get("userId")
		userId, err := strconv.Atoi(req)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		resp, err := h.OrderRepository.GetByUserId(userId)
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
