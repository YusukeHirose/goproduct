package router

import (
	"net/http"

	"github.com/labstack/echo"
)

func SetUrl(e *echo.Echo) {
	group := e.Group("/products")
	group.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
}
