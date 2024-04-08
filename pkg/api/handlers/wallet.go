package handlers

import (
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	walletUseCase services.WalletUseCase
}

func NewWalletHandler(walletusecase services.WalletUseCase) *WalletHandler {
	return &WalletHandler{
		walletUseCase: walletusecase,
	}
}
func (wa *WalletHandler) GetWallet(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Check the parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	wallet, err := wa.walletUseCase.GetWallet(userID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "Wallet is successfully shown", wallet, nil)
	c.JSON(http.StatusCreated, successRes)

}
func (wa *WalletHandler) WalletHistory(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Check the parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	walletHistory, err := wa.walletUseCase.WalletHistory(userID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "wallet history cannot be retrived", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "Wallet is successfully shown", walletHistory, nil)
	c.JSON(http.StatusCreated, successRes)
}
