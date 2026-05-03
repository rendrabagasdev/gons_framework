package storage

import (
	"context"
	"fmt"
	"io"
	"strings"

	"go-framework/app/contracts"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var _ contracts.Storage = (*S3Driver)(nil)

type S3Driver struct {
	Client       *s3.Client
	BucketName   string // Diganti agar tidak konflik dengan method Bucket()
	Region       string
	BaseURL      string
	UsePathStyle bool
}

func (s *S3Driver) Bucket(name string) contracts.Storage {
	// Return instance baru dengan bucket berbeda
	return &S3Driver{
		Client:       s.Client,
		BucketName:   name,
		Region:       s.Region,
		BaseURL:      s.BaseURL,
		UsePathStyle: s.UsePathStyle,
	}
}

func (s *S3Driver) Put(filePath string, content io.Reader) error {
	_, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(strings.TrimLeft(filePath, "/")),
		Body:   content,
	})
	return err
}

func (s *S3Driver) Get(filePath string) ([]byte, error) {
	result, err := s.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(strings.TrimLeft(filePath, "/")),
	})
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	return io.ReadAll(result.Body)
}

func (s *S3Driver) Exits(filePath string) bool {
	_, err := s.Client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(strings.TrimLeft(filePath, "/")),
	})
	return err == nil
}

func (s *S3Driver) Delete(filePath string) error {
	_, err := s.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(strings.TrimLeft(filePath, "/")),
	})
	return err
}

func (s *S3Driver) Url(filePath string) string {
	if s.BaseURL != "" {
		baseUrl := strings.TrimRight(s.BaseURL, "/")
		cleanPath := strings.TrimLeft(filePath, "/")
		return fmt.Sprintf("%s/%s", baseUrl, cleanPath)
	}

	// Default S3 URL format
	if s.UsePathStyle {
		return fmt.Sprintf("https://s3.%s.amazonaws.com/%s/%s", s.Region, s.BucketName, strings.TrimLeft(filePath, "/"))
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.BucketName, s.Region, strings.TrimLeft(filePath, "/"))
}
