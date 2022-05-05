package tests

import (
	"Delivery/Delivery/cfg"
	"net/http/httptest"
)

func Start(cfg *cfg.Config) *httptest.Server {
	//userRepository := NewUserRepositoryMock()
	//tokenService := services.NewTokenService(cfg)
	//
	//authHandler := handlers.NewAuthHandler(cfg, userRepository)
	//userHandler := NewUserHandler(tokenService, userRepository)
	//
	//mux := http.NewServeMux()
	//mux.HandleFunc("/login", authHandler.Login)
	//mux.HandleFunc("/profile", userHandler.GetProfile)
	//mux.HandleFunc("/refresh", authHandler.Refresh)
	//mux.HandleFunc("/sign_in", userHandler.SignIn)
	//
	//return httptest.NewServer(mux)
	return nil
}
