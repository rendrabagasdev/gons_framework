package cache

import (
	"errors"
	"sync"
	"time"
)

type item struct {
	value      string
	expiration int64
}

type MemoryDriver struct {
	store sync.Map
}

func (m *MemoryDriver) Get(key string) (string, error) {
	val, ok := m.store.Load(key)
	if !ok {
		return "", errors.New("cache: miss")
	}

	cachedItem := val.(item)
	if time.Now().Unix() > cachedItem.expiration {
		m.store.Delete(key)
		return "", errors.New("cache expired")
	}

	return cachedItem.value, nil
}

func (m *MemoryDriver) Set(key string, value string, ttl time.Duration) error {
	m.store.Store(key, item{
		value:      value,
		expiration: time.Now().Add(ttl).UnixNano(),
	})
	return nil
}

func (m *MemoryDriver) Forget(key string) error {
	m.store.Delete(key)
	return nil
}

func (m *MemoryDriver) Remember(key string, ttl time.Duration, callback func() (string, error)) (string, error) {
	val, err := m.Get(key)
	if err == nil {
		return val, nil
	}

	freshData, err := callback()
	if err != nil {
		return "", err
	}

	m.Set(key, freshData, ttl)

	return freshData, nil
}
