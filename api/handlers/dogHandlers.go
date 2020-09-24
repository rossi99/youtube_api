package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
)

// Dog is for the name and type of dog
type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func GetDogs(c echo.Context) error {
	dogName := c.QueryParam("name")
	dogType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Your dog name is: %s and its type is: %s\n", dogName, dogType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": dogName,
			"type": dogType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "You need to lets us know if you want data back",
	})
}

func AddDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &dog)

	if err != nil {
		log.Printf("Failed unmarshling in addDogs: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("This is your dog: %#v", dog)
	return c.String(http.StatusOK, "We got your dog!")
}
