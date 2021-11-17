package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type HM map[string]int

type Cache interface {
	HMSet(ctx context.Context, key string, value ...string) (bool, error)
	HMGetInt(ctx context.Context, key string, fields ...string) ([]int, error)
	Expire(ctx context.Context, key string, d time.Duration) (bool, error)
}

type Config struct {
	Addr string `json:"addr"`
	Pass string `json:"password"`
	DB   int    `json:"db"`
}

type cache struct {
	redis *redis.Client
}

func NewCache(cfg *Config) (*cache, func(), error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Pass,
		DB:       cfg.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, nil, err
	}

	return &cache{
			redis: rdb,
		}, func() {
			rdb.Close()
		}, nil
}

func (c *cache) Expire(ctx context.Context, key string, d time.Duration) (bool, error) {
	return c.redis.Expire(ctx, key, d).Result()
}

func (c *cache) HMSet(ctx context.Context, key string, value ...string) (bool, error) {
	return c.redis.HMSet(ctx, key, value).Result()
}

func (c *cache) HMGetInt(ctx context.Context, key string, fields ...string) ([]int, error) {
	res, err := c.redis.HMGet(ctx, key, fields...).Result()
	if err != nil {
		return []int{}, err
	}

	sliceInt := make([]int, len(res))
	for i, r := range res {
		if n, err := strconv.Atoi(fmt.Sprintf("%s", r)); err == nil {
			sliceInt[i] = n
			continue
		}
		return []int{}, err
	}

	return sliceInt, nil
}
