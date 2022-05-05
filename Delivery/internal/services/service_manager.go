package services

import (
	"Delivery/Delivery/cfg"
	"Delivery/Delivery/internal/repositories"
	"errors"
)

type ServiceManager struct {
	Token    ITokenService
	User     IUserService
	Supplier ISupplierService
	Product  IProductService
	Order    IOrderService
}

func NewServiceManager(store *repositories.Store, cfg *cfg.Config) (*ServiceManager, error) {
	if store == nil {
		return nil, errors.New("no store provided")
	}
	return &ServiceManager{
		User:     NewUserService(store),
		Token:    NewTokenService(cfg),
		Supplier: NewSupplierService(store),
		Product:  NewProductService(store),
		Order:    NewOrderService(store),
	}, nil
}
