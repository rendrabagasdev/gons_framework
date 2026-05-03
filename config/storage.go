package config

import (
	"context"
	"gons/app/contracts"
	"gons/app/utils/storage"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golobby/container/v3"
)

func init() {
	RegisterConfig(func() error {
		return container.Singleton(func() contracts.Storage {
			disk := GetEnv("STORAGE_DISK", "local")

			if disk == "s3" {
				key := GetEnv("S3_KEY", "")
				secret := GetEnv("S3_SECRET", "")
				region := GetEnv("S3_REGION", "us-east-1")
				bucket := GetEnv("S3_BUCKET", "")
				endpoint := GetEnv("S3_ENDPOINT", "")
				usePathStyle, _ := strconv.ParseBool(GetEnv("S3_USE_PATH_STYLE", "false"))

				customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					if endpoint != "" {
						return aws.Endpoint{
							URL:           endpoint,
							SigningRegion: region,
						}, nil
					}
					return aws.Endpoint{}, &aws.EndpointNotFoundError{}
				})

				cfg, _ := config.LoadDefaultConfig(context.TODO(),
					config.WithRegion(region),
					config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, "")),
					config.WithEndpointResolverWithOptions(customResolver),
				)

				client := s3.NewFromConfig(cfg, func(o *s3.Options) {
					o.UsePathStyle = usePathStyle
				})

				return &storage.S3Driver{
					Client:       client,
					BucketName:   bucket,
					Region:       region,
					BaseURL:      GetEnv("STORAGE_BASE_URL", ""),
					UsePathStyle: usePathStyle,
				}
			}

			// Default Local
			return &storage.LocalDriver{
				Root:    GetEnv("STORAGE_ROOT", "./public/storage"),
				BaseURL: GetEnv("STORAGE_BASE_URL", GetEnv("APP_URL", "http://localhost:8080")+"/storage"),
			}
		})
	})
}
