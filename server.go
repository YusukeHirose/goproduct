package main

import (
	"goproduct/router"

	"./api/middlewares"

	"./db"

	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	middlewares.SetMainMiddlewares(e)

	// Database
	db := db.DbManager()

	// Route => handler
	router.SetUrl(e)

	defer db.Close()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
