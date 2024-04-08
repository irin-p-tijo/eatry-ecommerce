package handlers

import (
	"eatry/pkg/domain"
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentUseCase services.PaymentUseCase
}

func NewPaymentHandler(usecase services.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: usecase,
	}
}
func (pay *PaymentHandler) AddPaymentMethods(c *gin.Context) {
	var addpayment domain.PaymentMethod

	if err := c.ShouldBindJSON(&addpayment); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	paymentresponse, err := pay.paymentUseCase.AddPaymentMethods(addpayment)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not add the payment methods", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the paymentmethod is added", paymentresponse, nil)
	c.JSON(http.StatusOK, successRes)
}
func (pay *PaymentHandler) DeletePaymentMethods(c *gin.Context) {
	paymentID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "parameters provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err1 := pay.paymentUseCase.DeletePaymentMethods(paymentID)
	if err1 != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not delete the payment methods", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the paymentmethod is deleetd", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (pay *PaymentHandler) GetAllPaymentMethods(c *gin.Context) {
	pagestr := c.Query("page")
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check the variables", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	pagesize, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the parameters provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	payment, err := pay.paymentUseCase.GetAllPaymentMethods(page, pagesize)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrive the data", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "the product is deleted", payment, nil)
	c.JSON(http.StatusOK, successRes)
}
func (pay *PaymentHandler) MakePaymentRazorPay(c *gin.Context) {
	orderID := c.Query("id")
	userID := c.Query("user_id")
	user_Id, _ := strconv.Atoi(userID)
	fmt.Println("order id is ", orderID)
	orderDetail, razorID, err := pay.paymentUseCase.MakePaymentRazorPay(orderID, user_Id)
	if err != nil {

		if strings.Contains(err.Error(), "Payment failed") {
			errorRes := response.ClientResponse(http.StatusInternalServerError, "Payment failed", nil, err.Error())
			c.JSON(http.StatusInternalServerError, errorRes)
			return
		} else {
			errorRes := response.ClientResponse(http.StatusInternalServerError, "could not generate order details", nil, err.Error())
			c.JSON(http.StatusInternalServerError, errorRes)
			return
		}

	}
	fmt.Println("orderDetails :", orderDetail)
	fmt.Println("order id is ", orderID)
	fmt.Println("razorID:", razorID)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"final_price": orderDetail.FinalPrice * 100,
		"razor_id":    razorID,
		"user_id":     userID,
		"order_id":    orderDetail.OrderId,
		"user_name":   orderDetail.Name,
		"total":       int(orderDetail.FinalPrice),
	})
}
func (p *PaymentHandler) VerifyPayment(c *gin.Context) {

	orderID := c.Query("order_id")
	paymentID := c.Query("payment_id")
	razorID := c.Query("razor_id")
	fmt.Println("payment id is ", paymentID)
	err := p.paymentUseCase.SavePaymentDetails(paymentID, razorID, orderID)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not update payment details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated payment details", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
