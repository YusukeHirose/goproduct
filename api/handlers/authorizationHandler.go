package handlers

import (
	"goproduct/api/models"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
)

func GetCode(c echo.Context) error {
	var authStruct models.Auth

	envconfig.Process("", &authStruct)
	clientId := authStruct.ClientId
	log.Println(clientId)
	return c.JSON(http.StatusOK, clientId)
}
