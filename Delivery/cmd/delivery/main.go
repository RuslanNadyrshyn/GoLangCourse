package main

import (
	"awesomeProject/Goroutines"
	"awesomeProject/internal/repositories/database"
	"awesomeProject/internal/repositories/database/Connection"
	"awesomeProject/internal/repositories/models"
	"awesomeProject/internal/worker_pool"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func main() {
	conn, err := Connection.Connect()
	if err != nil {
		log.Fatal(err)
	}

	//Supplier
	poolConst := func() worker_pool.Worker {
		return w{
			DB: conn,
		}
	}

	pool := worker_pool.NewPool(4, poolConst)
	///////Goroutines
	go pool.Run()

	test := Goroutines.GoRoutinesRepository{}

	var sup []models.Supplier
	for i := 0; i < 7; i++ {
		sup = append(sup, test.Insert("./Goroutines/Data/"+strconv.Itoa(i+1)+".json"))
		//fmt.Println("sup =", sup[i])
		pool.DataSource <- sup[i]
	}
	pool.Stop()
}

type w struct {
	DB *sql.DB
}

func (c w) Do(data interface{}, i int) {
	d, ok := data.(models.Supplier)
	if !ok {
		return
	}
	supplierRepo := database.NewSupplierRepository(c.DB)

	tx, _ := c.DB.Begin()
	supplierRepo.TX = tx

	_, err := supplierRepo.Insert(d)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
}

func (c w) Stop() {

}
