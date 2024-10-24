package services

import (
	"database/sql"
	"errors"
	"net/http"

	portRepository "github.com/backstagefood/backstagefood/internal/core/ports/repositories"
	portService "github.com/backstagefood/backstagefood/internal/core/ports/services"
	"github.com/backstagefood/backstagefood/pkg/transaction"
	"github.com/google/uuid"

	"github.com/backstagefood/backstagefood/internal/core/domain"
	paymentgateway "github.com/backstagefood/backstagefood/internal/core/services/payment_gateway"
)

var (
	errorOrderPending  = errors.New("order still pending")
	errorPaymentFailed = errors.New("payment failed")
	errorInsertOrder   = errors.New("create order failed")
	errorInvalidUUID   = errors.New("invalid uuid")
	errorDeleteOrder   = errors.New("error deleting an order")
)

type OrderService struct {
	transactionManager transaction.TransactionManagerInterface
	orderRepository    portRepository.Order
}

func NewOrderService(
	orderRepository portRepository.Order,
	transactionManager transaction.TransactionManagerInterface,
) portService.Order {
	return &OrderService{
		orderRepository:    orderRepository,
		transactionManager: transactionManager,
	}
}

func (o *OrderService) MakeCheckout(orderId string) (*portService.CheckoutServiceDTO, error) {
	transactionResult, err := o.transactionManager.RunWithTransaction(func(tx *sql.Tx) (interface{}, error) {
		updatedOrder, err := o.orderRepository.UpdateOrderStatus(tx, orderId)
		if err != nil {
			return &portService.CheckoutServiceDTO{
				PaymentSucceeded: true,
				OrderStatus:      domain.PENDING,
				Order:            updatedOrder,
			}, errorOrderPending
		}

		// TODO: FakeCheckout() need to be interfaced when the real webhook is implemented.
		paymentGatewayResponse := paymentgateway.PaymentCheckout()
		if paymentGatewayResponse != http.StatusOK {
			return &portService.CheckoutServiceDTO{
				PaymentSucceeded: false,
				OrderStatus:      domain.PAYMENT_FAILED,
			}, errorPaymentFailed
		}

		return &portService.CheckoutServiceDTO{
			PaymentSucceeded: true,
			OrderStatus:      domain.RECEIVED,
			Order:            updatedOrder,
		}, nil
	})

	return transactionResult.(*portService.CheckoutServiceDTO), err
}

func (o *OrderService) GetOrders() ([]*domain.Order, error) {
	return o.orderRepository.ListOrders()
}

func (o *OrderService) CreateOrder(order *domain.Order) (map[string]string, error) {
	// todo validar mascara e existencia dos campos recebidos: ID_CUSTOMER e LISTA DE PRODUTOS
	createOrder, err := o.orderRepository.CreateOrder(order)
	if err != nil {
		return nil, errorInsertOrder
	}
	return createOrder, err
}

func (o *OrderService) DeleteOrder(orderId string) error {
	if err := uuid.Validate(orderId); err != nil {
		return errorInvalidUUID
	}

	if err := o.orderRepository.DeleteOrder(orderId); err != nil {
		return errorDeleteOrder
	}

	return nil
}
