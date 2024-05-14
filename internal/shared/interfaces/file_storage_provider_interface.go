package interfaces

import "grpc/pkg/proto_gen/grpc"

type FileStorageProviderInterface interface {
	Upload(fileName string, filePath string) error
	UploadMediaFromStream(media *grpc.MediaUploadRequest) error
}
