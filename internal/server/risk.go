package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Constants
// const riskRegisterDbId = "bcf44593fa5245a4bb4c245b5fcdc5d0"

// Create the risk interface
// func (s *Server) Query(ctx context.Context) error {
// 	s.
// 	page, err := s.store.Page.Get(ctx, "c77982b36df84bc3a735dace06b66d71")
// 	return nil
// }

func (s *Server) RiskHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "All the risks",
	}

	return c.JSON(http.StatusOK, resp)
}

// func (s *Server) CreateEventRiskHandler(c echo.Context) error {
// 	// # 1.
// 	// s.store.

// 	client := s.store
// 	page, err := client.Page.Get(c.Request().Context(), "c77982b36df84bc3a735dace06b66d71")
// 	if err != nil {
// 		return c.JSON(500, map[string]string{"error": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, resp)
// }