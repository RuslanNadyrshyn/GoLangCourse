package main

import (
	config2 "Delivery/Delivery/internal/auth/config"
	"Delivery/Delivery/internal/auth/server"
)

func main() {
	cfg := config2.NewConfig(false)

	server.Start(cfg)
}
