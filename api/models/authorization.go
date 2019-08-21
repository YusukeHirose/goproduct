package models

type Auth struct {
	ClientId     string `envconfig:"QIITA_CLIANT" json:"client_id"`
	ClientSecret string `envconfig:"QIITA_SECRET" json:"client_secret"`
	Code         string `json:"code"`
}

type AuthResponse struct {
	ClientId string   `json:"client_id"`
	Scopes   []string `json:"scopes"`
	Token    string   `json:"token"`
}
