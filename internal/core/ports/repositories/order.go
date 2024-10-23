package repositories

import "github.com/backstagefood/backstagefood/internal/core/domain"

type Order interface {
	UpdateOrderStatus(orderId string) (*domain.Order, error)
	ListOrders() ([]*domain.Order, error)
	CreateOrder(product *domain.Order) (map[string]string, error)
}
