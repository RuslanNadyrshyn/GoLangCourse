package services

import (
	"awesomeProject/internal/auth/config"
	"awesomeProject/internal/auth/helper"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type TokenService struct {
	cfg *config.Config
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{
		cfg: cfg,
	}
}

func (s *TokenService) GenerateAccessToken(userID int) (string, error) {
	return s.GenerateToken(userID, s.cfg.AccessLifetimeMinutes, s.cfg.AccessSecret)
}

func (s *TokenService) GenerateRefreshToken(userID int) (string, error) {
	return s.GenerateToken(userID, s.cfg.RefreshLifetimeMinutes, s.cfg.RefreshSecret)
}

func (s *TokenService) GenerateToken(userID, lifetimeMinutes int, secret string) (string, error) {
	claims := &helper.JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifetimeMinutes)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (s *TokenService) ValidateAccessToken(tokenString string) (*helper.JwtCustomClaims, error) {
	return s.ValidateToken(tokenString, s.cfg.AccessSecret)
}

func (s *TokenService) ValidateRefreshToken(tokenString string) (*helper.JwtCustomClaims, error) {
	return s.ValidateToken(tokenString, s.cfg.RefreshSecret)
}

func (s *TokenService) ValidateToken(tokenString, secret string) (*helper.JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &helper.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*helper.JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse token claims")
	}

	return claims, nil
}

func (s *TokenService) GetTokenFromBearerString(bearerString string) string {
	if bearerString == "" {
		return ""
	}

	parts := strings.Split(bearerString, "Bearer ")
	if len(parts) != 2 {
		return ""
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}

	return token
}

func (s *TokenService) ParseToken(tokenString, secret string) (claims *helper.JwtCustomClaims, err error,
) {
	token, err := jwt.ParseWithClaims(tokenString, &helper.JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(*helper.JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, nil
}

func (s *TokenService) RefreshToken(accessToken, refreshToken, accessSecret, refreshSecret string) (string, string, error) {
	//claims, err := s.ParseToken(tokenString, secret)
	//if err != nil {
	//	return "", errors.New("not authorized")
	//}

	accToken, err := s.ParseToken(accessToken, accessSecret)
	if err != nil {
		return "", "", errors.New("not authorized")
	}
	refToken, err := s.ParseToken(refreshToken, refreshSecret)
	if err != nil {
		return "", "", errors.New("not authorized")
	}

	newAccessToken, err := s.GenerateAccessToken(accToken.ID)
	if err != nil {
		return "", "", err
	}
	newRefreshToken, err := s.GenerateRefreshToken(refToken.ID)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}
