package interfaces

import "eatry/pkg/utils/models"

type CartRepository interface {
	QuantityofProductInCart(userID, productID int) (float64, error)
	AddToCart(userID int, productID int, Quantity float64, productprice float64) error
	TotalPriceIncrementInCart(userID int, productID int) (float64, error)
	GetTotalPrice(userID int) (models.CartTotal, error)
	DisplayCart(userID int) ([]models.Cart, error)
	UpdateCart(userID int, productID int, Quantity float64, TotalPrice float64) error
	ProductExists(userID int, productID int) (bool, error)
	GetQuantityAndProductDetails(userID int, productID int, cartdetails struct {
		Quantity   int
		TotalPrice float64
	}) (struct {
		Quantity   int
		TotalPrice float64
	}, error)
	RemoveProductFromCart(userID int, productID int) error
	UpdateCartDetails(cartdetails struct {
		Quantity   int
		TotalPrice float64
	}, userId int, productId int) error
	RemoveFromCart(userID int) ([]models.Cart, error)
	GetAllItemsFromCart(userID int) ([]models.Cart, error)
}
