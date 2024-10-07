package repository

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/model"
	"github.com/sirupsen/logrus"
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
	logger := logrus.WithFields(
		logrus.Fields{
			"opt": opt,
		},
	)

	selectQ := sq.Select("id, title, content, author_id, created_at").
		From("stories")

	if opt.Keyword != "" {
		selectQ = selectQ.Where(sq.Like{"title": fmt.Sprintf("%%%s%%", opt.Keyword)})
	}

	if opt.Cursor != "" {
		selectQ = selectQ.Where(sq.LtOrEq{"id": opt.Cursor})
	}

	if opt.SortBy != "" {
		selectQ = selectQ.OrderBy("id DESC")
	}

	rowSQL, _, _ := selectQ.ToSql()
	fmt.Println(rowSQL)

	rows, err := selectQ.RunWith(r.db).QueryContext(ctx)
	if err != nil {
		logger.Error(err)
		return
	}

	for rows.Next() {
		var story model.Story
		if err := rows.Scan(
			&story.ID,
			&story.Title,
			&story.Content,
			&story.AuthorID,
			&story.CreatedAt); err != nil {
			logger.Error(err)
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
