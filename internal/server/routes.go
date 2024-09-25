package server

import (
	"net/http"

	api_location "scouts-risk/api/location"
	"scouts-risk/web"
	web_event "scouts-risk/web/event"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// static assets
	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))

	// web routes
	// e.GET("/", echo.WrapHandler(templ.Handler(web.Root())))
	e.GET("/risks/event", echo.WrapHandler(templ.Handler(web_event.Register())))
	// POST example
	// e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

	// HTMX routes
	// e.POST("/api/tst", api.PostTst)
	e.POST("/api/location/new", api_location.New)

	// api routes
	e.GET("/api/risk", s.RiskHandler)
	// e.GET("/api/risk/register", s.RiskHandler)

	return e
}
