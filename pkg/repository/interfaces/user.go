package interfaces

import "eatry/pkg/utils/models"

type UserRepository interface {
	CheckUserAvailability(email string) bool
	FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error)
	UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error)
	LoginHandler(user models.UserDetails) (models.UserDetailsResponse, error)
	AddAddress(userID int, address models.AddAddress) error
	DeleteAddress(userID, addressID int) error
	GetAddresses(id int) ([]models.Address, error)
	UserProfile(id int) (models.UserProfile, error)
	GetAllPaymentOptions() ([]models.PaymentDetails, error)
}
