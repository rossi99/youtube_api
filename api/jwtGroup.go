package api

import (
	"github.com/labstack/echo"
	"youtube_api/api/handlers"
)

func JwtGroup(g *echo.Group) {
	g.GET("/main", handlers.MainJWT)
}
