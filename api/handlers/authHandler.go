package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

// JwtClaims holds Token details
type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// check username & password against DB in real world use
	if username == "Ross" && password == "1234" {
		// create cookie instance
		cookie := &http.Cookie{}

		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		token, err := createJwtToken()

		if err != nil {
			log.Println("Error creating JWT token", err)
			return c.String(http.StatusInternalServerError, "Something went wrong")
		}

		// create cookie instance with token
		jwtCookie := &http.Cookie{}

		jwtCookie.Name = "JWTCookie"
		jwtCookie.Value = token
		jwtCookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You are logged in!",
			"token":   token,
		})
	}

	return c.String(http.StatusUnauthorized, "Oops, you don't have access to this!")
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"ross",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// this should be a long string in a server, never store in program
	token, err := rawToken.SignedString([]byte("mySecret"))

	if err != nil {
		return "", err
	}

	return token, nil
}
