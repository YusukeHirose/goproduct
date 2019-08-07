package router

import (
	"goproduct/api/handlers"

	"github.com/labstack/echo"
)

func SetUrl(e *echo.Echo) {
	group := e.Group("/products")
	group.GET("", handlers.GetProducts)
	group.GET("/:id", handlers.GetProduct)
	group.POST("", handlers.PostProduct)
	group.PATCH("/:id", handlers.UpdateProduct)
}