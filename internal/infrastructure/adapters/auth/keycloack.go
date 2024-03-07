package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	urlpack "net/url"
	"strings"
)

type Keycloak struct {
	BaseUrl    string
	Realm      string
	ClientId   string
	GrantType  string
	HTTPClient *http.Client
	RS256      string
	Secret     string
}

type Claims struct {
	jwt.StandardClaims
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type LoginRequest struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	GrantType string `json:"grant_type"`
	ClientId  string `json:"client_id"`
}

func (k Keycloak) Login(ctx context.Context, login string, password string) (jwt.Token, error) {
	if "" == k.RS256 || "" == k.Secret {
		return jwt.Token{}, errors.New("keycloak: rs256 key and secret are required")
	}

	url := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", k.BaseUrl, k.Realm)

	body := urlpack.Values{}
	body.Add("grant_type", k.GrantType)
	body.Add("client_id", k.ClientId)
	body.Add("client_secret", k.Secret)
	body.Add("username", login)
	body.Add("password", password)

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body.Encode()))
	if err != nil {
		return jwt.Token{}, err
	}
	var response map[string]interface{}
	err = k.sendRequest(req, &response)
	if err != nil {
		return jwt.Token{}, err
	}

	accessToken := response["access_token"].(string)

	token, _, err := k.parseToken(accessToken)

	if err != nil {
		return jwt.Token{}, err
	}

	return token, nil
}

func (k Keycloak) parseToken(accessToken string) (jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		secretKey := "-----BEGIN CERTIFICATE-----\n" +
			k.RS256 +
			"\n-----END CERTIFICATE-----"

		key, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(secretKey))

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	return *token, claims, err
}

func (k Keycloak) CheckLogin(token string) (bool, error) {
	accessToken, _, err := k.parseToken(token)
	if err != nil {
		return false, err
	}

	return accessToken.Valid, nil
}

func (k Keycloak) Can(action string, token string) (bool, error) {
	accessToken, claims, err := k.parseToken(token)
	if !accessToken.Valid || err != nil {
		return false, errors.New("jwt is not valid")
	}

	resourcesAccess := claims["resource_access"].(map[string]interface{})
	grpcAccess := resourcesAccess["grpc"].(map[string]interface{})
	grpcRoles := grpcAccess["roles"].([]interface{})

	for _, role := range grpcRoles {
		if role.(string) == action {
			return true, nil
		}
	}

	return false, errors.New("action is not allowed")

}

func (k Keycloak) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := k.HTTPClient.Do(req)
	if err != nil {
		return errors.New("error when sending request to keycloak." + err.Error())
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return err
	}

	return nil
}
