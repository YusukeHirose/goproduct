package handlers

import (
	"goproduct/api/models"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

const (
	name        string = "Name"
	price       string = "Price"
	description string = "Description"
	image       string = "Image"
	required    string = "required"
	max         string = "max"
)

func GenerateValidationErrorMessage(errors validator.ValidationErrors) map[string]models.ValidationErrors {
	var messages []string
	for _, err := range errors {
		field := err.Field()
		tag := err.Tag()
		var message string
		switch field {
		case name:
			switch tag {
			case required:
				message = "name is required"
			case max:
				message = "name is limited to 5 characters"
			}
		case price:
			switch tag {
			case required:
				message = "price is required"
			case max:
				message = "price is limited to 5 characters"
			}
		case description:
			switch tag {
			case required:
				message = "description is required"
			case max:
				message = "description is limited to 5 characters"
			}
		}
		messages = append(messages, message)
	}
	responseValue := models.ValidationErrors{
		Status:  http.StatusBadRequest,
		Message: messages,
	}
	return map[string]models.ValidationErrors{"errors": responseValue}
}
