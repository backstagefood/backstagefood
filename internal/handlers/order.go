package handlers

import (
	"fmt"
	"net/http"

	"github.com/backstagefood/backstagefood/internal/core/domain"
	portService "github.com/backstagefood/backstagefood/internal/core/ports/services"
	"github.com/backstagefood/backstagefood/internal/repositories"

	"github.com/backstagefood/backstagefood/internal/core/services"
	"github.com/backstagefood/backstagefood/pkg/transaction"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService portService.Order
}

func NewOrderHandler(databaseConnection *repositories.ApplicationDatabase) *OrderHandler {
	productRepository := repositories.NewOrderRepository(databaseConnection)
	transactionManager := transaction.New(databaseConnection.Client())

	return &OrderHandler{
		orderService: services.NewOrderService(productRepository, transactionManager),
	}
}

// Checkout godoc
// @Summary Checkout ensure the payment is succeeded.
// @Description If payment succeeded then update order status.
// @Tags checkout
// @Produce json
// @Param orderId path string true "orderId"
// @Success 201 {object} services.CheckoutServiceDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /checkout/{orderId} [post]
func (o *OrderHandler) Checkout(c echo.Context) error {
	orderId := c.Param("orderId")
	if orderId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "order id maybe not exist"})
	}

	result, err := o.orderService.MakeCheckout(orderId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

// ListAllOrders godoc
// @Summary List all orders
// @Description Get all orders available in the database.
// @Tags orders
// @Produce json
// @Success 200 {array} domain.Order
// @Failure 500 {object} map[string]string
// @Router /orders [get]
func (o *OrderHandler) ListAllOrders(c echo.Context) error {
	result, err := o.orderService.GetOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order in the database.
// @Tags orders
// @Accept json
// @Produce json
// @Param order body services.CreateOrderDTO true "CreateOrderDTO object"
// @Success 201 {object} services.CreateOrderDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders [post]
func (o *OrderHandler) CreateOrder(c echo.Context) error {
	createOrderDTO := new(portService.CreateOrderDTO)
	fmt.Println("orderBody=", createOrderDTO)
	if err := c.Bind(createOrderDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var productsList []domain.Product
	for _, p := range createOrderDTO.Products {
		productsList = append(productsList, domain.Product{
			ID: p,
		})
	}
	newOrder := domain.Order{CustomerID: createOrderDTO.CustomerID, Products: productsList}
	order, err := o.orderService.CreateOrder(&newOrder)
	fmt.Println("order ", order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, order)

}

// DeleteOrder godoc
// @Summary Delete an order
// @Description Delete an order in the database.
// @Tags orders
// @Produce json
// @Param order path string true "orderId"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{orderId} [delete]
func (o *OrderHandler) DeleteOrder(c echo.Context) error {
	orderId := c.Param("orderId")
	if orderId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "order id maybe not exist"})
	}

	if err := o.orderService.DeleteOrder(orderId); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusNoContent, nil)
}
