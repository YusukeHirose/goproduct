package handlers

import (
	"goproduct/api/models"
	"goproduct/db"
	"net/http"

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
	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	db.Save(&product)
	responseBody := map[string]models.Product{"product": product}
	return c.JSON(http.StatusOK, responseBody)
}
