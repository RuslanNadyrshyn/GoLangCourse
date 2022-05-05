package repositories

import "database/sql"

type Store struct {
	UserRepo              *UserRepo
	SupplierRepo          *SupplierRepo
	MenuRepo              *MenuRepo
	ProductRepo           *ProductRepo
	ProductIngredientRepo *ProductIngredientRepo
	IngredientRepo        *IngredientRepo
	OrderRepo             *OrderRepo
	OrderProductRepo      *OrderProductRepo
}

func NewStore(conn *sql.DB) *Store {
	UserRepo := NewUserRepo(conn)
	SupplierRepo := NewSupplierRepo(conn)
	MenuRepo := NewMenuRepo(conn)
	ProductRepo := NewProductRepo(conn)
	ProductIngredientRepo := NewProductIngredientRepo(conn)
	IngredientRepo := NewIngredientRepo(conn)
	OrderRepo := NewOrderRepo(conn)
	OrderProductRepo := NewOrderProductRepo(conn)

	return &Store{
		UserRepo:              UserRepo,
		SupplierRepo:          SupplierRepo,
		MenuRepo:              MenuRepo,
		ProductRepo:           ProductRepo,
		ProductIngredientRepo: ProductIngredientRepo,
		IngredientRepo:        IngredientRepo,
		OrderRepo:             OrderRepo,
		OrderProductRepo:      OrderProductRepo,
	}
}
