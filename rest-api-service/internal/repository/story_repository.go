package repository

import (
	"context"
	"database/sql"

	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/model"
)

type storyRepository struct {
	db *sql.DB
}

func NewStoryRepository(db *sql.DB) model.StoryRepository {
	return &storyRepository{db: db}
}

func (r *storyRepository) Create(ctx context.Context, data model.Story) (*model.Story, error) {
	// TODO
	panic("implement me!")
}

func (r *storyRepository) FindAll(ctx context.Context, opt model.StoryOpt) (results []*model.Story, total int64, err error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title FROM articles")
	if err != nil {
		return
	}

	for rows.Next() {
		var story model.Story
		if err := rows.Scan(&story.ID, &story.Title); err != nil {
			// TODO: use log
			continue
		}
		results = append(results, &story)
	}

	return
}

func (r *storyRepository) FindByID(ctx context.Context, id int64) (*model.Story, error) {
	// TODO
	panic("implement me!")
}

func (r *storyRepository) Update(ctx context.Context, id int64) error {
	// TODO
	panic("implement me!")
}

func (r *storyRepository) Delete(ctx context.Context, id int64) error {
	// TODO
	panic("implement me!")
}
