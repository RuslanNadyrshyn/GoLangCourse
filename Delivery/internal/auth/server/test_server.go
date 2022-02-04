package server

import (
	"awesomeProject/internal/auth/config"
	"awesomeProject/internal/auth/repositories"
	"awesomeProject/internal/auth/server/handlers"
	"awesomeProject/internal/auth/services"
	"log"
	"net/http"
)

func Start(cfg *config.Config) {
	userRepository := repositories.NewUserRepository()
	tokenService := services.NewTokenService(cfg)

	authHandler := handlers.NewAuthHandler(cfg)
	userHandler := handlers.NewUserHandler(tokenService, userRepository)

	mux := http.NewServeMux()
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/profile", userHandler.GetProfile)

	log.Fatal(http.ListenAndServe(cfg.Port, mux))
}
