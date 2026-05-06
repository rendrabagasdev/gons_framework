package database

import (
	"fmt"
	"gons/pkg/env"
	"log/slog"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// NewRedisClient initializes and returns a Redis client.
func NewRedisClient() *redis.Client {
	host := env.Get("REDIS_HOST", "127.0.0.1")
	port := env.Get("REDIS_PORT", "6379")
	password := env.Get("REDIS_PASSWORD", "")
	dbStr := env.Get("REDIS_DB", "0")
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		db = 0
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	slog.Info("Redis initialized", "host", host, "port", port, "db", db)

	return rdb
}