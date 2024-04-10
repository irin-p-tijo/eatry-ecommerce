package handlers

import (
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/models"
	"eatry/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// UserSignUp godoc
// @Summary User sign up
// @Description Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserDetails true "User details"
// @Success 201 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /users/signup [post]
func (u *UserHandler) UserSignUp(c *gin.Context) {

	var user models.UserDetails

	// bind the user details to the struct
	if err := c.ShouldBindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	// checking whether the data sent by the user has all the correct constraints specified by Users struct
	err := validator.New().Struct(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)

		return
	}
	userCreated, err := u.userUseCase.UserSignUp(user)

	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "User could not signed up", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "User successfully signed up", userCreated, nil)
	c.JSON(http.StatusCreated, successRes)

}

// LoginHandler godoc
// @Summary User login
// @Description Log in a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserLogin true "User credentials"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /users/login [post]
func (u *UserHandler) LoginHandler(c *gin.Context) {

	var user models.UserLogin

	if err := c.ShouldBindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := validator.New().Struct(user)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	userdetails, err := u.userUseCase.LoginHandler(user)

	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "User could not be logged in", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "User successfully logged in", userdetails, nil)
	c.JSON(http.StatusOK, successRes)

}
func (u *UserHandler) AddAddress(c *gin.Context) {

	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the parameters given are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	var address models.AddAddress
	if err := c.BindJSON(&address); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := u.userUseCase.AddAddress(userID, address); err != nil {

		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added address", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (u *UserHandler) DeleteAddress(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the parameters given are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	addressID, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the parameters given are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := u.userUseCase.DeleteAddress(userID, addressID); err != nil {

		errRes := response.ClientResponse(http.StatusBadRequest, "the address cannot be deleted ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the address is deleted", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (u *UserHandler) GetAllAddress(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user_id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	address, err := u.userUseCase.GetAllAddress(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not open checkout", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", address, nil)
	c.JSON(http.StatusOK, successRes)
}
func (u *UserHandler) UserProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints are given wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userdetails, err := u.userUseCase.UserProfile(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could retrive the userprofile details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the userprofie dtails are retrived", userdetails, nil)
	c.JSON(http.StatusOK, successRes)
}
func (u *UserHandler) CheckOut(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "parameters are given wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	checkout, err := u.userUseCase.CheckOut(userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Checkout Page loaded successfully", checkout, nil)
	c.JSON(http.StatusOK, successRes)
}
