package main

import (
	"Delivery/Delivery/cfg"
	"Delivery/Delivery/internal/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := cfg.NewConfig(false)
	server.Start(config)
}
