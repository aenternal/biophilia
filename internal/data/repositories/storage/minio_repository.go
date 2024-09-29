package storage

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
)

func createBucketIfNotExists(client *minio.Client, bucket string) error {
	ctx := context.Background()

	exists, err := client.BucketExists(ctx, bucket)
	if err != nil {
		return fmt.Errorf("failed to check if bucket exists: %w", err)
	}

	if !exists {
		err = client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("failed to create bucket: %w", err)
		}
	}
	return nil
}

type MinioRepository struct {
	client *minio.Client
	bucket string
}

func NewMinioRepository(endpoint, minioUser, minioPassword, bucket string, useSSL bool) (*MinioRepository, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioUser, minioPassword, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create minio client: %w", err)
	}

	err = createBucketIfNotExists(client, bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to create bucket: %w", err)
	}

	return &MinioRepository{
		client: client,
		bucket: bucket,
	}, nil
}

func (m *MinioRepository) UploadFile(objectName string, file io.Reader, fileSize int64) error {
	_, err := m.client.PutObject(context.Background(), m.bucket, objectName, file, fileSize, minio.PutObjectOptions{})
	return err
}

func (m *MinioRepository) GetFile(objectName string) (io.Reader, error) {
	return m.client.GetObject(context.Background(), m.bucket, objectName, minio.GetObjectOptions{})
}

func (m *MinioRepository) DeleteFile(objectName string) error {
	return m.client.RemoveObject(context.Background(), m.bucket, objectName, minio.RemoveObjectOptions{})
}
