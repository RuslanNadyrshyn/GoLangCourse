package handlers

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/models"
	"Delivery/Delivery/internal/repositories/requests"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SupplierHandler struct {
	SupplierRepository          repositories.ISupplierRepository
	MenuRepository              repositories.IMenuRepository
	ProductRepository           repositories.IProductRepository
	ProductIngredientRepository repositories.IProductIngredientRepository
	IngredientRepository        repositories.IIngredientRepository
}

func NewSupplierHandler(
	SupplierRepository repositories.ISupplierRepository,
	MenuRepository repositories.IMenuRepository,
	ProductRepository repositories.IProductRepository,
	ProductIngredientRepository repositories.IProductIngredientRepository,
	IngredientRepository repositories.IIngredientRepository,
) *SupplierHandler {
	return &SupplierHandler{
		SupplierRepository:          SupplierRepository,
		MenuRepository:              MenuRepository,
		ProductRepository:           ProductRepository,
		ProductIngredientRepository: ProductIngredientRepository,
		IngredientRepository:        IngredientRepository,
	}
}

// AddSupplier ?
func (h *SupplierHandler) AddSupplier(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	case "GET":
		req := new(requests.SupplierRequest)
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		supplier := models.Supplier{
			Id:    req.Id,
			Name:  req.Name,
			Type:  req.Type,
			Image: req.Image,
			WorkingHours: struct {
				Opening string `json:"opening"`
				Closing string `json:"closing"`
			}{
				Opening: req.WorkingHours.Opening,
				Closing: req.WorkingHours.Closing,
			},
		}
		req.Id, err = h.SupplierRepository.Insert(&supplier)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		//Menu
		menuId, err := h.MenuRepository.Insert(req.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		//Products
		for i := range req.Menu {
			req.Menu[i].MenuId = menuId
			req.Menu[i].Id, err = h.ProductRepository.Insert(req.Menu[i])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Println(err)
				return
			}

			//Ingredients
			for j := range req.Menu[i].Ingredients {
				ingredientId, err := h.IngredientRepository.Insert(req.Menu[i].Ingredients[j])
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					fmt.Println(err)
					return
				}
				//ProductIngredients
				_, err = h.ProductIngredientRepository.Insert(req.Menu[i].Id, ingredientId)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					fmt.Println(err)
					return
				}
			}
		}
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(req.Id)
		if err != nil {
			log.Println(err)
			return
		}
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *SupplierHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var resp []requests.SupplierRequest
		suppliers, err := h.SupplierRepository.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
		for _, supplier := range suppliers {
			menu, err := h.ProductRepository.GetBySupplierId(supplier.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				log.Println(err)
				return
			}
			resp = append(resp, requests.SupplierRequest{
				Id:    supplier.Id,
				Name:  supplier.Name,
				Type:  supplier.Type,
				Image: supplier.Image,
				WorkingHours: struct {
					Opening string `json:"opening"`
					Closing string `json:"closing"`
				}{
					Opening: supplier.WorkingHours.Opening,
					Closing: supplier.WorkingHours.Closing,
				},
				Menu: menu,
			})
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Println(err)
			return
		}
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
