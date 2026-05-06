package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gons/internal/contracts"
)

var _ contracts.Storage = (*LocalDriver)(nil)

type LocalDriver struct {
	Root    string
	BaseURL string
}

func (l *LocalDriver) Bucket(name string) contracts.Storage {
	// Untuk lokal, bucket dianggap sebagai sub-folder di dalam root
	return &LocalDriver{
		Root:    filepath.Join(l.Root, name),
		BaseURL: fmt.Sprintf("%s/%s", strings.TrimRight(l.BaseURL, "/"), name),
	}
}

func (l *LocalDriver) Put(filePath string, content io.Reader) error {
	fullPath := filepath.Join(l.Root, filePath)

	if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	return err
}

func (l *LocalDriver) Get(filePath string) ([]byte, error) {
	fullPath := filepath.Join(l.Root, filePath)
	return os.ReadFile(fullPath)
}

func (l *LocalDriver) Exits(filePath string) bool {
	fullPath := filepath.Join(l.Root, filePath)
	_, err := os.Stat(fullPath)
	return err == nil
}

func (l *LocalDriver) Delete(filePath string) error {
	fullPath := filepath.Join(l.Root, filePath)
	return os.Remove(fullPath)
}

func (l *LocalDriver) Url(filePath string) string {
	baseUrl := strings.TrimRight(l.BaseURL, "/")
	cleanPath := strings.TrimLeft(filePath, "/")
	return fmt.Sprintf("%s/%s", baseUrl, cleanPath)
}
