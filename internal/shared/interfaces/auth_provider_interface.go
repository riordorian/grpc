package interfaces

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type AuthProviderInterface interface {
	Login(ctx context.Context, login string, password string) (jwt.Token, error)
	CheckLogin(accessToken string) (bool, error)
	Can(action string, accessToken string) (bool, error)
	GetUserIdByToken(accessToken string) (uuid.UUID, error)
}
