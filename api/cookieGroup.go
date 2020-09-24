package api

import (
	"github.com/labstack/echo"
	"youtube_api/api/handlers"
)

func CookieGroup(g *echo.Group)  {
	g.GET("/main", handlers.MainCookie)
}
