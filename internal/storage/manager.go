package storage

import (
	"context"
	"gons/pkg/utils/storage"
	"gons/internal/contracts"
	"gons/pkg/env"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func customEndpointResolver(endpoint string) aws.EndpointResolverWithOptionsFunc {
	return func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if endpoint != "" {
			return aws.Endpoint{
				URL:           endpoint,
				SigningRegion: region,
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	}
}

// NewStorage returns a new Storage contract implementation based on environment configuration.
func NewStorage() contracts.Storage {
	disk := env.Get("STORAGE_DISK", "local")

	if disk == "s3" {
		key := env.Get("S3_KEY", "")
		secret := env.Get("S3_SECRET", "")
		region := env.Get("S3_REGION", "us-east-1")
		bucket := env.Get("S3_BUCKET", "")
		endpoint := env.Get("S3_ENDPOINT", "")
		usePathStyle, _ := strconv.ParseBool(env.Get("S3_USE_PATH_STYLE", "false"))

		cfg, _ := awsconfig.LoadDefaultConfig(context.TODO(),
			awsconfig.WithRegion(region),
			awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, "")),
			awsconfig.WithEndpointResolverWithOptions(customEndpointResolver(endpoint)),
		)

		client := s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = usePathStyle
		})

		return &storage.S3Driver{
			Client:       client,
			BucketName:   bucket,
			Region:       region,
			BaseURL:      env.Get("STORAGE_BASE_URL", ""),
			UsePathStyle: usePathStyle,
		}
	}

	// Default Local
	return &storage.LocalDriver{
		Root:    env.Get("STORAGE_ROOT", "./public/storage"),
		BaseURL: env.Get("STORAGE_BASE_URL", env.Get("APP_URL", "http://localhost:8080")+"/storage"),
	}
}
