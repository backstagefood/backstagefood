package domain

import (
	"errors"
	"time"
)

var (
	ErrCPFIsRequired = errors.New("cpf is required")
)

type Customer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
