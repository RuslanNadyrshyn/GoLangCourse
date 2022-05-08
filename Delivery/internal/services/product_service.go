package services

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/models"
	"Delivery/Delivery/internal/repositories/responses"
)

type ProductService struct {
	store *repositories.Store
}

func NewProductService(store *repositories.Store) *ProductService {
	return &ProductService{
		store: store,
	}
}

func (r *ProductService) GetAll() (*[]models.Product, error) {

	products, err := r.store.ProductRepo.GetAll()
	if err != nil {
		return nil, err
	}
	for _, p := range *products {
		p.Ingredients, err = r.store.IngredientRepo.GetByProductId(p.Id)
		if err != nil {
			return nil, err
		}
	}

	return products, nil
}

func (r *ProductService) GetById(id int64) (*responses.ProductResponse, error) {
	product, err := r.store.ProductRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	product.Ingredients, err = r.store.IngredientRepo.GetByProductId(product.Id)
	if err != nil {
		return nil, err
	}

	supplier, err := r.store.SupplierRepo.GetByProductId(product.MenuId)
	if err != nil {
		return nil, err
	}

	resp := responses.ProductResponse{
		Product:  *product,
		Supplier: *supplier,
	}
	return &resp, nil
}

func (r *ProductService) UpdatePrice(p *models.Product) (productId int64, err error) {
	productId, err = r.store.ProductRepo.UpdatePrice(p)
	if err != nil {
		return productId, err
	}

	return productId, nil
}
