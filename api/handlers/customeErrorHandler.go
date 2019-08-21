package handlers

import (
	"goproduct/api/models"
	"net/http"

	"github.com/labstack/echo"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var code int
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	var responseValue models.Error
	switch code {
	case http.StatusInternalServerError:
		responseValue = models.Error{
			Status:  code,
			Message: "Internal server error has occured",
		}
		response := map[string]models.Error{"error": responseValue}
		c.JSON(http.StatusInternalServerError, response)
	case http.StatusBadRequest:
		responseValue = models.Error{
			Status:  code,
			Message: "Bad request",
		}
		response := map[string]models.Error{"error": responseValue}
		c.JSON(http.StatusBadRequest, response)
	case http.StatusMethodNotAllowed:
		responseValue = models.Error{
			Status:  code,
			Message: "Method not arrowed",
		}
		response := map[string]models.Error{"error": responseValue}
		c.JSON(http.StatusMethodNotAllowed, response)
	case http.StatusNotFound:
		responseValue = models.Error{
			Status:  code,
			Message: "resource is not found",
		}
		response := map[string]models.Error{"error": responseValue}
		c.JSON(http.StatusNotFound, response)
	case http.StatusUnauthorized:
		responseValue = models.Error{
			Status:  http.StatusUnauthorized,
			Message: "Authorization faild",
		}
		response := map[string]models.Error{"error": responseValue}
		c.JSON(http.StatusUnauthorized, response)
	default:
		responseValue = models.Error{
			Status:  code,
			Message: "Unexpected error",
		}
		response := map[string]models.Error{"error": responseValue}
		c.JSON(http.StatusInternalServerError, response)
	}
}
