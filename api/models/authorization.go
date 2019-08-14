package models

type Auth struct {
	ClientId string `envconfig:"QIITA_CLIANT"`
}
