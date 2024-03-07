package users

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/mock"
	"grpc/internal/infrastructure/adapters/auth"
	"grpc/internal/shared/interfaces"
	"grpc/pkg/mocks"
	"reflect"
	"testing"
)

func TestLoginHandler_Handle(t *testing.T) {
	type fields struct {
		AuthProvider interfaces.AuthProviderInterface
	}

	authProvider := mocks.NewAuthProviderInterface(t)
	ctx := context.Background()
	authProvider.EXPECT().Login(ctx, "test", "test").RunAndReturn(func(ctx context.Context, login string, password string) (jwt.Token, error) {
		return *jwt.New(jwt.SigningMethodHS256), nil
	})
	authProvider.EXPECT().Login(ctx, "", mock.AnythingOfType("string")).RunAndReturn(func(ctx context.Context, login string, password string) (jwt.Token, error) {
		return jwt.Token{}, errors.New("invalid credentials")
	})
	authProvider.EXPECT().Login(ctx, mock.AnythingOfType("string"), "").RunAndReturn(func(ctx context.Context, login string, password string) (jwt.Token, error) {
		return jwt.Token{}, errors.New("invalid credentials")
	})

	type args struct {
		ctx context.Context
		req auth.LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    jwt.Token
		wantErr bool
	}{
		{
			name: "Correct request",
			fields: fields{
				AuthProvider: authProvider,
			},
			args: args{
				ctx: context.Background(),
				req: auth.LoginRequest{
					Login:    "test",
					Password: "test",
				},
			},
			want:    *jwt.New(jwt.SigningMethodHS256),
			wantErr: false,
		},
		{
			name: "Invalid login request",
			fields: fields{
				AuthProvider: authProvider,
			},
			args: args{
				ctx: context.Background(),
				req: auth.LoginRequest{
					Login:    "",
					Password: "test",
				},
			},
			want:    jwt.Token{},
			wantErr: true,
		},
		{
			name: "Invalid password request",
			fields: fields{
				AuthProvider: authProvider,
			},
			args: args{
				ctx: context.Background(),
				req: auth.LoginRequest{
					Login:    "test",
					Password: "",
				},
			},
			want:    jwt.Token{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LoginHandler{
				AuthProvider: tt.fields.AuthProvider,
			}
			got, err := l.Handle(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}
