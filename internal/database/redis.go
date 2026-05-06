package database

import (
	"context"
	"fmt"
	"gons/pkg/env"
	"log/slog"
	"os"
	"strconv"
	"time"

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

	// Verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		slog.Error("Failed to connect to Redis", "error", err, "host", host, "port", port)
		os.Exit(1)
	}

	slog.Info("Redis connected successfully", "host", host, "port", port, "db", db)

	return rdb
}
