package handlers

import (
	"net/http"

	db "github.com/backstagefood/backstagefood/internal/repositories"
	"github.com/backstagefood/backstagefood/internal/service"
	"github.com/backstagefood/backstagefood/pkg/transaction"
	"github.com/labstack/echo/v4"
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

// Health godoc
// @Summary Health check
// @Description Check if the server and the database are up and running.
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *Handler) Health() func(c echo.Context) error {
	return func(c echo.Context) error {
		databaseStatus := "UP"
		if err := h.database.DataBaseHeatlh(); err != nil {
			databaseStatus = "DOWN"
		}
		return c.JSON(http.StatusOK, map[string]string{"status": "UP", "database": databaseStatus})
	}
}

// ListAllProducts godoc
// @Summary List all products
// @Description Get all products available in the database.
// @Tags products
// @Produce json
// @Param description query string false "Descripion"
// @Success 200 {array} domain.Product
// @Failure 500 {object} map[string]string
// @Router /products [get]
func (h *Handler) ListAllProducts() func(c echo.Context) error {
	return func(c echo.Context) error {
		description := c.QueryParam("description")
		uc := service.NewProductService(h.database)
		products, err := uc.GetProducts(description)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, products)
	}
}

// FindProductById godoc
// @Summary Find product by ID
// @Description Get a specific product by its ID.
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} domain.Product
// @Failure 500 {object} map[string]string
// @Router /products/{id} [get]
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

// Checkout godoc
// @Summary Checkout ensure the payment is succeeded.
// @Description If payment succeeded then update order status.
// @Tags checkout
// @Produce json
// @Param orderId path string true "orderId"
// @Success 201 {object} service.CheckoutServiceDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /checkout/{orderId} [post]
func (h *Handler) Checkout(c echo.Context) error {
	orderId := c.Param("orderId")
	if orderId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "order id maybe not exist"})
	}

	transactionManager := transaction.New(h.database.Client())

	service := service.NewCheckout(h.database, orderId, transactionManager)
	result, err := service.MakeCheckout()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product in the database.
// @Tags products
// @Accept json
// @Produce json
// @Param product body domain.Product true "Product object"
// @Success 201 {object} domain.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [post]
func (h *Handler) CreateProduct() func(c echo.Context) error {
	return func(c echo.Context) error {
		productDTO := new(service.ProductDTO)
		if err := c.Bind(productDTO); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		uc := service.NewProductService(h.database)
		category, err := uc.GetCategoryID(c.Param("category"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		productDTO.IDCategory = category
		createdProduct, err := uc.CreateProduct(productDTO)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, createdProduct)
	}
}
