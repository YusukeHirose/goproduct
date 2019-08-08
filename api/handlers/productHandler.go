package handlers

import (
	"encoding/base64"
	"goproduct/api/models"
	"goproduct/db"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func GetProducts(c echo.Context) error {
	db := db.Connect()
	defer db.Close()
	products := []models.Product{}
	db.Find(&products)
	responseBody := map[string][]models.Product{"products": products}
	return c.JSON(http.StatusOK, responseBody)
}

func GetProduct(c echo.Context) error {
	db := db.Connect()
	defer db.Close()
	id := c.Param("id")
	product := models.Product{}
	if db.Where("id=?", id).Find(&product).RecordNotFound() {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	responseBody := map[string]models.Product{"product": product}
	return c.JSON(http.StatusOK, responseBody)
}

func PostProduct(c echo.Context) error {
	db := db.Connect()
	defer db.Close()
	product := models.Product{}
	err := c.Bind(&product)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	uploadImage(product.Image)
	//db.Create(&product)
	responseBody := map[string]models.Product{"product": product}
	return c.JSON(http.StatusOK, responseBody)
}

func UpdateProduct(c echo.Context) error {
	db := db.Connect()
	defer db.Close()
	id := c.Param("id")
	product := models.Product{}
	if db.Where("id=?", id).Find(&product).RecordNotFound() {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	db.Save(&product)
	responseBody := map[string]models.Product{"product": product}
	return c.JSON(http.StatusOK, responseBody)
}

func DeleteProduct(c echo.Context) error {
	db := db.Connect()
	defer db.Close()
	id := c.Param("id")
	product := models.Product{}
	if db.Where("id=?", id).Find(&product).RecordNotFound() {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	db.Delete(&product)
	return c.String(http.StatusNoContent, "success")
}

func FindByName(c echo.Context) error {
	db := db.Connect()
	defer db.Close()
	name := c.QueryParam("name")
	products := []models.Product{}
	if db.Where("name LIKE ?", "%"+name+"%").Find(&products).RecordNotFound() {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	responseBody := map[string][]models.Product{"products": products}
	return c.JSON(http.StatusOK, responseBody)
}

func uploadImage(image string) string {
	// base64形式のリクエストをデコードする
	if !strings.Contains(image, ",") {
		log.Println("base64の画像データではない")
	}
	imageData := strings.Split(image, ",")[1]
	imageByteData, err := base64.StdEncoding.DecodeString(imageData)
	if err != nil {
		log.Println("base64形式ではない")
	}
	imageType := http.DetectContentType(imageByteData)
	log.Println("imageType is %s", imageType)
	return imageType
}
