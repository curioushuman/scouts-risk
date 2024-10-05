package simple

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"scouts-risk/simpleapi"

	"github.com/a-h/templ"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	port string
	addr string
}

/**
 * Creates a new HTTP server
 */
func NewHttp() *http.Server {
	port := httpPort()
	NewServer := &Server{
		port: port,
		addr: fmt.Sprintf(":%s", port),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         NewServer.addr,
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// API/HTMX routes
	e.GET("/", s.HelloWorldHandler)
  e.GET("/api/activity", simpleapi.New)

	e.Logger.Fatal(e.Start(s.addr))

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}
	return c.JSON(http.StatusOK, resp)
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

func httpPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return
}