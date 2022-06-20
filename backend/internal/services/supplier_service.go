package services

import (
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories/models"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories/requests"
)

type SupplierService struct {
	store *repositories.Store
}

func NewSupplierService(store *repositories.Store) *SupplierService {
	return &SupplierService{
		store: store,
	}
}

func (r *SupplierService) Insert(s *requests.SupplierRequest) (id int64, err error) {
	err = r.store.SupplierRepo.BeginTx()
	if err != nil {
		return 0, err
	}
	//Supplier
	supplier := models.Supplier{
		Id:    s.Id,
		Name:  s.Name,
		Type:  s.Type,
		Image: s.Image,
		WorkingHours: struct {
			Opening string `json:"opening"`
			Closing string `json:"closing"`
		}{
			Opening: s.WorkingHours.Opening,
			Closing: s.WorkingHours.Closing,
		},
	}
	s.Id, err = r.store.SupplierRepo.Insert(&supplier)
	if err != nil {
		_ = r.store.SupplierRepo.RollbackTx()
		return 0, err
	}

	//Menu
	r.store.MenuRepo.SetTx(r.store.SupplierRepo.GetTx())
	menuId, err := r.store.MenuRepo.Insert(s.Id)
	if err != nil {
		_ = r.store.MenuRepo.RollbackTx()
		return 0, err
	}

	//Products
	r.store.ProductRepo.SetTx(r.store.SupplierRepo.GetTx())

	for _, p := range s.Menu {
		p.MenuId = menuId
		p.Id, err = r.store.ProductRepo.Insert(&p)
		if err != nil {
			_ = r.store.ProductRepo.RollbackTx()
			return 0, err
		}

		//Ingredients
		r.store.IngredientRepo.SetTx(r.store.SupplierRepo.GetTx())
		r.store.ProductIngredientRepo.SetTx(r.store.SupplierRepo.GetTx())
		for _, ingredient := range p.Ingredients {
			ingredientId, err := r.store.IngredientRepo.Insert(ingredient)
			if err != nil {
				_ = r.store.IngredientRepo.RollbackTx()
				return 0, err
			}

			//ProductIngredients
			_, err = r.store.ProductIngredientRepo.Insert(p.Id, ingredientId)
			if err != nil {
				_ = r.store.ProductIngredientRepo.RollbackTx()
				return 0, err
			}
		}
	}

	err = r.store.SupplierRepo.CommitTx()
	if err != nil {
		_ = r.store.SupplierRepo.RollbackTx()
		return 0, err
	}

	r.store.SupplierRepo.SetTx(nil)
	r.store.MenuRepo.SetTx(nil)
	r.store.ProductRepo.SetTx(nil)
	r.store.IngredientRepo.SetTx(nil)
	r.store.ProductIngredientRepo.SetTx(nil)

	return supplier.Id, nil
}

func (r *SupplierService) GetAll() (*[]models.Supplier, error) {
	suppliers, err := r.store.SupplierRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (r *SupplierService) GetTypes() (*[]string, error) {
	types, err := r.store.SupplierRepo.GetTypes()
	if err != nil {
		return nil, err
	}
	return &types, nil
}

func (r *SupplierService) GetByType(supType string) (suppliers *[]models.Supplier, err error) {
	switch {
	case supType == "":
		suppliers, err = r.store.SupplierRepo.GetAll()
	case supType == "workingHours":
		suppliers, err = r.store.SupplierRepo.GetByWorkingHours()
	case supType != "":
		suppliers, err = r.store.SupplierRepo.GetByType(supType)
	}

	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (r *SupplierService) GetByProductId(prodId int64) (*models.Supplier, error) {
	supplier, err := r.store.SupplierRepo.GetByProductId(prodId)
	if err != nil {
		return nil, err
	}
	return supplier, nil
}
