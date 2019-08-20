package handlers

import (
	"goproduct/api/models"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
)

func GetCode(c echo.Context) error {
	var authStruct models.Auth

	envconfig.Process("", &authStruct)
	clientId := authStruct.ClientId
	requestUrl := "https://qiita.com/api/v2/oauth/authorize?client_id=" + clientId
	return c.Redirect(http.StatusMovedPermanently, requestUrl)
}
