package handlers

import (
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	wishlistUseCase services.WishlistUseCase
}

func NewWishlistHandler(usecase services.WishlistUseCase) *WishlistHandler {
	return &WishlistHandler{
		wishlistUseCase: usecase,
	}
}
func (w *WishlistHandler) AddToWishList(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))
	id := c.Query("id")
	productID, err := strconv.Atoi(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "product id is in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = w.wishlistUseCase.AddToWishList(productID, userID)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to item to the wishlist", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "SuccessFully added product to the wishlist", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
func (w *WishlistHandler) RemoveFromWishList(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))
	id := c.Query("id")
	productID, err := strconv.Atoi(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "product id is in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = w.wishlistUseCase.RemoveFromWishList(productID, userID)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to remove item from wishlist", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "SuccessFully deleted product from wishlist", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
func (w *WishlistHandler) GetWishList(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))
	wishList, err := w.wishlistUseCase.GetWishList(userID)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve wishlist detailss", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "SuccessFully retrieved wishlist", wishList, nil)
	c.JSON(http.StatusOK, successRes)

}
