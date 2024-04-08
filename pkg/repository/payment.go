package repository

import (
	"eatry/pkg/domain"
	interfaces "eatry/pkg/repository/interfaces"
	"eatry/pkg/utils/models"
	"errors"

	"gorm.io/gorm"
)

type paymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) interfaces.PaymentRepository {
	return &paymentRepository{
		DB: db,
	}
}
func (pay *paymentRepository) AddPaymentMethods(addpayment domain.PaymentMethod) (domain.PaymentMethod, error) {
	var payments domain.PaymentMethod
	err := pay.DB.Raw("insert into payment_methods (payment_name) values (?) returning id", addpayment.Payment_Name).Scan(&payments).Error
	if err != nil {
		return domain.PaymentMethod{}, err
	}
	return payments, err

}
func (pay *paymentRepository) DeletePaymentMethods(paymentID int) error {
	err := pay.DB.Exec("delete from payment_methods where id=?", paymentID)
	if err.RowsAffected < 1 {
		return errors.New("could not delete the paymentmehods")
	}
	return nil
}
func (pay *paymentRepository) GetAllPaymentMethods(page int, count int) ([]domain.PaymentMethod, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var payments []domain.PaymentMethod
	err := pay.DB.Raw("select id, payment_name from payment_methods limit ? offset ?", count, offset).Scan(&payments).Error
	if err != nil {
		return []domain.PaymentMethod{}, err
	}
	return payments, nil
}
func (pay *paymentRepository) AddRazorPayDetails(orderID string, razorPayOrderID string) error {

	err := pay.DB.Exec("insert into razor_pays (order_id,razor_id) values (?,?)", orderID, razorPayOrderID).Error
	if err != nil {
		return err
	}
	return nil
}
func (pay *paymentRepository) GetOrderDetailsByOrderId(orderID string) (models.CombinedOrderDetails, error) {
	var orderDetails models.CombinedOrderDetails
	err := pay.DB.Raw("select  orders.order_id,orders.final_price,orders.shipment_status,orders.payment_status,users.name,users.email,users.phone,addresses.house_name,addresses.street,addresses.city,addresses.district,addresses.state,addresses.pin from orders join users on  orders.user_id = users.id join addresses on orders.address_id=addresses.id where orders.order_id=? ", orderID).Scan(&orderDetails).Error

	if err != nil {
		return models.CombinedOrderDetails{}, nil
	}
	return orderDetails, nil
}
func (o *paymentRepository) CheckPaymentStatus(orderID string) (string, error) {

	var paymentStatus string
	err := o.DB.Raw("select payment_status from orders where order_id = ? ", orderID).Scan(&paymentStatus).Error
	if err != nil {
		return "", err
	}

	return paymentStatus, nil

}

func (o *paymentRepository) UpdateShipmentAndPaymentByOrderID(shipmentStatus string, paymentStatus string, orderID string) error {

	err := o.DB.Exec("update orders set payment_status = ?, shipment_status = ? where order_id = ?", paymentStatus, shipmentStatus, orderID).Error
	if err != nil {
		return err
	}

	return nil

}

func (o *paymentRepository) UpdatePaymentDetails(orderID string, paymentID string) error {

	err := o.DB.Exec("update razor_pays set payment_id = ? where order_id = ?", paymentID, orderID).Error
	if err != nil {
		return err
	}
	return nil

}
