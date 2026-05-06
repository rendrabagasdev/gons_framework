package cache

import (
	"context"
	"gons/internal/contracts"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisDriver struct {
	Client *redis.Client
}

func (r *RedisDriver) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (r *RedisDriver) Set(key string, value string, ttl time.Duration) error {
	ctx := context.Background()
	return r.Client.Set(ctx, key, value, ttl).Err()
}

func (r *RedisDriver) Forget(key string) error {
	ctx := context.Background()
	return r.Client.Del(ctx, key).Err()
}

func (r *RedisDriver) Remember(key string, ttl time.Duration, callback func() (string, error)) (string, error) {
	val, err := r.Get(key)
	if err == nil && val != "" {
		return val, nil
	}

	freshVal, err := callback()
	if err != nil {
		return "", err
	}

	err = r.Set(key, freshVal, ttl)
	return freshVal, err
}

var _ contracts.Cache = (*RedisDriver)(nil)
