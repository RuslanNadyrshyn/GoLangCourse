package requests

import "Delivery/Delivery/internal/repositories/models"

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
