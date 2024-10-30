package services

import (
	"github.com/backstagefood/backstagefood/internal/core/domain"
)

type CreateOrderDTO struct {
	CustomerID string   `json:"id_customer"`
	Products   []string `json:"products_id"`
}

type UpdateStatusDTO struct {
	Status string `json:"status"`
}

type CheckoutServiceDTO struct {
	PaymentSucceeded bool               `json:"paymentSucceeded"`
	OrderStatus      domain.OrderStatus `json:"orderStatus"`
	Order            *domain.Order      `json:"order"`
}

type Order interface {
	MakeCheckout(orderId string) (*CheckoutServiceDTO, error)
	GetOrders(status *domain.OrderStatus) ([]*domain.Order, error)
	FindOrderById(id string) (*domain.Order, error)
	CreateOrder(product *domain.Order) (map[string]string, error)
	UpdateOrder(orderId string, status domain.OrderStatus) error
	DeleteOrder(orderId string) error
}
