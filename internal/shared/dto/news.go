package dto

import (
	"github.com/google/uuid"
	"grpc/pkg/proto_gen/grpc"
)

type Status int

type ListRequest struct {
	Sort   string
	Author uuid.UUID
	Status Status
	Query  string
	Page   int32
}

type CreateRequest struct {
	Text  string
	Title string
	Tags  []string
	Media []*grpc.MediaUploadRequest
}
