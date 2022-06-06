package main

import (
	"Delivery/backend/cfg"
	"Delivery/backend/internal/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := cfg.NewConfig(false)
	server.Start(config)
}
