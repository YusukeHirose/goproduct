package handlers

import (
	"bytes"
	"encoding/json"
	"goproduct/api/models"
	"io/ioutil"
	"log"
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

func GetAccessToken(c echo.Context) error {
	var requestBody models.Auth
	envconfig.Process("", &requestBody)
	requestBody.Code = c.QueryParam("code")
	requestByte, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}
	requestBuffer := bytes.NewBuffer(requestByte)
	requestUrl := "https://qiita.com/api/v2/access_tokens"
	res, err := http.Post(requestUrl, "application/json", requestBuffer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "リクエスト失敗")
	}
	if res.StatusCode != http.StatusCreated {
		return c.JSON(http.StatusInternalServerError, "リクエスト失敗")
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	return c.JSON(http.StatusOK, bodyString)
}
