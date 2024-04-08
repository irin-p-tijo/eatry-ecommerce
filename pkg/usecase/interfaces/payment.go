package interfaces

import (
	"eatry/pkg/domain"
	"eatry/pkg/utils/models"
)

type PaymentUseCase interface {
	AddPaymentMethods(payment domain.PaymentMethod) (domain.PaymentMethod, error)
	DeletePaymentMethods(paymentID int) error
	GetAllPaymentMethods(page int, count int) ([]domain.PaymentMethod, error)
	MakePaymentRazorPay(orderID string, user_Id int) (models.CombinedOrderDetails, string, error)
	SavePaymentDetails(paymentID string, razorID string, orderID string) error
}
