package api

import (
	"youtube_api/api/handlers"
	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/login", handlers.Login)
	e.GET("/greeting", handlers.Greeting)
	e.GET("/dogs/:data", handlers.GetDogs)

	// Post Endpoints
	e.POST("/dogs", handlers.AddDog)
	e.POST("/rabbit", handlers.AddRabbit)
	e.POST("/fish", handlers.AddFish)
}
