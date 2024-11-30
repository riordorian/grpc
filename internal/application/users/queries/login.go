package users

import (
	"context"
	"github.com/golang-jwt/jwt"
	"grpc/internal/shared/dto"
	"grpc/internal/shared/interfaces"
)

type LoginHandler struct {
	AuthProvider interfaces.AuthProviderInterface
}

type LoginHandlerInterface interface {
	Handle(ctx context.Context, req dto.LoginRequest) (jwt.Token, error)
}

func NewLoginHandler(authProvider interfaces.AuthProviderInterface) LoginHandlerInterface {
	return LoginHandler{
		AuthProvider: authProvider,
	}
}

func (l LoginHandler) Handle(ctx context.Context, req dto.LoginRequest) (jwt.Token, error) {
	// Todo: Add context
	token, err := l.AuthProvider.Login(ctx, req.Login, req.Password)
	if err != nil {
		return jwt.Token{}, err
	}

	return token, nil
}
