package server

import (
	"net/http"

	api_activity "scouts-risk/api/activity"
	api_feedback "scouts-risk/api/feedback"
	api_hazard "scouts-risk/api/hazard"
	api_location "scouts-risk/api/location"
	"scouts-risk/web"
	web_event "scouts-risk/web/event"
	web_location "scouts-risk/web/location"

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
	e.GET("/", echo.WrapHandler(templ.Handler(web.Home())))
	e.GET("/cool", echo.WrapHandler(templ.Handler(web.Cool())))
	e.GET("/risk/location", RenderWithContext(web_location.Home))
	e.GET("/risk/location/register", echo.WrapHandler(templ.Handler(web_location.Update())))
	// An unnecessary route, just used for the prototype
	e.GET("/risk/location/new", echo.WrapHandler(templ.Handler(web_location.Create())))
	e.GET("/risk/event", RenderWithContext(web_event.Home))
	e.GET("/risk/event/register", echo.WrapHandler(templ.Handler(web_event.Update())))
	// Unnecessary routes, just used for the prototype
	e.GET("/risk/event/new", echo.WrapHandler(templ.Handler(web_event.Create())))
	e.GET("/risk/event/before", echo.WrapHandler(templ.Handler(web_event.Before())))
	e.GET("/risk/event/during", echo.WrapHandler(templ.Handler(web_event.During())))
	// POST example
	// e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

	// API/HTMX routes
	// e.POST("/api/tst", api.PostTst)
	e.POST("/api/risk/activity/new", api_activity.New)
	e.POST("/api/risk/location/new", api_location.New)
	e.POST("/api/risk/hazard/new", api_hazard.New)
	e.POST("/api/risk/consequence/new", api_hazard.NewConsequence)
	e.POST("/api/risk/control/new", api_hazard.NewControl)
	// feedback
	e.POST("/api/feedback", api_feedback.Basic)

	e.Logger.Fatal(e.Start(s.addr))

	return e
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func RenderWithContext(funcTemplate func (c echo.Context) templ.Component) func(echo.Context) error {
	return func (c echo.Context) error {
		return Render(c, http.StatusOK, funcTemplate(c))
	}
}