package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"pms.pkg/logger"
)

type Config struct {
	Host     string `env:"HOST"`
	Password string `env:"PASSWORD"`
}

type Client[T Cachable] struct {
	r *redis.Client
}

func New[T Cachable](conf *Config, t T) *Client[T] {
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
	return c.r.Set(ctx, key, j, time.Duration(exp)*time.Hour).Err()
}

func (c *Client[T]) Rpush(ctx context.Context, key string, value T) error {
	j, err := json.Marshal(value)
	if err != nil {
		return err
	}
	cmd := c.r.RPush(ctx, key, j)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (c *Client[T]) Rpop(ctx context.Context, key string) (result T, err error) {
	cmd := c.r.RPop(ctx, key)
	if cmd.Err() != nil {
		return result, cmd.Err()
	}
	val, err := cmd.Bytes()
	if err != nil {
		return result, err
	}
	if err = json.Unmarshal(val, &result); err != nil {
		return result, err
	}
	return result, nil
}

func (c *Client[T]) Delete(ctx context.Context, key string) error {
	return c.r.Del(ctx, key).Err()
}

func (c *Client[T]) Get(ctx context.Context, key string) (T, error) {
	log := logger.Log.With("func", "Get")
	log.Debug("Get called")

	var t T
	val := c.r.Get(ctx, key).Val()
	log.Debugw("raw data from redis", "val", val, "key", key)
	if err := json.Unmarshal([]byte(val), &t); err != nil {
		return t, err
	}
	return t, nil
}
