package main

import (
	"goproduct/api/handlers"
	"goproduct/api/middlewares"
	"goproduct/api/models"

	"goproduct/db"
	"goproduct/router"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	// Echo instance
	e := echo.New()
	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler
	e.Validator = &models.CustomValidator{Validator: validator.New()}
	// Middleware
	middlewares.SetMainMiddlewares(e)

	// Database
	db := db.Manager()

	// Route => handler
	router.SetUrl(e)

	defer db.Close()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
