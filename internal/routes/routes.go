package routes

import (
	"fmt"

	"github.com/backstagefood/backstagefood/internal/handlers"
	db "github.com/backstagefood/backstagefood/internal/repositories"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Routes struct {
	echoEngine *echo.Echo
}

func New(echoEngine *echo.Echo) *Routes {
	return &Routes{
		echoEngine: echoEngine,
	}
}

func (s *Routes) Start(port string, databaseConnection *db.ApplicationDatabase) {
	handlers := handlers.New(s.echoEngine, databaseConnection)
	s.routes(handlers)

	err := s.echoEngine.Start(":" + port)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *Routes) routes(handlers *handlers.Handler) {
	s.echoEngine.GET("/health", handlers.Health())
	s.echoEngine.GET("/products", handlers.ListAllProducts())
	s.echoEngine.GET("/products/:id", handlers.FindProductById())

	s.echoEngine.GET("/swagger/*", echoSwagger.WrapHandler)
}
