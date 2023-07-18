package domain

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// JwtCustomClaims representa los claims personalizados que se incluyen en el JWT de la plataforma .
type JwtCustomClaims struct {
	Email  string    `json:"email"`
	UserID uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}

// JwtCustomRefreshClaims representa los claims personalizados que se incluyen en el JWT de actualización en la plataforma .
type JwtCustomRefreshClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}
