package main

import (
	"awesomeProject/internal/repositories"
	"awesomeProject/internal/repositories/database/Connection"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var (
	conn, _         = Connection.Connect()
	supplierService = repositories.NewSupplierService(conn)
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/get_products", GetProducts)
	server.HandleFunc("/get_suppliers", GetSuppliers)

	http.ListenAndServe(":8080", server)
}
func GetSuppliers(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resp, "not allowed", http.StatusMethodNotAllowed)
	}

	suppliers, err := supplierService.SupplierRepo.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	data, _ := json.Marshal(suppliers)
	resp.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(resp, string(data))
}

func GetProducts(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resp, "not allowed", http.StatusMethodNotAllowed)
	}

	products, err := supplierService.ProductRepo.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	data, _ := json.Marshal(products)
	resp.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(resp, string(data))
}
