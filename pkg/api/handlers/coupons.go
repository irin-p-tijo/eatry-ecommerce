package handlers

import (
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/models"
	"eatry/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	usecase services.CouponUseCase
}

func NewCouponHandler(use services.CouponUseCase) *CouponHandler {
	return &CouponHandler{
		usecase: use,
	}
}
func (coup *CouponHandler) CreateNewCoupon(c *gin.Context) {
	var coupon models.AddCoupon
	if err := c.BindJSON(&coupon); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := coup.usecase.AddCoupon(coupon)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Coupon", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
func (coup *CouponHandler) MakeCouponInvalid(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := coup.usecase.MakeCouponInvalid(id); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Coupon cannot be made invalid", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully made Coupon as invalid", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
func (co *CouponHandler) GetAllCoupons(c *gin.Context) {

	coupons, err := co.usecase.GetAllCoupons()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in getting coupons", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully got all coupons", coupons, nil)
	c.JSON(http.StatusOK, successRes)

}
