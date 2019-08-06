package models

import "time"

type Base struct {
	Id        int       `gorm:"primary_key"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"-"`
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp on update current_timestamp" json:"-"`
}

type Product struct {
	Base
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
