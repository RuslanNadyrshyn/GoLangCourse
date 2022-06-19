package responses

import "Delivery/backend/internal/repositories/models"

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Product struct {
	Id          int64    `json:"id"`
	MenuId      int64    `json:"menuId"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
	Counter     int64    `json:"counter"`
	OldPrice    float64  `json:"oldPrice"`
}

type OrderResponse struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"userId"`
	Address   string    `json:"address"`
	Price     float64   `json:"price"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Products  []Product `json:"products"`
}

type ProductResponse struct {
	Product  models.Product
	Supplier models.Supplier
}

type ProductsTypesResponse struct {
	Products []models.Product
	Types    []string
}

type SupplierResponse struct {
	Suppliers []models.Supplier
	Types     []string
}
