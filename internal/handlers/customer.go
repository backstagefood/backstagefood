package handlers

import (
	"github.com/backstagefood/backstagefood/internal/core/domain"
	portService "github.com/backstagefood/backstagefood/internal/core/ports/services"
	"github.com/backstagefood/backstagefood/internal/core/services"
	"github.com/backstagefood/backstagefood/internal/repositories"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CustomerHandler struct {
	customerService portService.Customer
}

func NewCustomerHandler(databaseConnection *repositories.ApplicationDatabase) *CustomerHandler {
	customerRepository := repositories.NewCustomerRepository(databaseConnection)

	return &CustomerHandler{
		customerService: services.NewCustomerService(customerRepository),
	}
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
		CreatedAt: domain.GetNow(),
		UpdatedAt: domain.GetNow(),
	}
}

// CustomerSignUp godoc
// @Summary Customer sign up
// @Description Create a customer.
// @Tags customers
// @Produce json
// @Param customer body handlers.SignUpCustomerDTO true "SignUpCustomerDTO object"
// @Success 201 {object} handlers.CustomerDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customers/sign-up [post]
func (h *CustomerHandler) CustomerSignUp(c echo.Context) error {
	productDTO := new(SignUpCustomerDTO)
	if err := c.Bind(productDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	customer, err := h.customerService.SignUp(productDTO.ToDomainCustomer())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, CustomerDTO{
		ID:    customer.ID,
		Name:  customer.Name,
		CPF:   customer.CPF,
		Email: customer.Email,
	})
}

// CustomerIdentify godoc
// @Summary Customer identification
// @Description Identify a customer.
// @Tags customers
// @Produce json
// @Param cpf path string true "Customer CPF"
// @Success 200 {object} handlers.CustomerDTO
// @Failure 500 {object} map[string]string
// @Router /customers/{cpf} [get]
func (h *CustomerHandler) CustomerIdentify(c echo.Context) error {
	cpf := c.Param("cpf")

	customer, err := h.customerService.Identify(cpf)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, CustomerDTO{
		ID:    customer.ID,
		Name:  customer.Name,
		CPF:   customer.CPF,
		Email: customer.Email,
	})
}
