package repository

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"

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
	timeNowUTC := time.Now().UTC()

	result, err := sq.Insert("stories").
		Columns("title", "content", "author_id", "created_at").
		Values(
			data.Title,
			data.Content,
			data.AuthorID,
			timeNowUTC).
		RunWith(r.db).
		ExecContext(ctx)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{
				"opt": data,
			},
		).Errorf("run query insert: %s", err.Error())

		return nil, err
	}

	newStoryID, _ := result.LastInsertId()
	data.ID = newStoryID
	data.CreatedAt = timeNowUTC

	return &data, nil
}

func (r *storyRepository) FindAll(ctx context.Context, opt model.StoryOpt) (results []*model.Story, total int64, err error) {
	logger := logrus.WithFields(
		logrus.Fields{
			"opt": opt,
		},
	)

	selectQ := sq.Select("id, title, content, author_id, created_at").
		From("stories")

	selectCount := sq.Select("count(id) as total").From("stories")

	if opt.Keyword != "" {
		selectQ = selectQ.Where(sq.Like{"title": fmt.Sprintf("%%%s%%", opt.Keyword)})
		selectCount = selectCount.Where(sq.Like{"title": fmt.Sprintf("%%%s%%", opt.Keyword)})
	}

	if opt.Cursor != "" {
		// Decode from base64 to string id
		byteCursor, err := base64.StdEncoding.DecodeString(opt.Cursor)
		if err != nil {
			return nil, 0, err
		}
		var idStr = string(byteCursor)
		selectQ = selectQ.Where(sq.Lt{"id": idStr})
	}

	if opt.SortBy != "" {
		selectQ = selectQ.OrderBy(fmt.Sprintf("id %s", opt.SortBy))
	}

	selectQ = selectQ.Limit(uint64(opt.Limit))

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

	row := selectCount.RunWith(r.db).QueryRowContext(ctx)
	row.Scan(&total)

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
