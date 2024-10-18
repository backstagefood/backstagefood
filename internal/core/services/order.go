package services

import (
	"errors"
	portRepository "github.com/backstagefood/backstagefood/internal/core/ports/repositories"
	portService "github.com/backstagefood/backstagefood/internal/core/ports/services"
	"github.com/backstagefood/backstagefood/pkg/transaction"
	"net/http"

	"github.com/backstagefood/backstagefood/internal/core/domain"
	paymentgateway "github.com/backstagefood/backstagefood/internal/core/services/payment_gateway"
)

var (
	error_order_pending  = errors.New("order still pending")
	error_payment_failed = errors.New("payment failed")
)

type CheckoutService struct {
	transactionManager transaction.TransactionManagerInterface
	orderRepository    portRepository.Order
}

func NewOrderService(
	orderRepository portRepository.Order,
	transactionManager transaction.TransactionManagerInterface,
) portService.Order {
	return &CheckoutService{
		orderRepository:    orderRepository,
		transactionManager: transactionManager,
	}
}

func (c *CheckoutService) MakeCheckout(orderId string) (*portService.CheckoutServiceDTO, error) {
	transactionResult, err := c.transactionManager.RunWithTransaction(func() (interface{}, error) {
		updatedOrder, err := c.orderRepository.UpdateOrderStatus(orderId)
		if err != nil {
			return &portService.CheckoutServiceDTO{
				PaymentSucceeded: true,
				OrderStatus:      domain.PENDING,
				Order:            updatedOrder,
			}, error_order_pending
		}

		// TODO: FakeCheckout() need to be interfaced when the real web hook is implemented.
		paymentGatewayResponse := paymentgateway.PaymentCheckout()
		if paymentGatewayResponse != http.StatusOK {
			return &portService.CheckoutServiceDTO{
				PaymentSucceeded: false,
				OrderStatus:      domain.PAYMENT_FAILED,
			}, error_payment_failed
		}

		return &portService.CheckoutServiceDTO{
			PaymentSucceeded: true,
			OrderStatus:      domain.RECEIVED,
			Order:            updatedOrder,
		}, nil
	})

	return transactionResult.(*portService.CheckoutServiceDTO), err
}
