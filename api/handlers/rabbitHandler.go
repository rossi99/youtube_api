package handlers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

// Rabbit is for the name and type of rabbit
type Rabbit struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func AddRabbit(c echo.Context) error {
	rabbit := Rabbit{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&rabbit)

	if err != nil {
		log.Printf("Failed processing addRabbit request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is your Rabbit: %#v", rabbit)
	return c.String(http.StatusOK, "We got your Rabbit!")
}