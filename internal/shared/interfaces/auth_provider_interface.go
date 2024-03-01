package interfaces

import "github.com/golang-jwt/jwt"

type AuthProviderInterface interface {
	Login(login string, password string) (jwt.Token, error)
	CheckLogin(accessToken string) (bool, error)
}
