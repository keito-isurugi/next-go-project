package domain

import (
	"mime/multipart"
)

type ObjectInput struct {
	Key         string
	ContentType string
	FileContent []byte
}

type StorageRepository interface {
	PutObject(file *multipart.FileHeader, bucketName, objectKey string) (string, error)
}
