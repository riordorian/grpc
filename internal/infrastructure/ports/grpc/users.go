package grpc

import (
	"context"
	"grpc/pkg/proto_gen/grpc"
)

func (NewsServer) mustEmbedUnimplementedUsersServer() {}
func (s NewsServer) Login(ctx context.Context, req *grpc.UserLoginRequest) (*grpc.UserLoginResponse, error) {
	loginRequest, err := s.Convertors.LoginRequest.Convert(req)
	if err != nil {
		return nil, err
	}

	token, err := s.Handlers.Queries.Login.Handle(ctx, loginRequest)
	if err != nil {
		return nil, err
	}
	rawToken, err := s.Convertors.LoginResponse.Convert(token)
	if err != nil {
		return nil, err

	}

	return rawToken, nil
}
