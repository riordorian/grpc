package auth

type Keycloak struct {
	BaseUrl   string
	Realm     string
	ClientId  string
	GrantType string
}
