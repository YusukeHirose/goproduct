package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goproduct/api/models"
	"goproduct/db"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
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

func GetQiitaAccessToken(c echo.Context) error {
	var requestBody models.Auth
	envconfig.Process("", &requestBody)
	requestBody.Code = c.QueryParam("code")
	requestByte, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	requestBuffer := bytes.NewBuffer(requestByte)
	req, err := http.NewRequest("POST", baseApiUrl+"/access_tokens", requestBuffer)
	if err != nil {
		log.Fatal(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	req.Header.Add("Content-type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		log.Fatal(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// レスポンスを受け取る
	var authResponse models.AuthResponse
	if err := json.NewDecoder(res.Body).Decode(&authResponse); err != nil {
		log.Fatal(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// ユーザーデータの作成
	db := db.Connect()
	defer db.Close()
	var user models.User
	user.QiitaId, _ = getUserQiitaId(user.QiitaAccessToken)
	if db.Where("qiita_id=?", user.QiitaId).Find(&user).RecordNotFound() {
		user.QiitaAccessToken = authResponse.Token
		user.AccessToken, user.ExpiredAt = generateAccessToken()
		db.Create(&user)
	} else {
		// ユーザーが存在していた場合更新処理をする
		user.QiitaAccessToken = authResponse.Token
		user.AccessToken, user.ExpiredAt = generateAccessToken()
		db.Save(&user)
	}
	return c.JSON(http.StatusOK, user)
}

func getUserQiitaId(qiitaAccessToken string) (string, error) {
	req, err := http.NewRequest("GET", baseApiUrl+"/authenticated_user", nil)
	if err != nil {
		log.Fatal(err)
		return "", echo.NewHTTPError(http.StatusInternalServerError)
	}
	client := &http.Client{}
	req.Header.Add("Authorization", "Bearer "+qiitaAccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return "", echo.NewHTTPError(http.StatusInternalServerError)
	}
	var user models.User
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		log.Fatal(err)
		return "", echo.NewHTTPError(http.StatusInternalServerError)
	}
	return user.QiitaId, nil
}

func generateAccessToken() (string, time.Time) {
	accessToken, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	//expiredAt := time.Now().AddDate(0,0,1)
	expiredAt := time.Now().Add(1 * time.Minute)
	log.Println(fmt.Sprint(expiredAt))
	return accessToken.String(), expiredAt
}
