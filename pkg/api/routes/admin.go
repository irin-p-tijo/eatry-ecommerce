package routes

import (
	"eatry/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, adminHandler *handlers.AdminHandler, categoryHandler *handlers.CategoryHandler, productHandler *handlers.ProductHandler, couponHandler *handlers.CouponHandler, paymentHandlers *handlers.PaymentHandler, orderHandler *handlers.OrderHandler) {
	router.POST("/adminlogin", adminHandler.LoginHandler)
	router.POST("/adminsignup", adminHandler.CreateAdmin)
	router.GET("/dashboard", adminHandler.Dashboard)
	router.GET("/sales-report", adminHandler.FilterSalesReport)
	router.GET("/salesreportbydate", adminHandler.SalesReportByDate)

	userdetails := router.Group("/users")
	{
		userdetails.GET("/listofusers", adminHandler.GetUsers)
		userdetails.GET("/blockusers", adminHandler.BlockUser)
		userdetails.GET("/unblockusers", adminHandler.UnBlockUser)

	}

	category := router.Group("/category")
	{
		category.POST("/addcategory", categoryHandler.AddCategory)
		category.DELETE("/deletecategory", categoryHandler.DeleteCategory)
		category.GET("/listofcategories", categoryHandler.GetCategory)
	}
	product := router.Group("products")
	{
		product.POST("/addproduct", productHandler.AddProduct)
		product.DELETE("/deleteproduct", productHandler.DeleteProduct)
		product.PUT("/updateproduct", productHandler.UpdateProduct)
		product.GET("/listproducts", productHandler.ListProduct)
		product.GET("/filtercategory", productHandler.FilterCategory)
	}

	coupon := router.Group("coupon")
	{
		coupon.POST("/addcoupon", couponHandler.CreateNewCoupon)
		coupon.DELETE("/deletecoupon", couponHandler.MakeCouponInvalid)
		coupon.GET("/getallcoupon", couponHandler.GetAllCoupons)
	}
	order := router.Group("/orders")
	{
		order.GET("/orderdetails", orderHandler.GetAllOrderDetailsForAdmin)
		order.GET("/approveorder", orderHandler.ApproveOrder)

	}
	payment := router.Group("payment")
	{
		payment.POST("/add paymentmethod", paymentHandlers.AddPaymentMethods)
		payment.DELETE("/deletepaymentmethods", paymentHandlers.DeletePaymentMethods)
		payment.GET("/getallpaymentmethods", paymentHandlers.GetAllPaymentMethods)
	}

}
