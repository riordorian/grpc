package interceptors

import (
	"context"
	"errors"
	gp "google.golang.org/grpc"
	"grpc/internal/infrastructure/ports/grpc/proto_gen/grpc"
	"grpc/internal/shared/interfaces"
)

type AuthInterceptor struct {
	AuthProvider interfaces.AuthProviderInterface
}

func (a AuthInterceptor) Get() gp.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *gp.UnaryServerInfo, handler gp.UnaryHandler) (resp any, err error) {
		if a.needAuth(info.FullMethod) {
			if err := a.authorize(req.(*grpc.ListRequest).GetToken()); err != nil {
				return nil, err
			}

			can, err := a.can(info.FullMethod, req.(*grpc.ListRequest).GetToken())
			if !can || err != nil {
				return nil, err
			}

			return handler(ctx, req)
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

func (a AuthInterceptor) authorize(token string) error {
	if token == "" {
		return errors.New("token is empty")
	}

	if _, err := a.AuthProvider.CheckLogin(token); err != nil {
		return err
	}

	return nil
}

func (a AuthInterceptor) can(action string, token string) (bool, error) {
	return a.AuthProvider.Can(action, token)
}
