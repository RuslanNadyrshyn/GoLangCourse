package services

import (
	"Delivery/Delivery/internal/helper"
	"Delivery/Delivery/internal/repositories/models"
	"Delivery/Delivery/internal/repositories/requests"
	"Delivery/Delivery/internal/repositories/responses"
)

type ITokenService interface {
	GenerateAccessToken(userID int64) (string, error)
	GenerateRefreshToken(userID int64) (string, error)
	GenerateToken(userID int64, lifetimeMinutes int, secret string) (string, error)
	ValidateAccessToken(tokenString string) (*helper.JwtCustomClaims, error)
	ValidateRefreshToken(tokenString string) (*helper.JwtCustomClaims, error)
	ValidateToken(tokenString, secret string) (*helper.JwtCustomClaims, error)
	GetTokenFromBearerString(bearerString string) string
	ParseToken(tokenString, secret string) (*helper.JwtCustomClaims, error)
	RefreshToken(accessToken, refreshToken, accessSecret, refreshSecret string) (string, string, error)
}

type IUserService interface {
	Insert(u *models.User) (id int64, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetById(id int64) (user *models.User, err error)
	Update(u *models.User) (id int64, err error)
}

type ISupplierService interface {
	Insert(s *requests.SupplierRequest) (int64, error)
	GetAll() (*[]requests.SupplierRequest, error)
}

type IProductService interface {
	GetAll() (*[]models.Product, error)
	GetById(id int64) (*responses.ProductResponse, error)
	UpdatePrice(p *models.Product) (int64, error)
}

type IOrderService interface {
	AddOrder(order *requests.OrderRequest) (int64, error)
	GetById(id int64) (*responses.OrderResponse, error)
	GetByUserId(id int64) (*[]models.Order, error)
}
