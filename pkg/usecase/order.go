package usecase

import (
	"eatry/pkg/domain"
	"eatry/pkg/helper"
	interfaces "eatry/pkg/repository/interfaces"
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/models"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"github.com/jung-kurt/gofpdf"
)

type OrderUseCase struct {
	orderRepository  interfaces.OrderRepository
	userRepository   interfaces.UserRepository
	cartRepository   interfaces.CartRepository
	walletRepository interfaces.WalletRepository
}

func NewOrderUseCase(orderRepository interfaces.OrderRepository, userRepository interfaces.UserRepository, cartRepository interfaces.CartRepository, walletRepository interfaces.WalletRepository) services.OrderUseCase {
	return &OrderUseCase{
		orderRepository:  orderRepository,
		userRepository:   userRepository,
		cartRepository:   cartRepository,
		walletRepository: walletRepository,
	}
}
func (o *OrderUseCase) OrderItemsFromCart(orderFromCart models.OrderFromCart, userID int) (domain.OrderSuccessResponse, error) {
	var orderBody models.OrderIncoming
	err := copier.Copy(&orderBody, &orderFromCart)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	orderBody.UserID = int(userID)
	cartExist, err := o.orderRepository.DoesCartExist(userID)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	if !cartExist {
		return domain.OrderSuccessResponse{}, errors.New("cart empty can't order")
	}
	addressExist, err := o.orderRepository.AddressExist(orderBody)
	if err != nil {

		return domain.OrderSuccessResponse{}, err
	}

	if !addressExist {
		return domain.OrderSuccessResponse{}, errors.New("address does not exist")
	}

	cartItems, err := o.cartRepository.GetAllItemsFromCart(orderBody.UserID)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	var orderDetails domain.Order
	var orderItemDetails domain.UserOrderItem

	orderDetails = helper.CopyOrderDetails(orderDetails, orderBody)

	for _, c := range cartItems {
		orderDetails.GrandTotal += c.TotalPrice
	}
	orderDetails.FinalPrice = orderDetails.GrandTotal

	//for cash on delivery
	if orderBody.PaymentID == 1 {

		if orderDetails.FinalPrice > 1000 {
			return domain.OrderSuccessResponse{}, errors.New("cash on delivery is not possible")
		}
		orderDetails.PaymentStatus = "not paid"
		orderDetails.ShipmentStatus = "pending"
	}
	//wallet
	if orderBody.PaymentID == 3 {

		walletAvailable, err := o.orderRepository.GetWalletAmount(orderBody.UserID)
		if err != nil {
			return domain.OrderSuccessResponse{}, err
		}

		// if wallet amount is less than final amount - make payment status - not paid and shipment status pending
		if walletAvailable < orderDetails.FinalPrice {
			orderDetails.PaymentStatus = "not paid"
			orderDetails.ShipmentStatus = "pending"
			return domain.OrderSuccessResponse{}, errors.New("wallet amount is less than total amount")
		} else {
			o.orderRepository.UpdateWalletAmount(walletAvailable-orderDetails.FinalPrice, orderBody.UserID)
			orderDetails.PaymentStatus = "paid"
		}

	}

	err = o.orderRepository.CreateOrder(orderDetails)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	for _, c := range cartItems {
		// for each order save details of products and associated details and use order_id as foreign key ( for each order multiple product will be there)
		orderItemDetails.OrderID = orderDetails.OrderId
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = int(c.Quantity)
		orderItemDetails.TotalPrice = c.TotalPrice

		err := o.orderRepository.AddOrderItems(orderItemDetails, orderDetails.UserID, uint(c.ProductID), c.Quantity)
		if err != nil {
			return domain.OrderSuccessResponse{}, err
		}

	}
	orderSuccessResponse, err := o.orderRepository.GetBriefOrderDetails(orderDetails.OrderId)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	return orderSuccessResponse, nil

}

func (o *OrderUseCase) GetOrderDetails(userID int, page int, count int) ([]models.FullOrderDetails, error) {

	fullOrderDetails, err := o.orderRepository.GetOrderDetails(userID, page, count)
	if err != nil {
		return []models.FullOrderDetails{}, err
	}

	return fullOrderDetails, nil

}
func (o *OrderUseCase) CancelOrders(orderID string, userID int) error {
	userTest, err := o.orderRepository.UserOrderRelationship(orderID, userID)
	if err != nil {
		return err
	}
	if userTest != userID {
		return errors.New("the order is not done by this user")
	}
	orderProductDetails, err := o.orderRepository.GetProductDetailsFromOrders(orderID)
	if err != nil {
		return err
	}
	paymentStatus, err := o.orderRepository.GetPaymentStatus(orderID)
	if err != nil {
		return errors.New("cannot show the payment status")
	}

	shipmentStatus, err := o.orderRepository.GetShipmentStatus(orderID)
	if err != nil {
		return err
	}
	if shipmentStatus == "delivered" {
		return errors.New("item already delivered, cannot cancel")
	}

	if shipmentStatus == "returned" || shipmentStatus == "return" {
		message := fmt.Sprint(shipmentStatus)

		return errors.New("the order is in" + message + ", so no point in cancelling")
	}

	if shipmentStatus == "cancelled" {
		return errors.New("the order is already cancelled, so no point in cancelling")
	}
	err = o.orderRepository.CancelOrders(orderID)
	if err != nil {
		return err
	}
	var walletHistory models.WalletHistory
	if paymentStatus == "paid" {
		amount, err := o.orderRepository.GetPriceoftheproduct(orderID)
		if err != nil {
			return err

		}
		walletHistory.WalletAmount = amount
		walletData, err := o.walletRepository.GetWalletData(userID)
		if err != nil {
			return err

		}
		walletHistory.WalletID = walletData.WalletID
		walletHistory.OrderID = orderID
		walletHistory.Status = "CREDITED"

		err = o.walletRepository.AddtoWallet(userID, amount)
		if err != nil {
			return err
		}
		err = o.walletRepository.AddToWalletHistory(walletHistory)
		if err != nil {
			return err
		}

	}
	err = o.orderRepository.UpdateQuantityOfProduct(orderProductDetails)
	if err != nil {
		return err
	}

	return nil

}

