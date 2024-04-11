package handlers

import (
	"eatry/pkg/domain"
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryUseCase services.CategoryUseCase
}

func NewCategoryHandler(usecase services.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{
		CategoryUseCase: usecase,
	}
}

// @Summary Add a new category
// @Description Creates a new category in the system
// @Tags category
// @Accept json
// @Produce json
// @Param category body domain.Category true "Category Details"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /category [post]
func (cat *CategoryHandler) AddCategory(c *gin.Context) {
	var category domain.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the parameters are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	CategoryResponse, err := cat.CategoryUseCase.AddCategory(category)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the category cannot be added", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the category is added successfully", CategoryResponse, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Delete a category
// @Description Deletes a category from the system by its ID
// @Tags category
// @Accept json
// @Produce json
// @Param id query int true "Category ID"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /category/:id [delete]
func (cat *CategoryHandler) DeleteCategory(c *gin.Context) {
	CategoryID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check the parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err1 := cat.CategoryUseCase.DeleteCategory(CategoryID)

	if err1 != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the Category", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Get all categories
// @Description Retrieves a list of categories based on pagination parameters
// @Tags category
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param count query int true "Number of items per page"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /categories [get]
func (cat *CategoryHandler) GetCategory(c *gin.Context) {
	pagestr := c.Query("page")
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check the parameters", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check the parameters", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	categories, err := cat.CategoryUseCase.GetCategory(page, count)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrive the categories", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the categories are retrived", categories, nil)
	c.JSON(http.StatusOK, successRes)

}
