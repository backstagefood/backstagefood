package services

import (
	"github.com/backstagefood/backstagefood/internal/core/domain"
	portRepository "github.com/backstagefood/backstagefood/internal/core/ports/repositories"
	portService "github.com/backstagefood/backstagefood/internal/core/ports/services"
	"github.com/backstagefood/backstagefood/pkg/cpf"
)

type CustomerService struct {
	customerRepository portRepository.Customer
}

func NewCustomerService(repository portRepository.Customer) portService.Customer {
	return &CustomerService{customerRepository: repository}
}

func (c *CustomerService) validateCPF(customerCPF string) error {
	if customerCPF == "" {
		return domain.ErrCPFIsRequired
	}

	cpfInstance := cpf.NewCPF(customerCPF)

	if !cpfInstance.IsValid() {
		return domain.ErrCPFIsInvalid
	}

	return nil
}

func (c *CustomerService) SignUp(customer *domain.Customer) (*domain.Customer, error) {
	if err := c.validateCPF(customer.CPF); err != nil {
		return nil, err
	}

	return c.customerRepository.SignUp(customer)
}

func (c *CustomerService) Identify(cpf string) (*domain.Customer, error) {
	if err := c.validateCPF(cpf); err != nil {
		return nil, err
	}

	return c.customerRepository.Identify(cpf)
}
