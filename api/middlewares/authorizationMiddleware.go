package middlewares

import (
	"goproduct/api/models"
	"goproduct/db"
	"net/http"
	"time"

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
		accessedTime := time.Now()
		expiredTime := user.ExpiredAt
		if !accessedTime.Before(expiredTime) {
			return false, echo.NewHTTPError(http.StatusUnauthorized)
		}
		// 認証された場合にトークンの期限を更新する
		user.ExpiredAt = time.Now().Add(1 * time.Minute)
		db.Save(&user)
		return true, nil
	}))
}
