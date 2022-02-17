package handlers

import (
	"awesomeProject/internal/auth/config"
	"awesomeProject/internal/auth/repositories"
	"awesomeProject/internal/auth/services"
	"database/sql"
	"net/http"
	"net/http/httptest"
)

func Start(cfg *config.Config, conn *sql.DB) *httptest.Server {
	userRepository := repositories.NewUserRepositoryMock()
	tokenService := services.NewTokenService(cfg)

	authHandler := NewAuthHandler(cfg, conn)
	userHandler := NewUserHandler(tokenService, userRepository, conn)

	mux := http.NewServeMux()
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/profile", userHandler.GetProfile)
	mux.HandleFunc("/refresh", authHandler.Refresh)
	mux.HandleFunc("/sign_in", userHandler.SignIn)

	return httptest.NewServer(mux)
}
