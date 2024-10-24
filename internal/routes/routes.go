package routes

import (
	"fmt"

	"github.com/backstagefood/backstagefood/internal/handlers"
	"github.com/backstagefood/backstagefood/internal/repositories"
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

func (s *Routes) Start(port string, databaseConnection *repositories.ApplicationDatabase) {
	defaultHandlers := handlers.New(s.echoEngine, databaseConnection)
	customerHandler := handlers.NewCustomerHandler(databaseConnection)
	productHandler := handlers.NewProductHandler(databaseConnection)
	orderHandler := handlers.NewOrderHandler(databaseConnection)

	s.routes(
		defaultHandlers,
		customerHandler,
		productHandler,
		orderHandler,
	)

	err := s.echoEngine.Start(":" + port)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *Routes) routes(
	defaultHandlers *handlers.Handler,
	customerHandler *handlers.CustomerHandler,
	productHandler *handlers.ProductHandler,
	orderHandler *handlers.OrderHandler,
) {
	s.echoEngine.GET("/health", defaultHandlers.Health)

	s.echoEngine.GET("/products", productHandler.ListAllProducts)
	s.echoEngine.GET("/products/:id", productHandler.FindProductById)
	s.echoEngine.GET("/categories", productHandler.ListAllCategories)
	s.echoEngine.POST("/products", productHandler.CreateProduct)

	s.echoEngine.POST("/customers/sign-up", customerHandler.CustomerSignUp)
	s.echoEngine.GET("/customers/:cpf", customerHandler.CustomerIdentify)

	s.echoEngine.POST("/checkout/:orderId", orderHandler.Checkout)
	s.echoEngine.GET("/orders", orderHandler.ListAllOrders)
	s.echoEngine.POST("/orders", orderHandler.CreateOrder)
	s.echoEngine.DELETE("/orders/:orderId", orderHandler.DeleteOrder)

	s.echoEngine.PUT("/products/:id", productHandler.UpdateProduct)
	s.echoEngine.DELETE("/products/:id", productHandler.DeleteProduct)

	s.echoEngine.GET("/swagger/*", echoSwagger.WrapHandler)
}
