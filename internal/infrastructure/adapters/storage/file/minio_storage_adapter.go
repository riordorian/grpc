package file

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/minio/minio-go"
	"grpc/pkg/proto_gen/grpc"
	"log"
	"os"
	"time"
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

func (f MinioStorageAdapter) UploadMediaFromStream(media *grpc.MediaUploadRequest) error {
	currentTime := time.Now().String()
	hashFunc := md5.New()
	hashFunc.Write([]byte(currentTime))
	hash := hex.EncodeToString(hashFunc.Sum(nil))
	uploadPath := "./runtime/upload/" + hash + "/"

	if err := os.Mkdir(uploadPath, os.ModePerm); err != nil {
		return err
	}

	defer func() {
		if err := os.RemoveAll(uploadPath); err != nil {
			log.Fatal(err.Error())
		}
	}()

	localFileName := uploadPath + media.GetFileName()
	file, err := os.Create(localFileName)
	if err != nil {
		return err
	}

	if _, err := file.Write(media.GetChunk()); err != nil {
		log.Fatal(err.Error())
	}

	err = f.Upload(media.GetFileName(), localFileName)
	if err != nil {
		return err
	}

	return nil
}
