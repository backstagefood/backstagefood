package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Echo *echo.Echo
}

func New() *Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &Server{
		Echo: e,
	}
}

func (s *Server) Start(port string) {
	s.Routes()

	err := s.Echo.Start(":" + port)

	if err != nil {
		return
	}
}
