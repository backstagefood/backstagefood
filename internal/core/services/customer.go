package services

import (
	"github.com/backstagefood/backstagefood/internal/core/domain"
	portRepository "github.com/backstagefood/backstagefood/internal/core/ports/repositories"
	portService "github.com/backstagefood/backstagefood/internal/core/ports/services"
)

type CustomerService struct {
	customerRepository portRepository.Customer
}

func NewCustomerService(repository portRepository.Customer) portService.Customer {
	return &CustomerService{customerRepository: repository}
}

func (c *CustomerService) SignUp(customer *domain.Customer) (*domain.Customer, error) {
	if customer.CPF == "" {
		return nil, domain.ErrCPFIsRequired
	}

	return c.customerRepository.SignUp(customer)
}

func (c *CustomerService) Identify(cpf string) (*domain.Customer, error) {
	if cpf == "" {
		return nil, domain.ErrCPFIsRequired
	}

	return c.customerRepository.Identify(cpf)
}
