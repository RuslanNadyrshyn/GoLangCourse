package main

import (
	"Delivery/Delivery/cfg"
	"Delivery/Delivery/internal/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//http://foodapi.true-tech.php.nixdev.co/swagger/doc.json
	config := cfg.NewConfig(false)
	server.Start(config)
}
