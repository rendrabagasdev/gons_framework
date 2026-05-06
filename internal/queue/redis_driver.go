package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"gons/internal/contracts"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisDriver struct {
	Client *redis.Client
}

// Push adds a new job to the end of the queue.
func (r *RedisDriver) Push(name string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("queue: failed to marshal payload: %w", err)
	}
	return r.Client.LPush(context.Background(), name, data).Err()
}

// Later schedules a job to be executed after a delay.
func (r *RedisDriver) Later(name string, payload any, delay time.Duration) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("queue: failed to marshal payload: %w", err)
	}
	
	at := time.Now().Add(delay).Unix()
	return r.Client.ZAdd(context.Background(), name+":delayed", redis.Z{
		Score:  float64(at),
		Member: data,
	}).Err()
}

// Pop retrieves and removes the next job from the queue.
func (r *RedisDriver) Pop(name string) (string, error) {
	ctx := context.Background()
	
	// Try to move any delayed jobs that are now ready
	r.moveDelayedJobs(name)

	val, err := r.Client.RPop(ctx, name).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

// moveDelayedJobs moves jobs from the delayed sorted set to the main list when they are ready.
func (r *RedisDriver) moveDelayedJobs(name string) {
	ctx := context.Background()
	delayedKey := name + ":delayed"
	now := time.Now().Unix()

	// Find jobs that are ready to be processed
	jobs, err := r.Client.ZRangeByScore(ctx, delayedKey, &redis.ZRangeBy{
		Min: "-inf",
		Max: fmt.Sprintf("%d", now),
	}).Result()

	if err != nil || len(jobs) == 0 {
		return
	}

	for _, job := range jobs {
		// Use a transaction or simple atomicity to move the job
		pipe := r.Client.Pipeline()
		pipe.LPush(ctx, name, job)
		pipe.ZRem(ctx, delayedKey, job)
		_, _ = pipe.Exec(ctx)
	}
}

var _ contracts.Queue = (*RedisDriver)(nil)
