package model

import (
	"context"
	"time"
)

type StoryRepository interface {
	Create(ctx context.Context, data Story) (*Story, error)
	FindAll(ctx context.Context, opt StoryOpt) ([]*Story, int64, error)
	FindByID(ctx context.Context, id int64) (*Story, error)
	Update(ctx context.Context, id int64) error
	Delete(ctx context.Context, id int64) error
}

type StoryUsecase interface {
	Create(ctx context.Context, data Story) (*Story, error)
	FindAll(ctx context.Context, opt StoryOpt) ([]*Story, int64, error)
	FindByID(ctx context.Context, id int64) (*Story, error)
	Update(ctx context.Context, id int64) error
	Delete(ctx context.Context, id int64) error
}

type Story struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  int64     `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-"`
}

type StoryOpt struct {
	SortBy  string `query:"sort_by"`
	Keyword string `query:"keyword"`
}
