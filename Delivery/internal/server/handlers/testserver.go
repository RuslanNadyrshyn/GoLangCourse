package handlers

//
//import (
//	"Delivery/Delivery/cfg"
//	"Delivery/Delivery/internal/repositories"
//	"Delivery/Delivery/internal/services"
//	"database/sql"
//	"net/http"
//	"net/http/httptest"
//)
//
//func Start(cfg *cfg.Config, conn *sql.DB) *httptest.Server {
//	//userRepository := repositories.NewUserRepositoryMock()
//	tokenService := services.NewTokenService(cfg)
//
//	authHandler := NewAuthHandler(cfg, conn)
//	//userHandler := NewUserHandler(tokenService, userRepository)
//
//	mux := http.NewServeMux()
//	mux.HandleFunc("/login", authHandler.Login)
//	mux.HandleFunc("/profile", userHandler.GetProfile)
//	mux.HandleFunc("/refresh", authHandler.Refresh)
//	mux.HandleFunc("/sign_in", userHandler.SignIn)
//
//	return httptest.NewServer(mux)
//}
