package repository

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"

	domain "github.com/keito-isurugi/next-go-project/internal/domain/storage"
	"github.com/keito-isurugi/next-go-project/internal/infra/env"
)

type s3Repository struct {
	ev       *env.Values
	S3Client s3iface.S3API
}

func NewS3Repository(ev *env.Values, s3 s3iface.S3API) domain.StorageRepository {
	return &s3Repository{
		ev:       ev,
		S3Client: s3,
	}
}

func (s3r *s3Repository) PutObject(file *multipart.FileHeader, bucketName, objectKey string) (string, error) {
	// ファイルを開く
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	// Content-Typeを決定
	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		// Content-Typeが指定されていない場合は、ファイルの内容から推測
		buffer := make([]byte, 512)
		_, err = src.Read(buffer)
		if err != nil {
			return "", fmt.Errorf("failed to read file content: %v", err)
		}
		contentType = http.DetectContentType(buffer)

		// ファイルポインタを先頭に戻す
		_, err = src.Seek(0, 0)
		if err != nil {
			return "", fmt.Errorf("failed to reset file pointer: %v", err)
		}
	}

	// PutObjectInput構造体を作成
	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectKey),
		Body:        src,
		ContentType: aws.String(contentType),
	}

	// S3にアップロード
	_, err = s3r.S3Client.PutObject(input)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	// アップロードされたオブジェクトのURLを生成
	url := fmt.Sprintf("http://localhost:4566/%s/%s", bucketName, objectKey)

	return url, nil
}
