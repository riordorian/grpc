package interceptors

import (
	"context"
	"errors"
	gp "google.golang.org/grpc"
)

type AuthInterceptor struct{}

func (a AuthInterceptor) Get() gp.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *gp.UnaryServerInfo, handler gp.UnaryHandler) (resp any, err error) {
		if a.needAuth(info.FullMethod) {
			if err := a.authorize(); err != nil {
				return nil, err
			}
		}

		return handler(ctx, req)
	}
}

func (a AuthInterceptor) needAuth(method string) bool {
	switch method {
	case "/grpc.Users/Login":
		return false
	default:
		return true
	}
}

func (a AuthInterceptor) authorize() error {
	return errors.New("authorization not implemented")
}
