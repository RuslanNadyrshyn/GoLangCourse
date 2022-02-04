package handlers

import (
	"awesomeProject/internal/auth/config"
	"awesomeProject/internal/auth/repositories"
	"awesomeProject/internal/auth/services"
	"net/http"
	"net/http/httptest"
)

func Start() *httptest.Server {
	cfg := config.NewConfig(false)

	userRepository := repositories.NewUserRepositoryMock()
	tokenService := services.NewTokenService(cfg)

	authHandler := NewAuthHandler(cfg)
	userHandler := NewUserHandler(tokenService, userRepository)

	mux := http.NewServeMux()
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/profile", userHandler.GetProfile)

	return httptest.NewServer(mux)
}
