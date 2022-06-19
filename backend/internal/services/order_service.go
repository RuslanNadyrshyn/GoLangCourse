package services

import (
	"Delivery/backend/internal/repositories"
	"Delivery/backend/internal/repositories/models"
	"Delivery/backend/internal/repositories/requests"
	"Delivery/backend/internal/repositories/responses"
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
	err = r.store.OrderRepo.BeginTx()
	if err != nil {
		return 0, err
	}

	if order.User.Id == 0 {
		r.store.UserRepo.SetTx(r.store.OrderRepo.GetTx())
		userId, err = r.store.UserRepo.InsertName(order.User)
		if err != nil {
			_ = r.store.UserRepo.RollbackTx()
			return 0, err
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
		_ = r.store.OrderRepo.RollbackTx()
		return 0, err
	}
	r.store.OrderProductRepo.SetTx(r.store.OrderRepo.GetTx())
	for _, product := range order.Products {
		_, err := r.store.OrderProductRepo.Insert(orderId, product.ProductId, product.Counter, product.ProductPrice)
		if err != nil {
			_ = r.store.OrderProductRepo.RollbackTx()
			return 0, err
		}
	}
	err = r.store.OrderRepo.CommitTx()
	if err != nil {
		_ = r.store.OrderRepo.RollbackTx()
		return 0, err
	}

	r.store.OrderRepo.SetTx(nil)
	r.store.UserRepo.SetTx(nil)
	r.store.OrderProductRepo.SetTx(nil)

	return orderId, nil
}

func (r *OrderService) GetById(id int64) (*responses.OrderResponse, error) {
	err := r.store.OrderRepo.BeginTx()
	if err != nil {
		return nil, err
	}
	order, err := r.store.OrderRepo.GetById(id)
	if err != nil {
		_ = r.store.OrderRepo.RollbackTx()
		return nil, err
	}

	r.store.OrderProductRepo.SetTx(r.store.OrderRepo.GetTx())
	orderProducts, err := r.store.OrderProductRepo.GetByOrderId(order.Id)
	if err != nil {
		_ = r.store.OrderProductRepo.RollbackTx()
		return nil, err
	}
	r.store.ProductRepo.SetTx(r.store.OrderRepo.GetTx())
	r.store.IngredientRepo.SetTx(r.store.OrderRepo.GetTx())

	var respProducts []responses.Product
	for _, orderProd := range *orderProducts {
		product, err := r.store.ProductRepo.GetById(orderProd.ProductId)
		if err != nil {
			_ = r.store.ProductRepo.RollbackTx()
			return nil, err
		}
		product.Ingredients, err = r.store.IngredientRepo.GetByProductId(product.Id)
		if err != nil {
			_ = r.store.IngredientRepo.RollbackTx()
			return nil, err
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
	err = r.store.OrderRepo.CommitTx()
	if err != nil {
		_ = r.store.OrderRepo.RollbackTx()
		return nil, err
	}

	r.store.OrderRepo.SetTx(nil)
	r.store.OrderProductRepo.SetTx(nil)
	r.store.ProductRepo.SetTx(nil)
	r.store.IngredientRepo.SetTx(nil)

	return resp, nil
}

func (r *OrderService) GetByUserId(id int64) (*[]models.Order, error) {
	orders, err := r.store.OrderRepo.GetByUserId(id)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
