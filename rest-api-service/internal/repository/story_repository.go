package repository

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/cache"
	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/model"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type storyRepository struct {
	db          *sql.DB
	redisClient *cache.RedisClient
}

func NewStoryRepository(db *sql.DB, redisClient *cache.RedisClient) model.StoryRepository {
	return &storyRepository{
		db:          db,
		redisClient: redisClient,
	}
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

	// TODO: insert dta to redis

	return &data, nil
}

func (r *storyRepository) FindAll(ctx context.Context, opt model.StoryOpt) (results []*model.Story, total int64, err error) {
	logger := logrus.WithFields(
		logrus.Fields{
			"opt": opt,
		},
	)

	// TODO: check is data stored in redis?
	// if yes, get data from redis and return directly
	bucketKey := storyListBucketKey
	cacheKey := newStoryListCacheKey(opt)
	err = r.redisClient.HGet(ctx, bucketKey, cacheKey, &results)
	if err != nil && err != redis.Nil {
		logger.Errorf("get data from redis: %s", err.Error())
	}
	if len(results) > 0 {
		return
	}

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

	// Store data to redis
	err = r.redisClient.HSet(ctx, bucketKey, cacheKey, results)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *storyRepository) FindByID(ctx context.Context, id int64) (story *model.Story, err error) {
	logger := logrus.WithFields(
		logrus.Fields{
			"id": id,
		},
	)

	cacheKey := newStoryDetailKey(id)
	err = r.redisClient.Get(ctx, cacheKey, &story)
	if err != nil && err != redis.Nil {
		logger.Errorf("get data from redis: %s", err.Error())
	}
	// if data is found in redis, return it
	if story.ID > 0 {
		return story, nil
	}

	// continue select from db
	row := sq.Select("id, title, content, author_id, created_at").
		From("stories").
		Where(sq.Eq{"id": id}).RunWith(r.db).QueryRowContext(ctx)
	if err = row.Scan(
		&story.ID,
		&story.Title,
		&story.Content,
		&story.AuthorID,
		&story.CreatedAt); err != nil {
		logger.Errorf("scan data from db: %s", err.Error())
		return
	}

	// Store data to redis
	err = r.redisClient.Set(ctx, cacheKey, story, 0)
	if err != nil {
		logger.Errorf("set data to redis: %s", err.Error())
	}

	return
}

func (r *storyRepository) Update(ctx context.Context, id int64) error {
	// Run Update query

	// Invalidate cache
	r.invalidateCache(ctx, id)

	return nil
}

func (r *storyRepository) Delete(ctx context.Context, id int64) error {
	// Run Delete query

	// Invalidate cache
	r.invalidateCache(ctx, id)

	return nil
}

func (r *storyRepository) invalidateCache(ctx context.Context, id int64) {
	r.redisClient.Del(ctx, newStoryDetailKey(id))
	r.redisClient.HDel(ctx, storyListBucketKey)
}
