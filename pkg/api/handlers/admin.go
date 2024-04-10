package handlers

import (
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/models"
	"eatry/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

func NewAdminHandler(usecase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
	}
}

// CreateAdmin godoc
// @Summary Create a new admin
// @Description Creates a new admin account
// @Tags Admin
// @Accept json
// @Produce json
// @Param body body models.AdminSignUp true "Admin signup details"
// @Success 201 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/create [post]

func (ad *AdminHandler) CreateAdmin(c *gin.Context) {
	var admin models.AdminSignUp
	if err := c.ShouldBindJSON(&admin); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	adminDetails, err := ad.adminUseCase.CreateAdmin(admin)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate Admin", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Successfully signed up the user", adminDetails, nil)
	c.JSON(http.StatusCreated, successRes)
}

// LoginHandler godoc
// @Summary Login as admin
// @Description Logs in as an admin
// @Tags Admin
// @Accept json
// @Produce json
// @Param body body models.AdminLogin true "Admin login details"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/login [post]
func (ad *AdminHandler) LoginHandler(c *gin.Context) { // login handler for the admin

	var adminDetails models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	admin, err := ad.adminUseCase.LoginHandler(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)

}

// GetUsers godoc
// @Summary Get users
// @Description Retrieves a list of users
// @Tags Admin
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param count query int true "Number of users per page"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/users [get]
func (ad *AdminHandler) GetUsers(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "page number is not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	pageSize, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user count in a page not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := ad.adminUseCase.GetUsers(page, pageSize)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved the users", users, nil)
	c.JSON(http.StatusOK, successRes)
}

// BlockUser godoc
// @Summary Block user
// @Description Blocks a user
// @Tags Admin
// @Accept json
// @Produce json
// @Param id query int true "User ID"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/block [post]
func (ad *AdminHandler) BlockUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check the parameters", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	err = ad.adminUseCase.BlockUser(userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not block the user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the user is blockes", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// UnBlockUser godoc
// @Summary Unblock user
// @Description Unblocks a user
// @Tags Admin
// @Accept json
// @Produce json
// @Param id query int true "User ID"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/unblock [post]
func (ad *AdminHandler) UnBlockUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check the parameters", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	err = ad.adminUseCase.UnBlockUser(userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not unblock the user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the user is unblocked", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// Dashboard godoc
// @Summary Get admin dashboard
// @Description Retrieves the admin dashboard
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/dashboard [get]
func (ad *AdminHandler) Dashboard(c *gin.Context) {

	adminDashboard, err := ad.adminUseCase.Dashboard()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "dashboard could not be displayed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "admin dashboard displayed fine", adminDashboard, nil)
	c.JSON(http.StatusOK, successRes)
}

// FilterSalesReport godoc
// @Summary Filter sales report
// @Description Filters the sales report by time period
// @Tags Admin
// @Accept json
// @Produce json
// @Param period query string true "Time period"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/sales-report [get]

func (ad *AdminHandler) FilterSalesReport(c *gin.Context) {

	timePeriod := c.Query("period")
	salesReport, err := ad.adminUseCase.FilterSalesReport(timePeriod)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "sales report could not be retrieved", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "sales report retrieved successfully", salesReport, nil)
	c.JSON(http.StatusOK, successRes)

}

// SalesReportByDate godoc
// @Summary Get sales report by date
// @Description Retrieves the sales report for a specific date range
// @Tags Admin
// @Accept json
// @Produce json
// @Param start query string true "Start date"
// @Param end query string true "End date"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/sales-report-by-date [get]
func (ad *AdminHandler) SalesReportByDate(c *gin.Context) {
	startDateStr := c.Query("start")
	endDateStr := c.Query("end")
	if startDateStr == "" || endDateStr == "" {
		err := response.ClientResponse(http.StatusBadRequest, "Check the parameters", nil, "Empty date string")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	report, err := ad.adminUseCase.ExecuteSalesReportByDate(startDateStr, endDateStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to get sales report by date", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	success := response.ClientResponse(http.StatusOK, "the sales report by date", report, nil)
	c.JSON(http.StatusOK, success)
}
