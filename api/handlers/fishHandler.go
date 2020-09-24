package handlers

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

// Fish is for the name and type of fish
type Fish struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func AddFish(c echo.Context) error {
	fish := Fish{}

	err := c.Bind(&fish)

	if err != nil {
		log.Printf("Failed processing addFish request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is your Rabbit: %#v", fish)
	return c.String(http.StatusOK, "We got your Rabbit!")
}
