package repositories

import (
	"Delivery/Delivery/internal/repositories/models"
	"errors"
)

type ISupplierRepository interface {
	GetAll() (suppliers []*models.Supplier, err error)
	Add(supplier *models.Supplier) (supplierId int, err error)
	Delete(supplierId int) error
}

type SupplierRepository struct {
	suppliers []*models.Supplier
}

func NewSupplierRepository() ISupplierRepository {
	suppliers := []*models.Supplier{
		{
			Id:    1,
			Name:  "pizza hub",
			Type:  "cafe",
			Image: "fsd",
			WorkingHours: struct {
				Opening string `json:"opening"`
				Closing string `json:"closing"`
			}{
				Opening: "7:00",
				Closing: "21:00",
			},
		},
		{
			Id:    2,
			Name:  "burger hub",
			Type:  "fastfood",
			Image: "fdfs",
			WorkingHours: struct {
				Opening string `json:"opening"`
				Closing string `json:"closing"`
			}{
				Opening: "7:00",
				Closing: "21:00",
			},
		},
	}

	return &SupplierRepository{suppliers: suppliers}
}

func (s *SupplierRepository) GetAll() (suppliers []*models.Supplier, err error) {
	return s.suppliers, nil
}

func (s *SupplierRepository) Add(supplier *models.Supplier) (supplierId int, err error) {
	s.suppliers = append(s.suppliers, supplier)

	return supplier.Id, nil
}

func (s *SupplierRepository) Delete(supplierId int) error {
	for i, sup := range s.suppliers {
		if sup.Id == supplierId {
			s.suppliers = append(s.suppliers[:i], s.suppliers[i+1:]...)
			return nil
		}
	}
	return errors.New("supplier not found")
}
