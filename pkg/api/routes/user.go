package routes

import (
	"eatry/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handlers.UserHandler, cartHandler *handlers.CartHandler, orderHandler *handlers.OrderHandler, paymentHandler *handlers.PaymentHandler, productHandler *handlers.ProductHandler, wishlistHandler *handlers.WishlistHandler, walletHandler *handlers.WalletHandler) {

	router.POST("/usersignup", userHandler.UserSignUp)
	router.POST("/userlogin", userHandler.LoginHandler)

	address := router.Group("/addresses")
	{
		address.POST("/addaddress", userHandler.AddAddress)
		address.DELETE("/deleteaddress", userHandler.DeleteAddress)
		address.GET("/alladdress", userHandler.GetAllAddress)
		address.GET("/userprofile", userHandler.UserProfile)
	}

	product := router.Group("products")
	{
		product.GET("/filtercategory", productHandler.FilterCategory)
		product.GET("/searchproduct", productHandler.SearchProduct)
	}
	cart := router.Group("/cart")
	{
		cart.POST("/addtocart", cartHandler.AddToCart)
		cart.DELETE("/removefromcart", cartHandler.RemoveFromCart)
		cart.GET("/displaycart", cartHandler.DisplayCart)

	}
	order := router.Group("/order")
	{
		order.POST("/orderfromcart", orderHandler.OrderItemsFromCart)
		order.DELETE("/cancelorders", orderHandler.CancelOrder)
		order.GET("/getorders", orderHandler.GetOrderDetails)
		order.GET("/generate", orderHandler.GenerateInvoice)

	}
	wallet := router.Group("/wallet")
	{
		wallet.GET("/getwallet", walletHandler.GetWallet)
		wallet.GET("/wallethistory", walletHandler.WalletHistory)
	}

	wishlist := router.Group("wishlist")
	{
		wishlist.POST("/wishlist", wishlistHandler.AddToWishList)
		wishlist.DELETE("/removewishlist", wishlistHandler.RemoveFromWishList)
		wishlist.GET("/getwishlist", wishlistHandler.GetWishList)
	}
	router.GET("/checkout", userHandler.CheckOut)
	router.GET("/razorpay", paymentHandler.MakePaymentRazorPay)
	router.GET("/verifypayment", paymentHandler.VerifyPayment)

}
