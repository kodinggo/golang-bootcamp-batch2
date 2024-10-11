package model

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type ContextAuthKey string

const BearerAuthKey ContextAuthKey = "BearerAuth"

type CustomClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

type AuthUsecase interface {
	Login(ctx context.Context, data Login) (string, error)
	Register(ctx context.Context, data Register) (string, error)
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
