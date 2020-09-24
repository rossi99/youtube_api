package router

import (
	"youtube_api/api"
	"youtube_api/api/middleware"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// create groups
	g := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	// set all middlewares
	middleware.SetMainMiddlewares(e)
	middleware.SetAdminMiddlewares(g)
	middleware.SetCookieMiddlewares(cookieGroup)
	middleware.SetJwtMiddlewares(jwtGroup)

	// set main routes
	api.MainGroup(e)

	// set group groups
	api.AdminGroup(g)
	api.CookieGroup(cookieGroup)
	api.JwtGroup(jwtGroup)

	return e
}
