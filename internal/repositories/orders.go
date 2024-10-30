package repositories

import (
	"database/sql"
	"errors"
	"log"

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
		&order.Customer.ID,
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

func (o *orderRepository) UpdateOrder(tx *sql.Tx, orderId string, status domain.OrderStatus) (int64, error) {
	query := `
		UPDATE orders
		SET status=$1, updated_at=now()
		WHERE id = $2	
	`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(status, orderId)

	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (o *orderRepository) ListOrders(status *domain.OrderStatus) ([]*domain.Order, error) {
	query := `
	SELECT o.id, o.id_customer, c."name" , o.status, o.notification_attempts, o.notified_at, o.created_at, o.updated_at
	 FROM orders o, customers c 
	WHERE o.id_customer = c.id
	  and o.status ILIKE '%' || $1 || '%' 
	`
	stmt, err := o.sqlClient.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Println("couldn't list orders - statement error:", err)
		return nil, err
	}

	rows, err := stmt.Query(status)
	if err != nil {
		log.Println("couldn't list orders - query error:", err)
		return nil, err
	}
	defer rows.Close()
	orders := make([]*domain.Order, 0)
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(
			&order.ID,
			&order.Customer.ID,
			&order.Customer.Name,
			&order.Status,
			&order.NotificationAttempts,
			&order.NotifiedAt,
			&order.CreatedAt,
			&order.UpdatedAt); err != nil {
			return nil, err
		}
		// find the products
		order.Products, err = o.listOrderProducts(order.ID)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

func (o *orderRepository) FindOrderById(id string) (*domain.Order, error) {
	query := `
	SELECT o.id, o.id_customer, c."name" , o.status, o.notification_attempts, o.notified_at, o.created_at, o.updated_at
	 FROM orders o, customers c 
	WHERE o.id_customer = c.id
	  AND o.id=$1
	`
	stmt, err := o.sqlClient.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Println("couldn't find order by ID - statement error:", err)
		return nil, err
	}
	var order domain.Order
	err = stmt.QueryRow(id).Scan(
		&order.ID,
		&order.Customer.ID,
		&order.Customer.Name,
		&order.Status,
		&order.NotificationAttempts,
		&order.NotifiedAt,
		&order.CreatedAt,
		&order.UpdatedAt)
	if err != nil {
		log.Println("couldn't find order by ID - query error:", err)
		return nil, err
	}
	// find the products
	order.Products, err = o.listOrderProducts(id)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *orderRepository) listOrderProducts(id string) ([]domain.Product, error) {
	query := `
		SELECT p.id, p.id_category, p.description, p.ingredients, p.price, p.created_at, p.updated_at, pc.id, pc.description 
		FROM order_products op, products p, product_categories pc 
		where op.id_product = p.id
		 and p.id_category = pc.id
		 and op.id_order = $1;
	`
	stmt, err := o.sqlClient.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Println("couldn't list order products - statement error:", err)
		return nil, err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("couldn't list order products - query error:", err)
		return nil, err
	}
	defer rows.Close()
	products := make([]domain.Product, 0)
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(
			&product.ID,
			&product.IDCategory,
			&product.Description,
			&product.Ingredients,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.ProductCategory.ID,
			&product.ProductCategory.Description); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (o *orderRepository) CreateOrder(order *domain.Order) (map[string]string, error) {
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
		log.Println("couldn't create order - statement error:", err)
		return nil, err
	}
	defer stmt.Close()
	var listIds []string
	for _, p := range order.Products {
		listIds = append(listIds, p.ID)
	}
	err = stmt.QueryRow(&order.Customer.ID, domain.PENDING, pq.Array(listIds)).Scan(
		&order.ID,
	)
	if err != nil {
		log.Println("couldn't create order - query error:", err)
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
