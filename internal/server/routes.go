package server

import (
	"net/http"

	"scouts-risk/api"
	"scouts-risk/web"

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
	e.GET("/", echo.WrapHandler(templ.Handler(web.Root())))
	e.GET("/risks/event", echo.WrapHandler(templ.Handler(web.EventRegister())))
	// POST example
	// e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

	// HTMX routes
	e.POST("/api/tst", api.PostTst)

	// api routes
	e.GET("/api/risk", s.RiskHandler)
	// e.GET("/api/risk/register", s.RiskHandler)

	return e
}
