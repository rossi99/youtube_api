package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func MainJWT(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)

	claim := token.Claims.(jwt.MapClaims)

	log.Println("User Name: ", claim["name"], "User ID: ", claim["jti"])

	return c.String(http.StatusOK, "You are on the top secret JWT page")
}