func (o *OrderUseCase) GetAllOrderDetailsForAdmin(page int) ([]models.CombinedOrderDetails, error) {

	orderDetails, err := o.orderRepository.GetOrderDetailsBrief(page)
	if err != nil {
		return []models.CombinedOrderDetails{}, err
	}
	return orderDetails, nil

}

func (o *OrderUseCase) CancelOrderFromAdminSide(orderID string) error {

	orderProducts, err := o.orderRepository.GetProductDetailsFromOrders(orderID)
	if err != nil {
		return err
	}

	err = o.orderRepository.CancelOrders(orderID)
	if err != nil {
		return err
	}

	// update the quantity to products since the order is cancelled
	err = o.orderRepository.UpdateQuantityOfProduct(orderProducts)
	if err != nil {
		return err
	}

	return nil

}
func (o *OrderUseCase) ApproveOrder(orderID string) error {

	// check whether the specified orderID exist
	ok, err := o.orderRepository.CheckOrderID(orderID)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("order id does not exist")
	}

	// check the shipment status - if the status cancelled, don't approve it
	shipmentStatus, err := o.orderRepository.GetShipmentStatus(orderID)
	if err != nil {
		return err
	}

	if shipmentStatus == "cancelled" {

		return errors.New("the order is cancelled, cannot approve it")
	}

	if shipmentStatus == "pending" {

		return errors.New("the order is pending, cannot approve it")
	}

	if shipmentStatus == "processing" {

		err := o.orderRepository.ApproveOrder(orderID)

		if err != nil {
			return err
		}
	}
	return nil

}

//------------------PDF--------------------

func (o *OrderUseCase) GenerateInvoice(orderID string, userID int) (*gofpdf.Fpdf, error) {
	ok, err := o.orderRepository.CheckOrderID(orderID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("orderId does not exists")
	}
	order, err := o.orderRepository.GetOrderDetailsByID(orderID)
	if err != nil {
		return nil, err
	}
	items, err := o.orderRepository.GetItemsByOrderId(orderID)
	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 30)
	pdf.SetTextColor(31, 73, 125)
	pdf.Cell(0, 20, "Invoice")
	pdf.Ln(20)

	pdf.SetFont("Arial", "I", 14)
	pdf.SetTextColor(51, 51, 51)
	pdf.Cell(0, 10, "Customer Details")
	pdf.Ln(10)

	customerDetails := []string{
		"Name: " + order.Name,
		"Email" + order.Email,
		"Phone" + order.Phone,
		"House Name: " + order.HouseName,
		"Street: " + order.Street,
		"City: " + order.City,
		"pincode:" + order.Pin,
		"State: " + order.State,
	}
	for _, detail := range customerDetails {
		pdf.Cell(0, 10, detail)
		pdf.Ln(10)
	}
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 16)
	pdf.SetFillColor(217, 217, 217)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(40, 10, "Item", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Price", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Quantity", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Final Price", "1", 0, "C", true, 0, "")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	pdf.SetFillColor(255, 255, 255)
	for _, item := range items {
		pdf.CellFormat(40, 10, item.Name, "1", 0, "L", true, 0, "")
		pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(item.Price, 'f', 2, 64), "1", 0, "C", true, 0, "")
		pdf.CellFormat(40, 10, strconv.Itoa(item.Quantity), "1", 0, "C", true, 0, "")
		pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(item.Price*float64(item.Quantity), 'f', 2, 64), "1", 0, "C", true, 0, "")
		pdf.Ln(10)
	}
	pdf.Ln(10)

	var totalPrice float64
	for _, item := range items {
		totalPrice += item.Price * float64(item.Quantity)
	}
	pdf.SetFont("Arial", "B", 16)
	pdf.SetFillColor(217, 217, 217)
	pdf.CellFormat(120, 10, "Total Price:", "1", 0, "R", true, 0, "")
	pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(totalPrice, 'f', 2, 64), "1", 0, "C", true, 0, "")
	pdf.Ln(10)

	pdf.SetFont("Arial", "I", 12)
	pdf.Cell(0, 10, "Generated by Eatry India Pvt Ltd. - "+time.Now().Format("2006-01-02 15:04:05"))
	pdf.Ln(10)

	return pdf, nil
}
