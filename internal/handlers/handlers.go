package handlers

import (
	"net/http"

	db "github.com/backstagefood/backstagefood/internal/repositories"
	"github.com/backstagefood/backstagefood/internal/service"
	"github.com/labstack/echo"
)

type Handler struct {
	echoEngine *echo.Echo
	database   *db.ApplicationDatabase
}

func New(echoEngine *echo.Echo, databaseConnection *db.ApplicationDatabase) *Handler {
	return &Handler{
		echoEngine: echoEngine,
		database:   databaseConnection,
	}
}

// TODO: criar um DTO padrão de saída para os handlers

func (h *Handler) Health() func(c echo.Context) error {
	return func(c echo.Context) error {
		databaseStatus := "UP"
		if err := h.database.SqlClient.Ping(); err != nil {
			databaseStatus = "DOWN"
		}
		return c.JSON(http.StatusOK, map[string]string{"status": "UP", "database": databaseStatus})
	}
}

func (h *Handler) ListAllProducts() func(c echo.Context) error {
	return func(c echo.Context) error {
		//TODO include query parameters to filter the list
		uc := service.NewProductService(h.database)
		products, err := uc.GetProducts()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, products)
	}
}

func (h *Handler) FindProductById() func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		uc := service.NewProductService(h.database)
		product, err := uc.GetProductById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, product)
	}
}
