package services

import "github.com/backstagefood/backstagefood/internal/core/domain"

type Customer interface {
	SignUp(customer *domain.Customer) (*domain.Customer, error)
	Identify(cpf string) (*domain.Customer, error)
}
