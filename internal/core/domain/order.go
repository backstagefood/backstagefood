package domain

import (
	"fmt"
	"strings"
	"time"
)

type OrderStatus string

const (
	PENDING        OrderStatus = "PENDING"
	RECEIVED       OrderStatus = "RECEIVED"
	PAYMENT_FAILED OrderStatus = "PAYMENT_FAILED"
	IN_PREPARATION OrderStatus = "IN_PREPARATION"
	READY          OrderStatus = "READY"
	COMPLETED      OrderStatus = "COMPLETED"
	CANCELLED      OrderStatus = "CANCELLED"
)

func (o *OrderStatus) GetOrderStatus(current string) (OrderStatus, error) {
	switch strings.ToUpper(current) {
	case string(PENDING):
		return PENDING, nil
	case string(RECEIVED):
		return RECEIVED, nil
	case string(PAYMENT_FAILED):
		return PAYMENT_FAILED, nil
	case string(IN_PREPARATION):
		return IN_PREPARATION, nil
	case string(READY):
		return READY, nil
	case string(COMPLETED):
		return COMPLETED, nil
	case string(CANCELLED):
		return CANCELLED, nil
	default:
		return "", fmt.Errorf("invalid order status")
	}
}

type Order struct {
	ID                   string     `json:"id"`
	Customer             Customer   `json:"customer"`
	Status               string     `json:"status"`
	NotificationAttempts int        `json:"notification_attempts"`
	NotifiedAt           *time.Time `json:"notified_at,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	Products             []Product  `json:"products,omitempty"`
}
