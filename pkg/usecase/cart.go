package usecase

import (
	interfaces "eatry/pkg/repository/interfaces"
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/models"
	"errors"
)

type CartUseCase struct {
	cartRepository    interfaces.CartRepository
	productRepository interfaces.ProductRepository
}

func NewCartUseCase(cartRepository interfaces.CartRepository, productrepository interfaces.ProductRepository) services.CartUseCase {
	return &CartUseCase{
		cartRepository:    cartRepository,
		productRepository: productrepository,
	}
}

func (car *CartUseCase) AddToCart(userID int, productID int) (models.CartResponse, error) {
	product, err := car.productRepository.CheckProduct(productID)
	if err != nil {
		return models.CartResponse{}, err
	}
	if !product {
		return models.CartResponse{}, errors.New("the product is exists in the cart")
	}

	stock, err := car.productRepository.CheckStock(productID)
	if err != nil {
		return models.CartResponse{}, err
	}
	if stock == 0 {
		return models.CartResponse{}, err
	}
	productInCart, err := car.cartRepository.QuantityofProductInCart(userID, productID)
	if err != nil {
		return models.CartResponse{}, err
	}
	productprice, err := car.productRepository.GetPriceofProduct(productID)
	if err != nil {
		return models.CartResponse{}, err
	}
	if productInCart == 0 {
		err := car.cartRepository.AddToCart(userID, productID, 1, productprice)
		if err != nil {
			return models.CartResponse{}, err
		}
	} else {
		TotalPrice, err := car.cartRepository.TotalPriceIncrementInCart(userID, productID)
		if err != nil {
			return models.CartResponse{}, err
		}
		err = car.cartRepository.UpdateCart(userID, productID, productInCart+1, TotalPrice+productprice)
		if err != nil {
			return models.CartResponse{}, err
		}
	}
	cartDetails, err := car.cartRepository.DisplayCart(userID)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartTotal, err := car.cartRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartResponse{}, err
	}
	cartResponse := models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       cartDetails,
	}

	return cartResponse, err
}
func (rt *CartUseCase) RemoveFromCart(userID int, productID int) (models.CartResponse, error) {
	ok, err := rt.cartRepository.ProductExists(userID, productID)
	if err != nil {
		return models.CartResponse{}, err
	}
	if !ok {
		return models.CartResponse{}, errors.New("the product does not exists in the cart")
	}
	var cartdetails struct {
		Quantity   int
		TotalPrice float64
	}
	cartdetails, err = rt.cartRepository.GetQuantityAndProductDetails(userID, productID, cartdetails)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartdetails.Quantity = cartdetails.Quantity - 1
	if cartdetails.Quantity == 0 {
		if err := rt.cartRepository.RemoveProductFromCart(userID, productID); err != nil {
			return models.CartResponse{}, err
		}
	}
	if cartdetails.Quantity != 0 {

		ProductPrice, err := rt.productRepository.GetPriceofProduct(productID)
		if err != nil {
			return models.CartResponse{}, err
		}
		cartdetails.TotalPrice = cartdetails.TotalPrice - ProductPrice
		err = rt.cartRepository.UpdateCartDetails(cartdetails, userID, productID)
		if err != nil {
			return models.CartResponse{}, err
		}
	}
	updatedCart, err := rt.cartRepository.RemoveFromCart(userID)
	if err != nil {
		return models.CartResponse{}, err
	}

	cartTotal, err := rt.cartRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartResponse{}, err
	}
	updatecart := models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       updatedCart,
	}
	return updatecart, nil
}
func (rt *CartUseCase) DisplayCart(userID int) (models.CartResponse, error) {
	displayCart, err := rt.cartRepository.DisplayCart(userID)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartTotal, err := rt.cartRepository.GetTotalPrice(userID)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartresponse := models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       displayCart,
	}
	return cartresponse, nil
}
