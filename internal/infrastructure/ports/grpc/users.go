package grpc

import (
	"context"
	"grpc/internal/shared/dto"
	"grpc/pkg/proto_gen/grpc"
	pg "grpc/pkg/proto_gen/grpc"
)

func (NewsServer) mustEmbedUnimplementedUsersServer() {}

func (s NewsServer) Login(ctx context.Context, req *grpc.UserLoginRequest) (*grpc.UserLoginResponse, error) {
	requestDto := new(dto.LoginRequest)
	_, err := s.Convertors.Request.Convert(req, requestDto)
	if err != nil {
		return nil, err
	}

	token, err := s.Handlers.Queries.Login.Handle(ctx, *requestDto)
	if err != nil {
		return nil, err
	}

	return &pg.UserLoginResponse{Token: token.Raw}, nil
}
