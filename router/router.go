package router

import (
	"../api/handlers"

	"github.com/labstack/echo"
)

func SetUrl(e *echo.Echo) {
	group := e.Group("/products")
	group.GET("", handlers.GetProducts)
}
