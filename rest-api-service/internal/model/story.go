package model

import (
	"context"
	"errors"
	"strings"
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

const DefaultLimit = 10

var AllowedRolesToManageStory = map[Role]bool{
	AdminRole:  true,
	AuthorRole: true,
}

type Story struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  int64     `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

type StoryOpt struct {
	SortBy  string `query:"sort_by"`
	Keyword string `query:"keyword"`
	Cursor  string `query:"cursor"`
	Limit   int64  `query:"limit"`
}

func (opt *StoryOpt) Validate() error {
	switch strings.ToLower(opt.SortBy) {
	case "", "asc", "desc":
	default:
		return errors.New("parameter sort_by is not valid")
	}

	if opt.Limit <= 0 || opt.Limit > 100 {
		opt.Limit = DefaultLimit
	}

	return nil
}
