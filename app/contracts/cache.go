package contracts

import "time"

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string, ttl time.Duration) error
	Forget(key string) error
	Remember(key string, ttl time.Duration, callback func() (string, error)) (string, error)
}