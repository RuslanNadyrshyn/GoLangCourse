package handlers

import (
	"Delivery/Delivery/cfg"
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/requests"
	"Delivery/Delivery/internal/repositories/responses"
	"Delivery/Delivery/internal/services"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type AuthHandler struct {
	cfg      *cfg.Config
	UserRepo repositories.IUserRepository
}

func NewAuthHandler(
	cfg *cfg.Config,
	UserRepo repositories.IUserRepository,
) *AuthHandler {
	return &AuthHandler{
		cfg:      cfg,
		UserRepo: UserRepo,
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
		user, err := h.UserRepo.GetByEmail(req.Email)
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
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Println(err)
			return
		}
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	case "POST":
		req := new(requests.RefreshRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		tokenService := services.NewTokenService(h.cfg)
		accessString, refreshString, err := tokenService.RefreshToken(req.AccessToken, req.RefreshToken, h.cfg.AccessSecret, h.cfg.RefreshSecret)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		resp := responses.LoginResponse{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Println(err)
			return
		}
	default:
		http.Error(w, "Only POST Method is Allowed", http.StatusMethodNotAllowed)
	}
}
