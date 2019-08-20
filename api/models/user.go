package models

import "time"

type User struct {
	Base
	QiitaId          string    `json:"id" validate:"required"`
	QiitaAccessToken string    `json:"-"`
	AccessToken      string    `json:"access_token"`
	ExpiredAt        time.Time `json:"-"`
}
