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
	url := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", k.BaseUrl, k.Realm)
	/*loginReq := LoginRequest{
		Login:     login,
		Password:  password,
		GrantType: k.GrantType,
		ClientId:  k.ClientId,
	}*/

	body := urlpack.Values{}
	body.Add("grant_type", k.GrantType)
	body.Add("client_id", k.ClientId)
	body.Add("username", login)
	body.Add("password", password)

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body.Encode()))
	if err != nil {
		return jwt.Token{}, err
	}
	az := k.sendRequest(req, &jwt.Token{})
	fmt.Println(az)
	return jwt.Token{}, nil
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

	fullResponse := successResponse{
		Data: v,
	}
	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}

	return nil
}
