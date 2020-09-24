package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Hey, form the server side!")
}
