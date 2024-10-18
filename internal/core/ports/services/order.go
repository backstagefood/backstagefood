package services

import (
	"github.com/backstagefood/backstagefood/internal/core/domain"
)

// todo: must be domain
type CheckoutServiceDTO struct {
	PaymentSucceeded bool               `json:"paymentSucceeded"`
	OrderStatus      domain.OrderStatus `json:"orderStatus"`
	Order            *domain.Order      `json:"order"`
}

type Order interface {
	// todo: must be domain
	MakeCheckout(orderId string) (*CheckoutServiceDTO, error)
}
