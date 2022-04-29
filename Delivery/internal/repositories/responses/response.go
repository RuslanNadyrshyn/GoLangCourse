package responses

import "Delivery/Delivery/internal/repositories/models"

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Product struct {
	Id          int      `json:"id"`
	MenuId      int      `json:"menuId"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
	Counter     int      `json:"counter"`
	OldPrice    float64  `json:"oldPrice"`
}

type OrderResponse struct {
	Id       int       `json:"id"`
	UserId   int       `json:"userId"`
	Address  string    `json:"address"`
	Price    float64   `json:"price"`
	Products []Product `json:"products"`
}

type ProductResponse struct {
	Product  models.Product
	Supplier models.Supplier
}
