package users

import (
	"context"
	"fmt"
	"grpc/internal/domain/news"
	"grpc/internal/infrastructure/adapters/auth"
	"grpc/internal/shared/interfaces"
)

type LoginHandler struct {
	AuthProvider interfaces.AuthProviderInterface
}

type LoginHandlerInterface interface {
	Handle(ctx context.Context, req auth.LoginRequest) ([]news.New, error)
}

func NewLoginHandler(authProvider interfaces.AuthProviderInterface) LoginHandlerInterface {
	return LoginHandler{
		AuthProvider: authProvider,
	}
}

func (l LoginHandler) Handle(ctx context.Context, req auth.LoginRequest) ([]news.New, error) {
	login, err := l.AuthProvider.Login(req.Login, req.Password)
	if err != nil {
		return nil, err
	}
	fmt.Println(login)

	return nil, nil
}
