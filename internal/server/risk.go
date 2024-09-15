package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) RiskHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "All the risks",
	}

	return c.JSON(http.StatusOK, resp)
}