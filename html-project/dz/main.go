package main

import (
	"awesomeProject/internal/repositories"
	"awesomeProject/internal/repositories/database"
	"awesomeProject/internal/repositories/database/Connection"
	"awesomeProject/internal/repositories/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var (
	conn, err       = Connection.Connect()
	supplierService = repositories.NewSupplierService(conn)
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/get_products", GetProducts)
	//server.HandleFunc("/add_product", AddProduct)

	http.ListenAndServe(":8080", server)
}

func AddProduct(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var id int
		product := new(models.Product)
		if err := json.NewDecoder(req.Body).Decode(&product); err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
		}
		prodRepo := database.NewProductRepository(conn)

		id, err = prodRepo.Insert(*product)
		if err != nil {
			http.Error(resp, "Insert error!", http.StatusBadRequest)
			log.Println(err)
			return
		}

		response := models.Product{
			Id:          id,
			Name:        product.Name,
			MenuId:      product.MenuId,
			Price:       product.Price,
			Image:       product.Image,
			Type:        product.Type,
			Ingredients: product.Ingredients,
		}

		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(response)
	default:
		http.Error(resp, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}

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
	//resp.Header().Add("Test", "Hello from keys man")

}
