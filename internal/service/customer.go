package service

import (
	"github.com/backstagefood/backstagefood/internal/domain"
	"github.com/backstagefood/backstagefood/internal/repositories"
)

type CustomerService struct {
	customerRepository repositories.CustomerRepository
}

func NewCustomerService(repository repositories.CustomerRepository) *CustomerService {
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
