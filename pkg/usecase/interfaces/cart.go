package interfaces

import "eatry/pkg/utils/models"

type CartUseCase interface {
	AddToCart(userID int, productID int) (models.CartResponse, error)
	RemoveFromCart(userID int, productID int) (models.CartResponse, error)
	DisplayCart(userID int) (models.CartResponse, error)
}
