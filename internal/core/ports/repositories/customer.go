package repositories

import "github.com/backstagefood/backstagefood/internal/core/domain"

type Customer interface {
	SignUp(*domain.Customer) (*domain.Customer, error)
	Identify(string) (*domain.Customer, error)
}
