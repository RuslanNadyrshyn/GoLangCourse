package main

import (
	"awesomeProject/Goroutines"
	config2 "awesomeProject/internal/auth/config"
	"awesomeProject/internal/auth/server"
	"awesomeProject/internal/repositories"
	"awesomeProject/internal/repositories/database/Connection"
	"awesomeProject/internal/repositories/models"
	"awesomeProject/internal/worker_pool"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func main() {
	//http://foodapi.true-tech.php.nixdev.co/swagger/doc.json

	conn, err := Connection.Connect()
	if err != nil {
		log.Fatal(err)
	}
	cfg := config2.NewConfig(false)

	server.Start(cfg)

	///Goroutines
	poolConst := func() worker_pool.Worker {
		return worker{
			DB: conn,
		}
	}

	pool := worker_pool.NewPool(4, poolConst)
	go pool.Run()

	test := Goroutines.GoRoutinesRepository{}

	supplierService := repositories.NewSupplierService(conn)
	err = supplierService.SupplierRepo.DeleteAll()
	if err != nil {
		fmt.Println(err)
	}

	//var sup1 []models.Supplier
	//sup1 = test.ParseHTTP("http://foodapi.true-tech.php.nixdev.co/suppliers")
	//
	//for i := range sup1 {
	//	//test.PrintSupplier(sup1[i])
	//	pool.DataSource <- sup1[i]
	//}
	//
	//test.UpdateHTTP(sup1, conn)

	var sup []models.Supplier
	for i := 0; i < 7; i++ {
		sup = append(sup, test.ParseJson("./Goroutines/Data/"+strconv.Itoa(i+1)+".json"))
		pool.DataSource <- sup[i]
	}
	pool.Stop()

	// TESTS
	//supplierService := repositories.NewSupplierService(conn)
	//
	//result, err := supplierService.ProductRepo.GetByType("sushi")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for i := range result {
	//	fmt.Println(result[i])
	//}
	//
	//err = supplierService.SupplierRepo.DeleteByName("Burger Club")
	////if err != nil {
	////	fmt.Println(err)
	////}
	//err = supplierService.SupplierRepo.DeleteAll()
	//if err != nil {
	//	fmt.Println(err)
	//}
}

type worker struct {
	DB *sql.DB
}

func (c worker) Do(data interface{}, _ int) {
	sup, ok := data.(models.Supplier)
	if !ok {
		return
	}
	supplierService := repositories.NewSupplierService(c.DB)

	supplierService.CreateSupplierTX(sup)
	//supplierService.CreateSupplier(sup)
}

func (c worker) Stop() {

}
