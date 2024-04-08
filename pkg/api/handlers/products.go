package handlers

import (
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/models"
	"eatry/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUseCase services.ProductUseCase
}

func NewProductHandler(usecase services.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: usecase,
	}
}
func (p *ProductHandler) AddProduct(c *gin.Context) {
	var addproduct models.AddProduct
	if err := c.ShouldBindJSON(&addproduct); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the constraints are given wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	products, err := p.productUseCase.AddProduct(addproduct)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the products cannot be added", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the product added successfully", products, nil)
	c.JSON(http.StatusOK, successRes)
}
func (p *ProductHandler) DeleteProduct(c *gin.Context) {
	ProductID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check the parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err1 := p.productUseCase.DeleteProduct(ProductID)
	if err1 != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the product cannot be deleted", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the product is deleted", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check the parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var Products models.ProductResponse
	if err := c.ShouldBindJSON(&Products); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the constraints are given wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	updatedproduct, err := p.productUseCase.UpdateProduct(Products, productID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "the product cannot be updated", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "the product is updated", updatedproduct, nil)
	c.JSON(http.StatusOK, successRes)

}
func (p *ProductHandler) ListProduct(c *gin.Context) {
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

	products, err := p.productUseCase.ListProducts(page, pagesize)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrive the data", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "the products data is retrived", products, nil)
	c.JSON(http.StatusOK, successRes)
}

func (p *ProductHandler) FilterCategory(c *gin.Context) {
	var data map[string]int
	if err := c.ShouldBindJSON(&data); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	productCategory, err := p.productUseCase.FilterCategory(data)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve products by category", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully filtered the category", productCategory, nil)
	c.JSON(http.StatusOK, successRes)
}
func (p *ProductHandler) SearchProduct(c *gin.Context) {

	var prefix models.SearchItems

	if err := c.ShouldBindJSON(&prefix); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	productDetails, err := p.productUseCase.SearchItemBasedOnPrefix(prefix.Name)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve products by prefix search", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Here is the product", productDetails, nil)
	c.JSON(http.StatusOK, successRes)

}
