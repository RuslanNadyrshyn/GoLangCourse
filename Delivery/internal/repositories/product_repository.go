package repositories

import (
	"Delivery/Delivery/internal/repositories/models"
	"errors"
)

type IProductRepository interface {
	GetAll() (products []*models.Product, err error)
	GetById(id int) (product *models.Product, err error)
}

type ProductRepository struct {
	products []*models.Product
}

func NewProductRepository() IProductRepository {
	products := []*models.Product{
		{
			Id:          1,
			Name:        "MyName",
			MenuId:      1,
			Price:       432.2,
			Image:       "fds",
			Type:        "pizza",
			Ingredients: []string{"pasta", "cheese"},
		},
		{
			Id:          2,
			Name:        "MyName2",
			MenuId:      2,
			Price:       4.2,
			Image:       "fdscas",
			Type:        "burger",
			Ingredients: []string{"beef", "souse", "onion"},
		},
	}

	return &ProductRepository{products: products}
}

func (p ProductRepository) GetAll() (products []*models.Product, err error) {
	return p.products, nil
}

func (p ProductRepository) GetById(id int) (product *models.Product, err error) {
	for _, product := range p.products {
		if product.Id == id {
			return product, nil
		}
	}
	return nil, errors.New("product not found")
}
