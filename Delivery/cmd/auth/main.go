package main

import (
	config2 "awesomeProject/internal/auth/config"
	"awesomeProject/internal/auth/server"
)

func main() {
	cfg := config2.NewConfig(false)

	server.Start(cfg)
}
