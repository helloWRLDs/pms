package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	cachemodels "pms.auth/internal/modules/cache/models"
)

type Config struct {
	Host     string `env:"HOST"`
	Password string `env:"PASSWORD"`
}

type Client[T cachemodels.Cachable] struct {
	r *redis.Client
}

func New[T cachemodels.Cachable](conf Config, t T) *Client[T] {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		Password: conf.Password,
		DB:       t.GetDB(),
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Errorf("redis conn failed: %w", err))
	}

	return &Client[T]{
		r: rdb,
	}
}

func (c *Client[T]) Set(ctx context.Context, key string, t T, exp int64) error {
	j, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return c.r.Set(ctx, key, j, time.Duration(exp)*time.Second).Err()
}

func (c *Client[T]) Get(ctx context.Context, key string) (T, error) {
	var t T
	val, err := c.r.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return t, nil
		}
		return t, err
	}
	if err := json.Unmarshal([]byte(val), &t); err != nil {
		return t, err
	}
	return t, nil
}
