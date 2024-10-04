package model

import "context"

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
