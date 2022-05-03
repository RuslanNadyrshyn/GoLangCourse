package middleware

import (
	"Delivery/Delivery/internal/services"
	"net/http"
)

type Middleware struct {
	service *services.TokenService
}

func NewMiddleware(service *services.TokenService) *Middleware {
	return &Middleware{
		service: service,
	}
}

func (m *Middleware) AuthCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := m.service.GetTokenFromBearerString(r.Header.Get("Authorization"))

		_, err := m.service.ValidateAccessToken(accessToken)
		if err != nil {
			http.Error(w, "(middleware) not authorized", http.StatusUnauthorized)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
