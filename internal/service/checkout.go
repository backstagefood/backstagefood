package service

import (
	"errors"
	"net/http"

	"github.com/backstagefood/backstagefood/internal/domain"
	paymentgateway "github.com/backstagefood/backstagefood/internal/service/payment_gateway"
)

var (
	error_order_pending  = errors.New("order still pending")
	error_payment_failed = errors.New("payment failed")
)

type OrderInterface interface {
	UpdateOrderStatus(orderId string) (*domain.Order, error)
}

type TransactionManagerInterface interface {
	RunWithTransaction(callback func() (interface{}, error)) (interface{}, error)
}

type CheckoutService struct {
	transactionManager TransactionManagerInterface
	orderRepository    OrderInterface
	orderId            string
}

type CheckoutServiceDTO struct {
	PaymentSucceeded bool               `json:"paymentSucceeded"`
	OrderStatus      domain.OrderStatus `json:"orderStatus"`
	Order            *domain.Order      `json:"order"`
}

func NewCheckout(
	orderRepository OrderInterface,
	orderId string,
	transactionManager TransactionManagerInterface,
) *CheckoutService {
	return &CheckoutService{
		orderRepository:    orderRepository,
		orderId:            orderId,
		transactionManager: transactionManager,
	}
}

func (c *CheckoutService) MakeCheckout() (*CheckoutServiceDTO, error) {
	transactionResult, err := c.transactionManager.RunWithTransaction(func() (interface{}, error) {
		updatedOrder, err := c.orderRepository.UpdateOrderStatus(c.orderId)
		if err != nil {
			return &CheckoutServiceDTO{
				PaymentSucceeded: true,
				OrderStatus:      domain.PENDING,
				Order:            updatedOrder,
			}, error_order_pending
		}

		// TODO: FakeCheckout() need to be interfaced when the real web hook is implemented.
		paymentGatewayResponse := paymentgateway.PaymentCheckout()
		if paymentGatewayResponse != http.StatusOK {
			return &CheckoutServiceDTO{
				PaymentSucceeded: false,
				OrderStatus:      domain.PAYMENT_FAILED,
			}, error_payment_failed
		}

		return &CheckoutServiceDTO{
			PaymentSucceeded: true,
			OrderStatus:      domain.RECEIVED,
			Order:            updatedOrder,
		}, nil
	})

	return transactionResult.(*CheckoutServiceDTO), err
}
