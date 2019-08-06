package handlers

import (
	"net/http"

	"../models"

	"../../db"

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
