package services

import (
	"github.com/backstagefood/backstagefood/internal/core/domain"
)

type CreateOrderDTO struct {
	CustomerID string   `json:"id_customer"`
	Products   []string `json:"products_id"`
}

type CheckoutServiceDTO struct {
	PaymentSucceeded bool               `json:"paymentSucceeded"`
	OrderStatus      domain.OrderStatus `json:"orderStatus"`
	Order            *domain.Order      `json:"order"`
}

type Order interface {
	MakeCheckout(orderId string) (*CheckoutServiceDTO, error)
	GetOrders() ([]*domain.Order, error)
	FindOrderById(id string) (*domain.Order, error)
	CreateOrder(product *domain.Order) (map[string]string, error)
	DeleteOrder(orderId string) error
}
