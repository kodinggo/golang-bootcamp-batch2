package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/config"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisClient struct {
	client *redis.Client
}

func InitRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost(),
		Password: config.RedisPass(),
		DB:       config.RedisDB(),
	})

	return &RedisClient{
		client: client,
	}
}

func (r *RedisClient) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	err := r.client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		logrus.Errorf("set data to redis: %v", err.Error())
	}
	return err
}

func (r *RedisClient) Get(ctx context.Context, key string, data any) error {
	strData, err := r.client.Get(ctx, key).Result()
	if err != nil {
		logrus.Errorf("get data from redis: %v", err.Error())
		return err
	}
	json.Unmarshal([]byte(strData), &data)
	return nil
}

func (r *RedisClient) Del(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		logrus.Errorf("delete data from redis: %v", err.Error())
	}
	return err
}

func (r *RedisClient) HSet(ctx context.Context, bucketKey string, cacheKey string, data any) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		logrus.Errorf("marshal data to storing redis: %v", err.Error())
		return err
	}
	err = r.client.HSet(ctx, bucketKey, cacheKey, byteData).Err()
	if err != nil {
		logrus.Errorf("set data hash to redis: %v", err.Error())
	}
	return err
}

func (r *RedisClient) HGet(ctx context.Context, bucketKey, cacheKey string, data any) error {
	strData, err := r.client.HGet(ctx, bucketKey, cacheKey).Result()
	if err != nil {
		logrus.Errorf("get data hash from redis: %v", err.Error())
		return err
	}
	json.Unmarshal([]byte(strData), &data)
	return err
}

func (r *RedisClient) HDel(ctx context.Context, bucketKey string) error {
	err := r.client.HDel(ctx, bucketKey).Err()
	if err != nil {
		logrus.Errorf("delete data hash from redis: %v", err.Error())
	}
	return err
}
