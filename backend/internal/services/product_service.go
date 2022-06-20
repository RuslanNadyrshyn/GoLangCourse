package services

import (
	"fmt"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories/models"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories/requests"
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
	for i, p := range products {
		products[i].Ingredients, err = r.store.IngredientRepo.GetByProductId(p.Id)
		if err != nil {
			return nil, err
		}
	}

	return &products, nil
}

func (r *ProductService) GetById(id int64) (*models.Product, error) {
	product, err := r.store.ProductRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	product.Ingredients, err = r.store.IngredientRepo.GetByProductId(product.Id)
	if err != nil {
		return nil, err
	}

	//supplierId, err := r.store.MenuRepo.GetSupIdById(product.MenuId)
	//if err != nil {
	//	return nil, err
	//}
	//
	//supplier, err := r.store.SupplierRepo.GetById(supplierId)
	//if err != nil {
	//	return nil, err
	//}
	//
	//resp := responses.ProductResponse{
	//	Product:  *product,
	//	Supplier: *supplier,
	//}
	return product, nil
}

func (r *ProductService) GetByParams(params requests.SortRequest) (*[]models.Product, error) {
	var prod []models.Product
	var products []models.Product
	var err error

	if params.ProductType != "" && params.SupplierType != "" {
		params.ProductType = fmt.Sprintf(" AND type = '%s'", params.ProductType)
	}

	switch {
	case params.SupplierId != 0: // SupplierId != 0
		products, err = r.store.ProductRepo.GetBySupId(params.SupplierId, params.ProductType)
		if err != nil {
			return nil, err
		}
	case params.SupplierType == "workingHours": // SupplierId == 0 && SupplierType == "WorkingHours"
		products, err = r.store.ProductRepo.GetByWorkingHours(params.ProductType)
		if err != nil {
			return nil, err
		}
	case params.SupplierType != "": // SupplierId == 0 && SupplierType != ""
		products, err = r.store.ProductRepo.GetBySupType(params.SupplierType, params.ProductType)
		if err != nil {
			return nil, err
		}
	case params.ProductType != "": // SupplierId == 0 && SupplierType == "" && ProductType != 0
		products, err = r.store.ProductRepo.GetByType(params.ProductType)
		if err != nil {
			return nil, err
		}
	case params.ProductType == "": // SupplierId == 0 && SupplierType == "" && ProductType == 0
		products, err = r.store.ProductRepo.GetAll()
		if err != nil {
			return nil, err
		}
	}

	for i, p := range products {
		products[i].Ingredients, err = r.store.IngredientRepo.GetByProductId(p.Id)
		if err != nil {
			return nil, err
		}
		prod = append(prod, p)
	}

	return &products, nil
}

func (r *ProductService) GetTypes(params requests.SortRequest) (t *[]string, err error) {
	var types []string

	switch {
	case params.SupplierId != 0: // SupplierId == 0
		types, err = r.store.ProductRepo.GetTypesBySupId(params.SupplierId)
		if err != nil {
			return nil, err
		}
	case params.SupplierType != "": // SupplierId == 0 && SupplierType != 0
		types, err = r.store.ProductRepo.GetTypesBySupType(params.SupplierType)
		if err != nil {
			return nil, err
		}
	case params.SupplierType == "": // SupplierId == 0 && SupplierType == 0
		types, err = r.store.ProductRepo.GetAllTypes()
		if err != nil {
			return nil, err
		}
	}
	return &types, nil
}

func (r *ProductService) UpdatePrice(p *models.Product) (productId int64, err error) {
	productId, err = r.store.ProductRepo.UpdatePrice(p)
	if err != nil {
		return productId, err
	}

	return productId, nil
}
