package interfaces

import (
	"eatry/pkg/utils/models"
)

type UserUseCase interface {
	UserSignUp(user models.UserDetails) (models.TokenUsers, error)
	LoginHandler(user models.UserLogin) (models.TokenUsers, error)
	AddAddress(userID int, address models.AddAddress) error
	DeleteAddress(userID, addressID int) error
	GetAllAddress(id int) ([]models.Address, error)
	UserProfile(id int) (models.UserProfile, error)
	CheckOut(userID int) (models.CheckoutDetails, error)
}
