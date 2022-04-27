package main

import (
	config2 "Delivery/Delivery/internal/auth/config"
	"Delivery/Delivery/internal/auth/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//http://foodapi.true-tech.php.nixdev.co/swagger/doc.json
	cfg := config2.NewConfig(false)

	server.Start(cfg)
}
