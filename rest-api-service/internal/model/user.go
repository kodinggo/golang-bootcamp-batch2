package model

import (
	"context"
	"time"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, data User) (*User, error)
}

type Role string

const (
	AuthorRole Role = "AUTHOR"
	AdminRole  Role = "ADMIN"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-"`
}
