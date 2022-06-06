package handlers

import (
	"Delivery/backend/cfg"
	"Delivery/backend/internal/repositories/requests"
	"Delivery/backend/internal/repositories/responses"
	"Delivery/backend/internal/services"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type AuthH struct {
	services *services.ServiceManager
	cfg      *cfg.Config
}

func NewAuthH(
	services *services.ServiceManager,
	cfg *cfg.Config,
) *AuthH {
	return &AuthH{
		services: services,
		cfg:      cfg,
	}
}

func (h *AuthH) Login(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	case "POST":
		req := new(requests.LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		user, err := h.services.User.GetByEmail(req.Email)
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

		accessString, err := h.services.Token.GenerateAccessToken(user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		refreshString, err := h.services.Token.GenerateRefreshToken(user.Id)
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

func (h *AuthH) Refresh(w http.ResponseWriter, r *http.Request) {
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

		accessString, refreshString, err := h.services.Token.RefreshToken(req.AccessToken, req.RefreshToken, h.cfg.AccessSecret, h.cfg.RefreshSecret)
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
