package convertors

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"grpc/internal/shared/dto"
	pg "grpc/pkg/proto_gen/grpc"
	"log"
	"os"
	"time"
)

type CreateRequestConvertorInterface interface {
	Convert(request *pg.CreateRequest) (dto.CreateRequest, error)
}

type CreateRequestConvertor struct {
}

func (l CreateRequestConvertor) Convert(req *pg.CreateRequest) (dto.CreateRequest, error) {
	currentTime := time.Now().String()
	hashFunc := md5.New()
	hashFunc.Write([]byte(currentTime))
	hash := hex.EncodeToString(hashFunc.Sum(nil))
	uploadPath := "./runtime/upload/" + hash + "/"

	if err := os.Mkdir(uploadPath, os.ModePerm); err != nil {
		log.Fatal(err.Error())
	}

	defer func() {
		if err := os.RemoveAll(uploadPath); err != nil {
			log.Fatal(err.Error())
		}
	}()

	//for {
	media := req.GetMedia()
	for i := range media {
		file, err := os.Create(uploadPath + media[i].GetFileName())
		if err != nil {
			fmt.Println("Error creating file:", err)
			return dto.CreateRequest{}, err
		}
		if _, err := file.Write(media[i].GetChunk()); err != nil {
			log.Fatal(err.Error())
		}
	}
	//}

	createRequest := dto.CreateRequest{
		Title: req.GetTitle(),
		Text:  req.GetText(),
		Tags:  req.GetTags(),
	}

	return createRequest, nil
}
