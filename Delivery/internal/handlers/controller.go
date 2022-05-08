package handlers

import (
	"Delivery/Delivery/cfg"
	"Delivery/Delivery/internal/services"
)

type Controller struct {
	Auth     *AuthH
	User     *UserH
	Supplier *SupplierH
	Product  *ProductH
	Order    *OrderH
}

func NewController(
	services *services.ServiceManager, cfg *cfg.Config,
) *Controller {
	return &Controller{
		Auth:     NewAuthH(services, cfg),
		User:     NewUserH(services),
		Supplier: NewSupplierH(services),
		Product:  NewProductH(services),
		Order:    NewOrderH(services),
	}
}
