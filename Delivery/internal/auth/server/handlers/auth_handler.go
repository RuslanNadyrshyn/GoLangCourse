package handlers

import (
	"Delivery/Delivery/internal/auth/config"
	"Delivery/Delivery/internal/auth/requests"
	"Delivery/Delivery/internal/auth/responses"
	"Delivery/Delivery/internal/auth/services"
	"Delivery/Delivery/internal/repositories/database"
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type AuthHandler struct {
	cfg  *config.Config
	conn *sql.DB
}

func NewAuthHandler(cfg *config.Config, conn *sql.DB) *AuthHandler {
	return &AuthHandler{
		cfg:  cfg,
		conn: conn,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)

	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	case "POST":
		req := new(requests.LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		userRepo := database.NewUserRepository(h.conn)

		user, err := userRepo.GetByEmail(req.Email)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		tokenService := services.NewTokenService(h.cfg)
		accessString, err := tokenService.GenerateAccessToken(user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		refreshString, err := tokenService.GenerateRefreshToken(user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := responses.LoginResponse{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := new(requests.RefreshRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tokenService := services.NewTokenService(h.cfg)

		accessString, refreshString, err := tokenService.RefreshToken(req.AccessToken, req.RefreshToken, h.cfg.AccessSecret, h.cfg.RefreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := responses.LoginResponse{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {

}
