package repositories

import (
	"database/sql"
	"github.com/backstagefood/backstagefood/internal/core/domain"
	portRepository "github.com/backstagefood/backstagefood/internal/core/ports/repositories"
)

type orderRepository struct {
	sqlClient *sql.DB
}

func NewOrderRepository(database *ApplicationDatabase) portRepository.Order {
	return &orderRepository{
		sqlClient: database.sqlClient,
	}
}

func (s *orderRepository) UpdateOrderStatus(orderId string) (*domain.Order, error) {
	query := `
		WITH updated_order AS (
			UPDATE orders
			SET status='Received', updated_at=now()
			WHERE id = $1 AND status = $2
			RETURNING id, id_customer, status, notification_attempts, notified_at, created_at, updated_at
		)
		SELECT id, id_customer, status, notification_attempts, notified_at, created_at, updated_at
		FROM updated_order;
	`
	stmt, err := s.sqlClient.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var order domain.Order
	err = stmt.QueryRow(orderId, domain.PENDING).Scan(
		&order.ID,
		&order.CustomerID,
		&order.Status,
		&order.NotificationAttempts,
		&order.NotifiedAt,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
