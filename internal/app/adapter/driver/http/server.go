package server

import (
	"net/http"

	database "github.com/backstagefood/backstagefood/internal/app/adapter/driven/postgresql"
	"github.com/backstagefood/backstagefood/internal/app/application/core/usecases"
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
	s.echoEngine.GET("/health", health(s.databaseConnection))
	s.echoEngine.GET("/products", ListAllProducts(s.databaseConnection))
	s.echoEngine.GET("/products/:id", findProductById(s.databaseConnection))
}

func health(d *database.SqlDb) func(c echo.Context) error {
	return func(c echo.Context) error {
		databaseStatus := "UP"
		if err := d.SqlClient.Ping(); err != nil {
			databaseStatus = "DOWN"
		}
		return c.JSON(http.StatusOK, map[string]string{"status": "UP", "database": databaseStatus})
	}
}

func ListAllProducts(db *database.SqlDb) func(c echo.Context) error {
	return func(c echo.Context) error {
		//TODO include query parameters  to filter the list
		uc := usecases.NewProductsRepository(db)
		products, err := uc.GetProducts()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, products)
	}
}

func findProductById(db *database.SqlDb) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		uc := usecases.NewProductsRepository(db)
		product, err := uc.GetProductById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, product)
	}
}
