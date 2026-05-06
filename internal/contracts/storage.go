package contracts

import (
	"io"
)

type Storage interface {
	Bucket(name string) Storage // Untuk ganti bucket saat runtime
	Put(path string, content io.Reader) error
	Get(path string) ([]byte, error)
	Exits(path string) bool
	Delete(path string) error
	Url(path string) string
}
