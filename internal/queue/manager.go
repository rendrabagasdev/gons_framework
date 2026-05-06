package queue

import (
	"gons/internal/contracts"
	"gons/pkg/env"
	"gons/pkg/utils/queue"
	"log/slog"
	"strconv"

	"github.com/golobby/container/v3"
	"github.com/redis/go-redis/v9"
)

// NewQueue returns a new Queue contract implementation based on environment configuration.
func NewQueue() contracts.Queue {
	driver := env.Get("QUEUE_DRIVER", "sync")
	bufferSizeStr := env.Get("QUEUE_BUFFER_SIZE", "100")
	bufferSize, _ := strconv.Atoi(bufferSizeStr)

	if driver == "redis" {
		var client *redis.Client
		if err := container.Resolve(&client); err != nil {
			slog.Error("Gons: Failed to resolve Redis client for Queue", "error", err)
			return queue.NewGoroutineDriver(bufferSize)
		}
		return &RedisDriver{Client: client}
	}

	return queue.NewGoroutineDriver(bufferSize)
}
