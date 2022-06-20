package middleware

import (
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/repositories/requests"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/services"
	"net/http"
)

type Middleware struct {
	service *services.ServiceManager
}

func NewMiddleware(service *services.ServiceManager) *Middleware {
	return &Middleware{
		service: service,
	}
}

func (m *Middleware) Validation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.SetupCORS(&w, r)
		switch r.Method {
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
		case "GET":
			accessToken := m.service.Token.GetTokenFromBearerString(r.Header.Get("Authorization"))
			_, err := m.service.Token.ValidateAccessToken(accessToken)
			if err != nil {
				http.Error(w, "(middleware) not authorized", http.StatusUnauthorized)
				return
			}
			w.Header().Add("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		}
	})
}
