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

type OrderResponse struct {
	Id       int              `json:"id"`
	User     models.User      `json:"user"`
	Address  string           `json:"address"`
	Products []models.Product `json:"products"`
	Price    float32          `json:"price"`
}
