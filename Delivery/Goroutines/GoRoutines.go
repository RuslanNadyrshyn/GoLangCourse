package Goroutines

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/models"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type GoRoutinesRepository struct {
}

func (grt GoRoutinesRepository) ParseHTTP(SuppliersLink string) (supplier []models.Supplier) {
	client := http.DefaultClient

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, SuppliersLink, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	suppliersMap := make(map[string][]models.Supplier, 0)

	err = json.NewDecoder(res.Body).Decode(&suppliersMap)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	suppliers := make([]models.Supplier, 0)

	suppliers = suppliersMap["suppliers"]

	for i := range suppliers {
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		req, err = http.NewRequestWithContext(ctx, http.MethodGet,
			"http://foodapi.true-tech.php.nixdev.co/suppliers/"+strconv.Itoa(suppliers[i].Id)+"/menu",
			nil)
		if err != nil {
			log.Fatal(err)
		}
		res, err = client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		menuMap := make(map[string][]models.Product, 0)
		json.NewDecoder(res.Body).Decode(&menuMap)
		suppliers[i].Menu = menuMap["menu"]
		//grt.PrintSupplier(suppliers[i])
		res.Body.Close()
		cancel()
	}
	return suppliers
}

func (grt GoRoutinesRepository) UpdateHTTP(suppliers []models.Supplier, conn *sql.DB) {
	supplierService := repositories.NewSupplierService(conn)
	client := http.DefaultClient

	for {
		time.Sleep(time.Minute)
		for i, sup := range suppliers {
			for j, prod := range sup.Menu {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				req, err := http.NewRequestWithContext(ctx, http.MethodGet,
					"http://foodapi.true-tech.php.nixdev.co/suppliers/"+
						strconv.Itoa(sup.Id)+"/menu/"+strconv.Itoa(prod.Id),
					nil)
				res, err := client.Do(req)
				if err != nil {
					log.Println(err)
				}
				var p models.Product
				err = json.NewDecoder(res.Body).Decode(&p)
				if err != nil {
					log.Println(err)
				}
				_, err = supplierService.ProductRepo.UpdatePrice(p)
				if err != nil {
					log.Println(err)
				}
				fmt.Println(p.Name, prod.Price, "-", p.Price)
				suppliers[i].Menu[j].Price = p.Price
				res.Body.Close()
				cancel()
			}
		}
	}
}

func (grt GoRoutinesRepository) ParseJson(jsonFile string) (supplier models.Supplier) {
	supplierData := []byte{}

	file, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for {
		chunk := make([]byte, 64)
		n, err := file.Read(chunk)
		if err == io.EOF { //
			break //
		}
		if err != nil {
			panic(err)
		}

		supplierData = append(supplierData, chunk[:n]...)
	}
	err = json.Unmarshal(supplierData, &supplier)
	if err != nil {
		panic(err)
	}

	//PrintSupplier(supplier)
	return supplier
}

func (grt GoRoutinesRepository) PrintSupplier(supplier models.Supplier) {
	fmt.Println("id:", supplier.Id,
		"\nname:", supplier.Name,
		"\ntype:", supplier.Type,
		"\nimage:", supplier.Image,
		"\nopening:", supplier.WorkingHours.Opening,
		"\nclosing:", supplier.WorkingHours.Closing,
		"\nMenu:")
	for i := range supplier.Menu {
		fmt.Println("\tId:", supplier.Menu[i].Id,
			"\n\tMenuId:", supplier.Menu[i].MenuId,
			"\n\tName:", supplier.Menu[i].Name,
			"\n\tPrice:", supplier.Menu[i].Price,
			"\n\tImage:", supplier.Menu[i].Image,
			"\n\tType:", supplier.Menu[i].Type,
			"\n\tIngredients:")
		for j := range supplier.Menu[i].Ingredients {
			fmt.Println("\t", supplier.Menu[i].Ingredients[j])
		}
	}
}
