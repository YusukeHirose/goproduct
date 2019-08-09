package handlers

// func CustomHTTPErrorHandler(err error, c echo.Context) {
// 	code := http.StatusInternalServerError
// 	if he, ok := err.(*echo.HTTPError); ok {
// 		code = he.Code
// 	}
// 	responseValue := models.Error{
// 		Status:  code,
// 		Message: "Internal Server Error has occured",
// 	}
// 	response := map[string]models.Error{"error": responseValue}
// 	c.JSON(http.StatusInternalServerError, response)
// }
