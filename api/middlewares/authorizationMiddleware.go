package middlewares

import (
	"goproduct/api/models"
	"goproduct/db"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Authorization(group *echo.Group) {
	group.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		db := db.Connect()
		defer db.Close()
		var user models.User
		if db.Where("access_token=?", key).Find(&user).RecordNotFound() {
			return false, echo.NewHTTPError(http.StatusUnauthorized)
		}

		// トークンの期限を確かめる

		// 認証された場合にトークンの期限を更新する
		return true, nil
	}))
}