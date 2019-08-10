package main

import (
	"goproduct/api/handlers"
	"goproduct/api/middlewares"
	"goproduct/db"
	"goproduct/router"

	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()
	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler
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
