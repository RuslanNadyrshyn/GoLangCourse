package server

import (
	"awesomeProject/internal/auth/config"
	"awesomeProject/internal/auth/repositories"
	"awesomeProject/internal/auth/server/handlers"
	"awesomeProject/internal/auth/services"
	"database/sql"
	"log"
	"net/http"
)

func Start(cfg *config.Config, conn *sql.DB) {
	userRepository := repositories.NewUserRepository()
	tokenService := services.NewTokenService(cfg)

	authHandler := handlers.NewAuthHandler(cfg, conn)
	userHandler := handlers.NewUserHandler(tokenService, userRepository, conn)

	mux := http.NewServeMux()
	mux.HandleFunc("/sign_in", userHandler.SignIn)
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/profile", userHandler.GetProfile)
	mux.HandleFunc("/refresh", authHandler.Refresh)

	log.Fatal(http.ListenAndServe(cfg.Port, mux))
}
