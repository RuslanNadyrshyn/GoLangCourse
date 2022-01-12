package main

import (
	"awesomeProject/internal/auth"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

const (
	httpPort = ":8080"

	accessSecret  = "access_secret_string"
	refreshSecret = "refresh_secret_string"

	accessLifetimeMinutes  = 5
	refreshLifetimeMinutes = 60
)

func main() {
	http.HandleFunc("/login", Login)
	//http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/profile", GetProfile)

	log.Fatal(http.ListenAndServe(httpPort, nil))
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := new(auth.LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := auth.NewUserRepository().GetUserByEmail(req.Email)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		accessString, err := auth.GenerateToken(user.ID, accessLifetimeMinutes, accessSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		refreshString, err := auth.GenerateToken(user.ID, refreshLifetimeMinutes, refreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := auth.LoginResponse{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		claims, err := auth.ValidateToken(auth.GetTokenFromBearerString(r.Header.Get("Authorization")), refreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		user, err := auth.NewUserRepository().GetUserByID(claims.ID)
		if err != nil {
			http.Error(w, "User does not exist", http.StatusBadRequest)
			return
		}

		resp := auth.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
