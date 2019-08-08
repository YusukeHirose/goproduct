package handlers

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"goproduct/api/models"
	"goproduct/db"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

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
	imagePath := uploadImage(product.Image)
	imageFileName := strings.TrimLeft(imagePath, imagesDir)
	product.Image = imageFileName
	db.Create(&product)
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
	deleteImageFile := product.Image
	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	imagePath := uploadImage(product.Image)
	imageFileName := strings.TrimLeft(imagePath, imagesDir)
	product.Image = imageFileName
	db.Save(&product)
	// 更新された画像ファイルは削除
	deleteUploadedImageFile(imagesDir + deleteImageFile)
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

const (
	Png       = "image/png"
	Jpg       = "image/jpg"
	Jpeg      = "image/jpeg"
	imagesDir = "static/images/"
)

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
	filePath := generateFilePath(imageType)

	writeImageData(filePath, imageByteData)

	return filePath
}

func generateFilePath(imageType string) string {
	var extention string
	switch imageType {
	case Png:
		extention = ".png"
	case Jpg:
		extention = ".jpg"
	case Jpeg:
		extention = ".jpeg"
	default:
		log.Println("画像データではない")
	}
	return imagesDir + fmt.Sprint(time.Now().Format("20060102150405")) + extention
}

func writeImageData(filePath string, imageByteData []byte) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.Write(imageByteData)
	if err != nil {
		log.Println("writing is faild")
		log.Fatal(err)
	}
	writer.Flush()
}

func deleteUploadedImageFile(imagePath string) {
	// ファイルの存在確認
	_, err := os.Stat(imagePath)
	if err != nil {
		log.Fatal(err)
		log.Println("削除するファイルが存在しない")
		return
	}
	if err := os.Remove(imagePath); err != nil {
		log.Fatal(err)
	}
}
