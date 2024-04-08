package handlers

import (
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartUseCase services.CartUseCase
}

func NewCartHandler(usecase services.CartUseCase) *CartHandler {
	return &CartHandler{
		cartUseCase: usecase,
	}
}
func (rt *CartHandler) AddToCart(c *gin.Context) {

	productID, err := strconv.Atoi(c.Query("product_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the parameters are given wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the parameters are given wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	cart, err := rt.cartUseCase.AddToCart(userID, productID)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the product cannot be added to cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the product is added to the cart successfully", cart, nil)
	c.JSON(http.StatusBadRequest, successRes)
}

func (rt *CartHandler) RemoveFromCart(c *gin.Context) {
	productID, err := strconv.Atoi(c.Query("product_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the parameters are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the parameters are wrong", nil, err.Error())

		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	updatecart, err := rt.cartUseCase.RemoveFromCart(userID, productID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not remove the product from cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	succesRes := response.ClientResponse(http.StatusOK, "product removed successfully", updatecart, nil)
	c.JSON(http.StatusOK, succesRes)

}

func (rt *CartHandler) DisplayCart(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check the parameters", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	cart, err := rt.cartUseCase.DisplayCart(userID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not display cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Cart items displayed successfully", cart, nil)
	c.JSON(http.StatusOK, successRes)
}
