package convertors

import (
	"github.com/golang-jwt/jwt"
	pg "grpc/pkg/proto_gen/grpc"
)

type UserLoginResponseConvertorInterface interface {
	Convert(token jwt.Token) (*pg.UserLoginResponse, error)
}

type UserLoginResponseConvertor struct{}

func (r UserLoginResponseConvertor) Convert(token jwt.Token) (*pg.UserLoginResponse, error) {
	return &pg.UserLoginResponse{Token: token.Raw}, nil
}
