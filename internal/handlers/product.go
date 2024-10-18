package handlers

import (
	portService "github.com/backstagefood/backstagefood/internal/core/ports/services"
	"github.com/backstagefood/backstagefood/internal/core/services"
	"github.com/backstagefood/backstagefood/internal/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductHandler struct {
	productService portService.Product
}

func NewProductHandler(databaseConnection *repositories.ApplicationDatabase) *ProductHandler {
	productRepository := repositories.NewProductRepository(databaseConnection)

	return &ProductHandler{
		productService: services.NewProductService(productRepository),
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
func (h *ProductHandler) ListAllProducts(c echo.Context) error {
	description := c.QueryParam("description")
	products, err := h.productService.GetProducts(description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)

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
func (h *ProductHandler) FindProductById(c echo.Context) error {
	id := c.Param("id")
	product, err := h.productService.GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, product)

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
func (h *ProductHandler) CreateProduct(c echo.Context) error {
	productDTO := new(portService.ProductDTO)
	if err := c.Bind(productDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	category, err := h.productService.GetCategoryID(c.Param("category"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	productDTO.IDCategory = category
	createdProduct, err := h.productService.CreateProduct(productDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, createdProduct)

}
