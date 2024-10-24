package repositories

import (
	"database/sql"
	"errors"

	"github.com/backstagefood/backstagefood/internal/core/domain"
	portRepository "github.com/backstagefood/backstagefood/internal/core/ports/repositories"
	"github.com/lib/pq"
)

var (
	errorMoreRowsWereDeletedThenExpected = errors.New("more rows were deleted than expected")
)

type orderRepository struct {
	sqlClient *sql.DB
}

func NewOrderRepository(database *ApplicationDatabase) portRepository.Order {
	return &orderRepository{
		sqlClient: database.sqlClient,
	}
}

func (o *orderRepository) UpdateOrderStatus(tx *sql.Tx, orderId string) (*domain.Order, error) {
	query := `
		WITH updated_order AS (
			UPDATE orders
			SET status='RECEIVED', updated_at=now()
			WHERE id = $1 AND status = $2
			RETURNING id, id_customer, status, notification_attempts, notified_at, created_at, updated_at
		)
		SELECT id, id_customer, status, notification_attempts, notified_at, created_at, updated_at
		FROM updated_order;
	`
	stmt, err := tx.Prepare(query)
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

func (o *orderRepository) DeleteOrder(orderId string) error {
	query := `
		WITH deleted_order_products AS (
			DELETE FROM order_products
			WHERE id_order = $1
			RETURNING id_order
		)
		DELETE FROM orders
		WHERE id IN (SELECT id_order FROM deleted_order_products)
	`

	stmt, err := o.sqlClient.Prepare(query)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(orderId)
	if err != nil {
		return err
	}

	oneLineExpected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if oneLineExpected != 1 {
		return errorMoreRowsWereDeletedThenExpected
	}

	return nil
}
