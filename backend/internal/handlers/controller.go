package handlers

import (
	"Delivery/backend/cfg"
	"Delivery/backend/internal/services"
)

type Controller struct {
	Auth     *AuthHandler
	User     *UserHandler
	Supplier *SupplierHandler
	Product  *ProductHandler
	Order    *OrderHandler
}

func NewController(
	services *services.ServiceManager, cfg *cfg.Config,
) *Controller {
	return &Controller{
		Auth:     NewAuthHandler(services, cfg),
		User:     NewUserHandler(services),
		Supplier: NewSupplierHandler(services),
		Product:  NewProductHandler(services),
		Order:    NewOrderHandler(services),
	}
}
