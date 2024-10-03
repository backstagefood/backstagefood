package http

import "github.com/labstack/echo/v4"

func (s *Server) Routes() {
	s.Echo.GET("/health", func(c echo.Context) error {
		return c.String(200, "healthy")
	})
}
