package service

import (
	"fmt"

	"github.com/backstagefood/backstagefood/internal/domain"
	"github.com/backstagefood/backstagefood/internal/service/checkout"
)

type OrderInterface interface {
	UpdateOrderStatus(orderId string) (*domain.Order, error)
}

type CheckoutService struct {
	orderRepository OrderInterface
	orderId         string
}

type CheckoutServiceDTO struct {
	PaymentSucceeded bool               `json:"paymentSucceeded"`
	OrderStatus      domain.OrderStatus `json:"orderStatus"`
	Order            *domain.Order      `json:"order"`
}

func NewCheckout(orderRepository OrderInterface, orderId string) *CheckoutService {
	return &CheckoutService{
		orderRepository: orderRepository,
		orderId:         orderId,
	}
}

func (c *CheckoutService) MakeCheckout() (*CheckoutServiceDTO, error) {
	// TODO: FakeCheckout() need to be interfaced when the real web hook is implemented.
	if !checkout.FakeCheckout() {
		return &CheckoutServiceDTO{
			PaymentSucceeded: false,
			OrderStatus:      domain.PAYMENT_FAILED,
		}, fmt.Errorf("payment failed")
	}

	// TODO: Checkout OK but for some reason DB task failed, what to do?
	// retry? requeue? reimbursement? (╥‸╥)
	updatedOrder, err := c.orderRepository.UpdateOrderStatus(c.orderId)
	if err != nil {
		return &CheckoutServiceDTO{
			PaymentSucceeded: true,
			OrderStatus:      domain.PENDING,
			Order:            updatedOrder,
		}, fmt.Errorf("order still pending")
	}

	return &CheckoutServiceDTO{
		PaymentSucceeded: true,
		OrderStatus:      domain.RECEIVED,
		Order:            updatedOrder,
	}, nil
}
