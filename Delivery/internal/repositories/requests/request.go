package requests

import (
	"Delivery/Delivery/internal/repositories/models"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	UserID       int    `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func SetupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

type OrderRequest struct {
	User       *models.User `json:"user"`
	TotalPrice float64      `json:"totalPrice"`
	Address    string       `json:"address"`
	Products   []struct {
		ProductId    int     `json:"id"`
		ProductPrice float64 `json:"price"`
		Counter      int     `json:"counter"`
	} `json:"products"`
}

type GetOrderRequest struct {
	Id int `json:"id"`
}
