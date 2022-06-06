package handlers

import (
	"Delivery/backend/internal/repositories/models"
	"Delivery/backend/internal/repositories/requests"
	"Delivery/backend/internal/repositories/responses"
	"Delivery/backend/internal/services"
	"encoding/json"
	"log"
	"net/http"
)

type UserHandler struct {
	services *services.ServiceManager
}

func NewUserHandler(services *services.ServiceManager) *UserHandler {
	return &UserHandler{
		services: services,
	}
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	case "POST":
		var id int64
		req := new(models.User)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			log.Println(err)
			return
		}

		_, err := h.services.User.GetByEmail(req.Email)
		if err != nil {
			id, err = h.services.User.Insert(req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				log.Println(err)
				return
			}
		} else {
			http.Error(w, "User with this email already exists", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		resp := responses.UserResponse{
			ID:    id,
			Name:  req.Name,
			Email: req.Email,
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

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	case "GET":
		requestToken := h.services.Token.GetTokenFromBearerString(r.Header.Get("Authorization"))
		claims, err := h.services.Token.ValidateAccessToken(requestToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			log.Println(err)
			return
		}

		user, err := h.services.User.GetById(claims.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}

		resp := responses.UserResponse{
			ID:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Println(err)
			return
		}
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
