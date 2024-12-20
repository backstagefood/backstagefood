package repositories

import (
	"database/sql"

	"github.com/backstagefood/backstagefood/internal/core/domain"
)

type Order interface {
	UpdateOrderStatus(tx *sql.Tx, orderId string) (*domain.Order, error)
	ListOrders(status *domain.OrderStatus) ([]*domain.Order, error)
	FindOrderById(id string) (*domain.Order, error)
	CreateOrder(product *domain.Order) (map[string]string, error)
	UpdateOrder(tx *sql.Tx, orderId string, status domain.OrderStatus) (int64, error)
	DeleteOrder(orderId string) error
}
