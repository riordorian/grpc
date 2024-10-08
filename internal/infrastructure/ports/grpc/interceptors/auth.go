package interceptors

import (
	"context"
	"errors"
	gp "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc/internal/shared/interfaces"
	"strings"
)

type AuthInterceptor struct {
	AuthProvider interfaces.AuthProviderInterface
}

func (a AuthInterceptor) Get() gp.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *gp.UnaryServerInfo, handler gp.UnaryHandler) (resp any, err error) {

		if a.needAuth(info.FullMethod) {
			bearerToken, err := a.GetTokenFromContext(ctx)
			if err != nil {
				return nil, err
			}
			if err := a.authorize(bearerToken); err != nil {
				return nil, err
			}

			can, err := a.can(info.FullMethod, bearerToken)
			if !can || err != nil {
				return nil, err
			}

			return handler(ctx, req)
		}

		return handler(ctx, req)
	}
}

func (a AuthInterceptor) GetStream() gp.StreamServerInterceptor {
	return func(
		srv any,
		ss gp.ServerStream,
		info *gp.StreamServerInfo,
		handler gp.StreamHandler,
	) error {
		//b := ss.RecvMsg(grpc.CreateRequest.ProtoMessage)
		if a.needAuth(info.FullMethod) {
			ctx := ss.Context()
			token, err := a.GetTokenFromContext(ctx)
			if err != nil {
				return err
			}

			if err := a.authorize(token); err != nil {
				return err
			}

			can, err := a.can(info.FullMethod, token)
			if !can || err != nil {
				return err
			}

			userId, err := a.AuthProvider.GetUserIdByToken(token)
			if err != nil {
				return err
			}

			err = handler(srv, &serverStream{ServerStream: ss, newCtx: context.WithValue(ctx, "userId", userId)})
			return err
		}

		return handler(srv, ss)
	}
}

// monitoredStream wraps grpc.ServerStream allowing each Sent/Recv of message to report.
type serverStream struct {
	gp.ServerStream

	newCtx context.Context
}

func (s *serverStream) Context() context.Context {
	return s.newCtx
}

func (s *serverStream) SendMsg(m any) error {
	err := s.ServerStream.SendMsg(m)
	return err
}

func (s *serverStream) RecvMsg(m any) error {
	err := s.ServerStream.RecvMsg(m)
	return err
}

func (a AuthInterceptor) needAuth(method string) bool {
	switch method {
	case "/grpc.Users/Login",
		"/grpc.Search/SearchQuery",
		"/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo":
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

func (a AuthInterceptor) GetTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata")
	}
	bearer := md["authorization"]
	bearerToken := bearer[0]
	bearerToken = strings.TrimPrefix(bearerToken, "Bearer ")

	return bearerToken, nil
}
