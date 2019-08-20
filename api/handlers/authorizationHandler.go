package handlers

import (
	"bytes"
	"encoding/json"
	"goproduct/api/models"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
)

const baseApiUrl = "https://qiita.com/api/v2"

func GetCode(c echo.Context) error {
	var authStruct models.Auth

	envconfig.Process("", &authStruct)
	clientId := authStruct.ClientId
	requestUrl := baseApiUrl + "/oauth/authorize?client_id=" + clientId + "&scope=read_qiita"
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
	requestUrl := baseApiUrl + "/access_tokens"
	res, err := http.Post(requestUrl, "application/json", requestBuffer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "リクエスト失敗")
	}
	if res.StatusCode != http.StatusCreated {
		return c.JSON(http.StatusInternalServerError, "リクエスト失敗")
	}
	var authResponse models.AuthResponse
	if err := json.NewDecoder(res.Body).Decode(&authResponse); err != nil {
		log.Fatal(err)
	}
	name := getUserQiitaId(authResponse.Token)
	return c.JSON(http.StatusOK, name)
}

func getUserQiitaId(qiitaAccessToken string) string {
	req, err := http.NewRequest("GET", baseApiUrl+"/authenticated_user", nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	req.Header.Add("Authorization", "Bearer "+qiitaAccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var user models.User
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		log.Fatal(err)
	}
	return user.QiitaId
}
