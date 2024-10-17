package service

import (
	"github.com/backstagefood/backstagefood/internal/domain"
	"github.com/backstagefood/backstagefood/internal/repositories"
	"github.com/google/uuid"
	"time"
)

type CustomerService struct {
	customerRepository repositories.CustomerRepository
}

type SignUpCustomerDTO struct {
	Name  string `json:"name"`
	CPF   string `json:"cpf"`
	Email string `json:"email"`
}

type CustomerDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	CPF   string `json:"cpf"`
	Email string `json:"email"`
}

func (customerDTO *SignUpCustomerDTO) ToDomainCustomer() *domain.Customer {
	return &domain.Customer{
		ID:        uuid.New().String(),
		Name:      customerDTO.Name,
		CPF:       customerDTO.CPF,
		Email:     customerDTO.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewCustomerService(repository repositories.CustomerRepository) *CustomerService {
	return &CustomerService{customerRepository: repository}
}

func (c *CustomerService) SignUp(customerDTO *domain.Customer) (*domain.Customer, error) {
	if customerDTO.CPF == "" {
		return nil, domain.ErrCPFIsRequired
	}

	return c.customerRepository.SignUp(customerDTO)
}

func (c *CustomerService) Identify(cpf string) (*domain.Customer, error) {
	if cpf == "" {
		return nil, domain.ErrCPFIsRequired
	}

	return c.customerRepository.Identify(cpf)
}
