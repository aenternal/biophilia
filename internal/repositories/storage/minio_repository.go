package storage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
)

type MinioRepository struct {
	client *minio.Client
	bucket string
}

func NewMinioRepository(endpoint, accessKey, secretKey, bucket string, useSSL bool) (*MinioRepository, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
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
