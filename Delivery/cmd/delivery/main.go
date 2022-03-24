package main

import (
	config2 "Delivery/Delivery/internal/auth/config"
	"Delivery/Delivery/internal/auth/server"
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//http://foodapi.true-tech.php.nixdev.co/swagger/doc.json
	cfg := config2.NewConfig(false)

	server.Start(cfg)
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
