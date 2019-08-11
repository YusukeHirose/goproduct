package models

import "time"

import "gopkg.in/go-playground/validator.v9"

type Base struct {
	Id        int       `gorm:"primary_key"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"-"`
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp on update current_timestamp" json:"-"`
}

type Product struct {
	Base
	Name        string `json:"name" validate:"required,max=5"`
	Price       int    `json:"price" validate:"required,max=5"`
	Description string `json:"description" validate:"required,max=5"`
	Image       string `json:"image"`
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
