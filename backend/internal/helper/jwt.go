package helper

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	ID int64 `json:"id"`
	jwt.StandardClaims
}
