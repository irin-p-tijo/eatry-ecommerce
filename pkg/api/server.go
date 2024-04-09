package api

import (
	"eatry/pkg/api/handlers"
	"eatry/pkg/api/routes"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
	adminHandler *handlers.AdminHandler,
	userHandler *handlers.UserHandler,
	categoryHandler *handlers.CategoryHandler,
	productHandler *handlers.ProductHandler,
	cartHandler *handlers.CartHandler,
	couponHandler *handlers.CouponHandler,
	orderHandler *handlers.OrderHandler,
	paymentHandler *handlers.PaymentHandler,
	wishlistHandler *handlers.WishlistHandler,
	walletHandler *handlers.WalletHandler,
) *ServerHTTP {

	router := gin.New()

	router.LoadHTMLGlob("templates/*.html")

	router.Use(gin.Logger())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	/////////////////routes//////////////////////////
	routes.AdminRoutes(router.Group("/admin"), adminHandler, categoryHandler, productHandler, couponHandler, paymentHandler, orderHandler)
	routes.UserRoutes(router.Group("/user"), userHandler, cartHandler, orderHandler, paymentHandler, productHandler, wishlistHandler, walletHandler)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start(infoLog *log.Logger, errorLog *log.Logger) {

	infoLog.Printf("starting server on :8000")
	err := sh.engine.Run(":8000")
	errorLog.Fatal(err)
}
