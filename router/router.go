package router

import (
	"goproduct/api/handlers"

	"github.com/labstack/echo"
)

func SetUrl(e *echo.Echo) {
	products := e.Group("/products")
	middlewares.Authorization(products)
	products.GET("", handlers.GetProducts)
	products.GET("/:id", handlers.GetProduct)
	products.POST("", handlers.PostProduct)
	products.PATCH("/:id", handlers.UpdateProduct)
	products.DELETE("/:id", handlers.DeleteProduct)
	products.GET("/search", handlers.FindByName)

	auth := e.Group("/login")
	auth.GET("", handlers.GetCode)
	auth.GET("/redirect", handlers.GetQiitaAccessToken)
}
