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

// @Summary Add a product to cart
// @Description Adds a product to a user's shopping cart
// @Tags cart
// @Accept json
// @Produce json
// @Param product_id query int true "Product ID"
// @Param user_id query int true "User ID"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /cart/add [post]
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

// @Summary Remove a product from cart
// @Description Removes a product from a user's shopping cart
// @Tags cart
// @Accept json
// @Produce json
// @Param product_id query int true "Product ID"
// @Param user_id query int true "User ID"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /cart/remove [delete]
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

// @Summary Get all items in a user's cart
// @Description Retrieves a list of all products in a user's shopping cart
// @Tags cart
// @Accept json
// @Produce json
// @Param user_id query int true "User ID"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /cart [get]
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
