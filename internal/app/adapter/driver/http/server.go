package server

import (
	"net/http"

	database "github.com/backstagefood/backstagefood/internal/app/adapter/driven/postgresql"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echoEngine         *echo.Echo
	databaseConnection *database.SqlDb
}

func New(echoEngine *echo.Echo, databaseConnection *database.SqlDb) *Server {
	return &Server{
		echoEngine:         echoEngine,
		databaseConnection: databaseConnection,
	}
}

func (s *Server) Start(port string) {
	s.routes()

	err := s.echoEngine.Start(":" + port)
	if err != nil {
		return
	}
}

func (s *Server) routes() {
	s.echoEngine.GET("/health", healthy())
}

func healthy() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy")
	}
}
