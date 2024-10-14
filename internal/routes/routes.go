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
	defaultHandlers := handlers.New(s.echoEngine, databaseConnection)
	customerHandler := handlers.NewCustomerHandler(databaseConnection)

	s.routes(
		defaultHandlers,
		customerHandler,
	)

	err := s.echoEngine.Start(":" + port)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *Routes) routes(defaultHandlers *handlers.Handler, customerHandler *handlers.CustomerHandler) {
	s.echoEngine.GET("/health", defaultHandlers.Health())
	s.echoEngine.GET("/products", defaultHandlers.ListAllProducts())
	s.echoEngine.GET("/products/:id", defaultHandlers.FindProductById())

	s.echoEngine.POST("/customers/sign-up", customerHandler.CustomerSignUp)
	s.echoEngine.GET("/customers/:cpf", customerHandler.CustomerIdentify)

	s.echoEngine.GET("/swagger/*", echoSwagger.WrapHandler)
}
