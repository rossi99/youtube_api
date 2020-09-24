package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func MainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Yay, you are on the Admin Page")
}
