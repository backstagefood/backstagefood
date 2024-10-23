package repositories

import (
	"database/sql"
	"github.com/backstagefood/backstagefood/internal/core/domain"
	portRepository "github.com/backstagefood/backstagefood/internal/core/ports/repositories"
	"github.com/lib/pq"
)

type orderRepository struct {
	sqlClient *sql.DB
}

func NewOrderRepository(database *ApplicationDatabase) portRepository.Order {
	return &orderRepository{
		sqlClient: database.sqlClient,
	}
}

func (o *orderRepository) UpdateOrderStatus(orderId string) (*domain.Order, error) {
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
	stmt, err := o.sqlClient.Prepare(query)
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

func (o *orderRepository) ListOrders() ([]*domain.Order, error) {
	query := "SELECT id, id_customer, status, notification_attempts, notified_at, created_at, updated_at FROM orders"
	rows, err := o.sqlClient.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := make([]*domain.Order, 0)
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.Status,
			&order.NotificationAttempts,
			&order.NotifiedAt,
			&order.CreatedAt,
			&order.UpdatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

func (o *orderRepository) CreateOrder(order *domain.Order) (map[string]string, error) {
	//insertOrder := `
	//	INSERT INTO orders
	//	(id, id_customer, status, notification_attempts, notified_at, created_at, updated_at)
	//	VALUES(gen_random_uuid(), $1, $2, 0, null, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	//	RETURNING id, status, notified_at, created_at, updated_at
	//`
	query := `
		WITH InsertedOrder AS (
			INSERT INTO orders
			(id, id_customer, status, notification_attempts, notified_at, created_at, updated_at)
			VALUES (gen_random_uuid(), $1, $2, 0, null, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
			RETURNING id
		)
		INSERT INTO order_products (id_order, id_product)
		SELECT (SELECT id FROM InsertedOrder), unnest($3::uuid[])
		RETURNING id_order
	`
	stmt, err := o.sqlClient.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var listIds []string
	for _, p := range order.Products {
		listIds = append(listIds, p.ID)
	}
	err = stmt.QueryRow(&order.CustomerID, domain.PENDING, pq.Array(listIds)).Scan(
		&order.ID,
		//&order.Status,
		//&order.NotifiedAt,
		//&order.CreatedAt,
		//&order.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return map[string]string{"id": order.ID}, nil

}
