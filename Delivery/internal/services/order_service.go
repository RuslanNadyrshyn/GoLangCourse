package services

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/models"
	"Delivery/Delivery/internal/repositories/requests"
	"Delivery/Delivery/internal/repositories/responses"
)

type OrderService struct {
	store *repositories.Store
}

func NewOrderService(store *repositories.Store) *OrderService {
	return &OrderService{
		store: store,
	}
}

func (r *OrderService) AddOrder(order *requests.OrderRequest) (userId int64, err error) {
	if order.User.Id == 0 {
		userId, err = r.store.UserRepo.InsertName(order.User)
		if err != nil {
			return userId, err
		}
	} else {
		userId = order.User.Id
	}
	var o = models.Order{
		Price:   order.TotalPrice,
		UserId:  userId,
		Address: order.Address,
	}
	orderId, err := r.store.OrderRepo.Insert(&o)
	if err != nil {
		return 0, err
	}

	for _, product := range order.Products {
		_, err := r.store.OrderProductRepo.Insert(orderId, product.ProductId, product.Counter, product.ProductPrice)
		if err != nil {
			return 0, err
		}
	}
	return orderId, nil
}

func (r *OrderService) GetById(id int64) (*responses.OrderResponse, error) {
	order, err := r.store.OrderRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	orderProducts, err := r.store.OrderProductRepo.GetByOrderId(order.Id)
	if err != nil {
		return nil, err
	}

	var respProducts []responses.Product
	for _, orderProd := range *orderProducts {
		product, err := r.store.ProductRepo.GetById(orderProd.ProductId)
		if err != nil {
			return nil, err
		}
		product.Ingredients, err = r.store.IngredientRepo.GetByProductId(product.Id)

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
	return resp, nil
}

func (r *OrderService) GetByUserId(id int64) (*[]models.Order, error) {
	orders, err := r.store.OrderRepo.GetByUserId(id)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
