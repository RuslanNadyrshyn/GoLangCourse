package services

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/models"
	"Delivery/Delivery/internal/repositories/requests"
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

func (r *SupplierService) GetAll() (*[]requests.SupplierRequest, error) {
	suppliers, err := r.store.SupplierRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var resp []requests.SupplierRequest
	for _, s := range *suppliers {
		menuId, err := r.store.MenuRepo.GetIdBySupId(s.Id)
		menu, err := r.store.ProductRepo.GetByMenuId(menuId)
		if err != nil {
			return nil, err
		}
		for i, p := range menu {
			menu[i].Ingredients, err = r.store.IngredientRepo.GetByProductId(p.Id)
			if err != nil {
				return nil, err
			}
		}

		resp = append(resp, requests.SupplierRequest{
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
			Menu: menu,
		})
	}
	return &resp, nil
}
