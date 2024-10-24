package repositories

import (
	"database/sql"

	"github.com/backstagefood/backstagefood/internal/core/domain"
)

type Order interface {
	UpdateOrderStatus(tx *sql.Tx, orderId string) (*domain.Order, error)
	ListOrders() ([]*domain.Order, error)
	FindOrderById(id string) (*domain.Order, error)
	CreateOrder(product *domain.Order) (map[string]string, error)
	DeleteOrder(orderId string) error
}
