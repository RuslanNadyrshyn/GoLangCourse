package Goroutines

import (
	"awesomeProject/internal/repositories/models"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type GoRoutinesRepository struct {
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

	PrintSupplier(supplier)
	return supplier
}

func PrintSupplier(supplier models.Supplier) {
	fmt.Println("id:", supplier.Id,
		"\nname:", supplier.Name,
		"\ntype:", supplier.Type,
		"\nimage:", supplier.Image,
		"\nopening:", supplier.WorkingHours.Opening,
		"\nclosing:", supplier.WorkingHours.Closing,
		"\nMenu:")
	for i := range supplier.Menu {
		fmt.Println("\tId:", supplier.Menu[i].Id,
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
