package handlers

import (
	"Delivery/Delivery/internal/repositories"
	"Delivery/Delivery/internal/repositories/database"
	"Delivery/Delivery/internal/repositories/models"
	"Delivery/Delivery/internal/repositories/requests"
	"Delivery/Delivery/internal/repositories/responses"
	"Delivery/Delivery/internal/services"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	tokenService   *services.TokenService
	userRepository repositories.IUserRepository
	conn           *sql.DB
}

func NewUserHandler(
	tokenService *services.TokenService,
	userRepository repositories.IUserRepository,
	conn *sql.DB,
) *UserHandler {
	return &UserHandler{
		tokenService:   tokenService,
		userRepository: userRepository,
		conn:           conn,
	}
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	case "POST":
		var id int
		req := new(models.User)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		userRepo := database.NewUserRepository(h.conn)

		_, err := userRepo.GetByEmail(req.Email)
		if err != nil {
			id, err = userRepo.Insert(req)
			if err != nil {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
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
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "GET":
		requestToken := h.tokenService.GetTokenFromBearerString(r.Header.Get("Authorization"))
		claims, err := h.tokenService.ValidateAccessToken(requestToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		user, err := h.userRepository.GetUserByID(claims.ID)
		if err != nil {
			http.Error(w, "User does not exist", http.StatusBadRequest)
			return
		}

		resp := responses.UserResponse{
			ID:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	requests.SetupCORS(&w, r)
	switch r.Method {
	case "GET":
		req := r.URL.Query().Get("id")
		id, err := strconv.Atoi(req)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		userRepo := database.NewUserRepository(h.conn)

		user, err := userRepo.GetById(id)
		if err != nil {
			http.Error(w, "User does not exist", http.StatusBadRequest)
			return
		}

		resp := responses.UserResponse{
			ID:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
