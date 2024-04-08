package handlers

import (
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/models"
	"eatry/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase services.OrderUseCase
}

func NewOrderHandler(useCase services.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: useCase,
	}
}
func (o *OrderHandler) OrderItemsFromCart(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the parameters are given wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var orderFromCart models.OrderFromCart
	if err := c.ShouldBindJSON(&orderFromCart); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	orderSuccessResponse, err := o.orderUseCase.OrderItemsFromCart(orderFromCart, userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully created the order", orderSuccessResponse, nil)
	c.JSON(http.StatusOK, successRes)
}

func (o *OrderHandler) GetOrderDetails(c *gin.Context) {

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	pageSize, err := strconv.Atoi(c.Query("count"))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	userID, _ := strconv.Atoi(c.Query("user_id"))

	fullOrderDetails, err := o.orderUseCase.GetOrderDetails(userID, page, pageSize)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Full Order Details", fullOrderDetails, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *OrderHandler) CancelOrder(c *gin.Context) {

	orderID := c.Query("id")
	fmt.Println(orderID)

	userID, _ := strconv.Atoi(c.Query("user_id"))

	err := o.orderUseCase.CancelOrders(orderID, userID)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not cancel the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Cancel Successfull", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
func (o *OrderHandler) GetAllOrderDetailsForAdmin(c *gin.Context) {

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	allOrderDetails, err := o.orderUseCase.GetAllOrderDetailsForAdmin(page)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not retrieve order details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Order Details Retrieved successfully", allOrderDetails, nil)
	c.JSON(http.StatusOK, successRes)

}
func (o *OrderHandler) ApproveOrder(c *gin.Context) {

	orderID := c.Query("order_id")

	err := o.orderUseCase.ApproveOrder(orderID)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not approve the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Order approved successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

//---------------------------PDF-------------------------------

func (o *OrderHandler) GenerateInvoice(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Check the parametrs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	orderID := c.Query("order_id")

	pdf, err := o.orderUseCase.GenerateInvoice(orderID, userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not generate invoice", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	c.Header("Content-Disposition", "attachment;filename=invoice.pdf")

	pdfFilePath := "salesReport/invoice.pdf"

	err = pdf.OutputFileAndClose(pdfFilePath)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "could not generate the pdf", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	c.Header("Content-Disposition", "attachment; filename=sales_report.pdf")
	c.Header("Content-Type", "application/pdf")

	c.File(pdfFilePath)

	c.Header("Content-Type", "application/pdf")

	err = pdf.Output(c.Writer)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "could not print invoice", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "the invoice is generated", pdf, nil)
	c.JSON(http.StatusOK, successRes)
}
