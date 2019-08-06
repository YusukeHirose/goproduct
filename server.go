package main

import (
	"net/http"

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
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	defer db.Close()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
