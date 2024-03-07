package interfaces

import (
	"context"
	"github.com/golang-jwt/jwt"
)

type AuthProviderInterface interface {
	Login(ctx context.Context, login string, password string) (jwt.Token, error)
	CheckLogin(accessToken string) (bool, error)
	Can(action string, accessToken string) (bool, error)
}
