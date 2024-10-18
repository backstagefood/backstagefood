package handlers

import (
	portService "github.com/backstagefood/backstagefood/internal/core/ports/services"
	"github.com/backstagefood/backstagefood/internal/repositories"
	"net/http"

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
func (h *OrderHandler) Checkout(c echo.Context) error {
	orderId := c.Param("orderId")
	if orderId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "order id maybe not exist"})
	}

	result, err := h.orderService.MakeCheckout(orderId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}
