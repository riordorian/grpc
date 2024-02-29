package auth

import (
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

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type LoginRequest struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	GrantType string `json:"grant_type"`
	ClientId  string `json:"client_id"`
}

func (k Keycloak) Login(login string, password string) (jwt.Token, error) {
	if "" == k.RS256 || "" == k.Secret {
		return jwt.Token{}, errors.New("keycloak: rs256 key and secret are required")
	}

	url := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", k.BaseUrl, k.Realm)

	body := urlpack.Values{}
	body.Add("grant_type", k.GrantType)
	body.Add("client_id", k.ClientId)
	body.Add("client_secret", "oxDr0qGGSObGDYnoCaHGOgJoVnjCx4A2")
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

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		secretKey := "-----BEGIN CERTIFICATE-----\n" +
			k.RS256 +
			"\n-----END CERTIFICATE-----"

		key, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(secretKey))

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return jwt.Token{}, err
	}

	return *token, nil
}

func (k Keycloak) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := k.HTTPClient.Do(req)
	if err != nil {
		return err
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
