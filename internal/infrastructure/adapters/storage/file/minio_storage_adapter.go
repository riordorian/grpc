package file

import (
	"github.com/minio/minio-go"
)

type MinioStorageAdapter struct {
	Client *minio.Client
}

func (f MinioStorageAdapter) Upload(fileName string, filePath string) error {
	if _, err := f.Client.FPutObject("grpc", fileName, filePath, minio.PutObjectOptions{}); err != nil {
		return err
	}

	return nil
}
