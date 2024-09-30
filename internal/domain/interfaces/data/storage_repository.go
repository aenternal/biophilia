package data

import (
	"io"
)

type StorageRepository interface {
	UploadFile(objectName string, file io.Reader, fileSize int64) error
	GetFile(objectName string) (io.Reader, error)
	DeleteFile(objectName string) error
}
