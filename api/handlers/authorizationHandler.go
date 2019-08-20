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
	return c.JSON(http.StatusOK, clientId)
}
