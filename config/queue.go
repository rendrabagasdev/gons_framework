package config

import (
	"go-framework/app/contracts"
	"go-framework/app/utils/queue"
	"strconv"

	"github.com/golobby/container/v3"
)

func init() {
	RegisterConfig(func() error {
		return container.Singleton(func() contracts.Queue {
			driver := GetEnv("QUEUE_DRIVER", "sync")
			bufferSizeStr := GetEnv("QUEUE_BUFFER_SIZE", "100")
			bufferSize, _ := strconv.Atoi(bufferSizeStr)

			if driver == "redis" {
				// return
			}

			return queue.NewGoroutineDriver(bufferSize)
		})
	})
}
