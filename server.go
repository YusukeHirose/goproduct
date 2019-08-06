package main

import (
	"./router"

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
	db := db.Manager()

	// Route => handler
	router.SetUrl(e)

	defer db.Close()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
