package services

import (
	"Delivery/backend/internal/helper"
	"Delivery/backend/internal/repositories/models"
	"Delivery/backend/internal/repositories/requests"
	"Delivery/backend/internal/repositories/responses"
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
	Insert(u *models.User) (int64, error)
	GetByEmail(email string) (*models.User, error)
	GetById(id int64) (*models.User, error)
	Update(u *models.User) (int64, error)
}

type ISupplierService interface {
	Insert(s *requests.SupplierRequest) (int64, error)
	GetAll() (*[]models.Supplier, error)
	GetByType(t string) (*[]models.Supplier, error)
	GetByProductId(prodId int64) (*models.Supplier, error)
	GetTypes() (*[]string, error)
}

type IProductService interface {
	GetAll() (*[]models.Product, error)
	GetById(id int64) (*models.Product, error)
	GetByParams(params requests.SortRequest) (*[]models.Product, error)
	GetTypes(params requests.SortRequest) (*[]string, error)
	UpdatePrice(p *models.Product) (int64, error)
}

type IOrderService interface {
	AddOrder(order *requests.OrderRequest) (int64, error)
	GetById(id int64) (*responses.OrderResponse, error)
	GetByUserId(id int64) (*[]models.Order, error)
}
