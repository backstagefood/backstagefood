package service

type CustomerInterface interface {
}

type CustomerService struct {
	customerRepository CustomerInterface
}

type CustomerDTO struct {
}
