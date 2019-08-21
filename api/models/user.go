package models

import "time"

type User struct {
	Base
	QiitaId          string    `json:"id" validate:"required"`
	QiitaAccessToken string    `json:"qiita_access_token"`
	AccessToken      string    `json:"access_token"`
	ExpiredAt        time.Time `json:"-"`
}
