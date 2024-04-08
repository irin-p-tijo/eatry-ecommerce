package usecase

import (
	"eatry/pkg/domain"
	interfaces "eatry/pkg/repository/interfaces"
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/models"
	"errors"

	"github.com/razorpay/razorpay-go"
)

type PaymentUsecase struct {
	paymentRepository interfaces.PaymentRepository
	orderRepository   interfaces.OrderRepository
}

func NewPaymentUseCase(repo interfaces.PaymentRepository, orderRepo interfaces.OrderRepository) services.PaymentUseCase {
	return &PaymentUsecase{
		paymentRepository: repo,
		orderRepository:   orderRepo,
	}
}
func (p *PaymentUsecase) AddPaymentMethods(payment domain.PaymentMethod) (domain.PaymentMethod, error) {
	payments, err := p.paymentRepository.AddPaymentMethods(payment)
	if err != nil {
		return domain.PaymentMethod{}, err
	}
	return payments, nil
}
func (p *PaymentUsecase) DeletePaymentMethods(paymentID int) error {
	err := p.paymentRepository.DeletePaymentMethods(paymentID)
	if err != nil {
		return err
	}
	return nil
}
func (p *PaymentUsecase) GetAllPaymentMethods(page int, count int) ([]domain.PaymentMethod, error) {
	payment, err := p.paymentRepository.GetAllPaymentMethods(page, count)
	if err != nil {
		return []domain.PaymentMethod{}, err
	}
	return payment, nil
}

// /////////////////RAZOR PAY////////////////////////
func (p *PaymentUsecase) MakePaymentRazorPay(orderID string, user_Id int) (models.CombinedOrderDetails, string, error) {

	combinedOrderDetails, err := p.paymentRepository.GetOrderDetailsByOrderId(orderID)
	if err != nil {
		return models.CombinedOrderDetails{}, "", err
	}
	client := razorpay.NewClient("rzp_test_jkUtyHkUKEWZte", "R097URXZzlQjfuH8ROJGbws3")

	data := map[string]interface{}{
		"amount":   int(combinedOrderDetails.FinalPrice) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		return models.CombinedOrderDetails{}, "", err
	}
	razorPayOrderID := body["id"].(string)

	err = p.paymentRepository.AddRazorPayDetails(orderID, razorPayOrderID)
	if err != nil {
		return models.CombinedOrderDetails{}, "", err
	}

	return combinedOrderDetails, razorPayOrderID, nil
}
func (p *PaymentUsecase) SavePaymentDetails(paymentID string, razorID string, orderID string) error {

	// to check whether the order is already paid
	status, err := p.paymentRepository.CheckPaymentStatus(orderID)
	if err != nil {
		return err
	}
	if status == "failed" {
		return errors.New("payment failed in razorpay")
	}
	if status == "not paid" {

		err = p.paymentRepository.UpdatePaymentDetails(orderID, paymentID)
		if err != nil {
			return err
		}

		err := p.paymentRepository.UpdateShipmentAndPaymentByOrderID("processing", "paid", orderID)
		if err != nil {
			return err
		}

		return nil

	}

	return errors.New("already paid")

}
