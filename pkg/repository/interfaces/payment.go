package interfaces

import (
	"eatry/pkg/domain"
	"eatry/pkg/utils/models"
)

type PaymentRepository interface {
	AddPaymentMethods(addpayment domain.PaymentMethod) (domain.PaymentMethod, error)
	DeletePaymentMethods(paymentID int) error
	GetAllPaymentMethods(page int, count int) ([]domain.PaymentMethod, error)
	AddRazorPayDetails(orderID string, razorPayOrderID string) error
	GetOrderDetailsByOrderId(orderID string) (models.CombinedOrderDetails, error)
	UpdateShipmentAndPaymentByOrderID(shipmentStatus string, paymentStatus string, orderID string) error
	CheckPaymentStatus(orderID string) (string, error)
	UpdatePaymentDetails(orderID string, paymentID string) error
}
