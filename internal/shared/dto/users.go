package dto

type LoginRequest struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	GrantType string `json:"grant_type"`
	ClientId  string `json:"client_id"`
}
