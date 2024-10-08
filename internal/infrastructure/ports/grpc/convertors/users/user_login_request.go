package users

import (
	"grpc/internal/infrastructure/adapters/auth"
	pg "grpc/pkg/proto_gen/grpc"
)

type UserLoginRequestConvertorInterface interface {
	Convert(request *pg.UserLoginRequest) (auth.LoginRequest, error)
}

type UserLoginRequestConvertor struct{}

func (r UserLoginRequestConvertor) Convert(request *pg.UserLoginRequest) (auth.LoginRequest, error) {
	return auth.LoginRequest{
		Login:    request.GetLogin(),
		Password: request.GetPassword(),
	}, nil
}
