package api

import (
	"github.com/labstack/echo"
	"youtube_api/api/handlers"
)

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
