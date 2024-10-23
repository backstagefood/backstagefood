package domain

import (
	"time"
)

type OrderStatus string

const (
	PENDING        OrderStatus = "PENDING"
	RECEIVED       OrderStatus = "RECEIVED"
	PAYMENT_FAILED OrderStatus = "PAYMENT_FAILED"
)

type Order struct {
	ID                   string     `json:"id"`
	CustomerID           string     `json:"id_customer"`
	Status               string     `json:"status"`
	NotificationAttempts int        `json:"notification_attempts"`
	NotifiedAt           *time.Time `json:"notified_at,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	Products             []Product  `json:"products,omitempty"`
}
