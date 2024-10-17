package repositories

import (
	"database/sql"
	"github.com/backstagefood/backstagefood/internal/domain"
)

type CustomerRepository interface {
	SignUp(*domain.Customer) (*domain.Customer, error)
	Identify(string) (*domain.Customer, error)
}

type customerRepository struct {
	sqlClient *sql.DB
}

func NewCustomerRepository(database *ApplicationDatabase) CustomerRepository {
	return &customerRepository{
		sqlClient: database.sqlClient,
	}
}

func (r *customerRepository) SignUp(customer *domain.Customer) (*domain.Customer, error) {
	query := "INSERT INTO customers (id, name, cpf, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	stmt, err := r.sqlClient.Prepare(query)

	defer stmt.Close()

	if err != nil {
		return nil, err
	}

	if err = stmt.QueryRow(
		customer.ID,
		customer.Name,
		customer.CPF,
		customer.Email,
		customer.CreatedAt,
		customer.UpdatedAt,
	).Scan(&customer.ID); err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *customerRepository) Identify(cpf string) (*domain.Customer, error) {
	query := "SELECT id, name, cpf, email, created_at, updated_at FROM customers WHERE cpf = $1"
	stmt, err := r.sqlClient.Prepare(query)

	defer stmt.Close()

	if err != nil {
		return nil, err
	}

	customer := new(domain.Customer)

	if err = stmt.QueryRow(
		cpf,
	).Scan(
		&customer.ID,
		&customer.Name,
		&customer.CPF,
		&customer.Email,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return customer, nil
}
