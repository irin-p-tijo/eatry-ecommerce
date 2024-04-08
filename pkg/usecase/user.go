package usecase

import (
	"eatry/pkg/helper"
	interfaces "eatry/pkg/repository/interfaces"
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/models"
	"errors"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepo        interfaces.UserRepository
	cartReopository interfaces.CartRepository
}

func NewUserUseCase(repo interfaces.UserRepository, cartrepo interfaces.CartRepository) services.UserUseCase {
	return &UserUseCase{
		userRepo:        repo,
		cartReopository: cartrepo,
	}
}
func (u *UserUseCase) UserSignUp(user models.UserDetails) (models.TokenUsers, error) {

	userExist := u.userRepo.CheckUserAvailability(user.Email)

	if userExist {
		return models.TokenUsers{}, errors.New("user already exist, sign in")
	}

	if user.Password != user.ConfirmPassword {
		return models.TokenUsers{}, errors.New("password does not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return models.TokenUsers{}, errors.New("internal server error")
	}

	user.Password = string(hashedPassword)

	userData, err := u.userRepo.UserSignUp(user)
	if err != nil {
		return models.TokenUsers{}, err
	}

	// crete a JWT token string for the user
	accessToken, err := helper.GenerateAccessToken(userData)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	refreshToken, err := helper.GenerateRefreshToken(userData)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	// copies all the details except the password of the user
	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &userData)
	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users:        userDetails,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (u *UserUseCase) LoginHandler(user models.UserLogin) (models.TokenUsers, error) {

	ok := u.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.TokenUsers{}, errors.New("the user does not exist")
	}

	// Get the user details in order to check the password, in this case ( The same function can be reused in future )
	user_details, err := u.userRepo.FindUserByEmail(user)
	if err != nil {
		return models.TokenUsers{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password), []byte(user.Password))
	if err != nil {
		return models.TokenUsers{}, errors.New("password incorrect")
	}

	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &user_details)
	if err != nil {
		return models.TokenUsers{}, err
	}

	// crete a JWT token string for the user
	accessToken, err := helper.GenerateAccessToken(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	refreshToken, err := helper.GenerateRefreshToken(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	return models.TokenUsers{
		Users:        userDetails,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
func (u *UserUseCase) AddAddress(userID int, address models.AddAddress) error {
	err := u.userRepo.AddAddress(userID, address)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) DeleteAddress(userID, addressID int) error {
	err := u.userRepo.DeleteAddress(userID, addressID)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserUseCase) GetAllAddress(id int) ([]models.Address, error) {
	address, err := u.userRepo.GetAddresses(id)
	if err != nil {
		return []models.Address{}, err
	}
	return address, nil
}
func (u *UserUseCase) UserProfile(id int) (models.UserProfile, error) {

	userprofile, err := u.userRepo.UserProfile(id)
	if err != nil {
		return models.UserProfile{}, err
	}
	return userprofile, nil
}
func (u *UserUseCase) CheckOut(userID int) (models.CheckoutDetails, error) {
	id := userID
	alluseraddress, err := u.userRepo.GetAddresses(id)
	if err != nil {
		return models.CheckoutDetails{}, err
	}
	paymentdetails, err := u.userRepo.GetAllPaymentOptions()
	if err != nil {
		return models.CheckoutDetails{}, err
	}
	Cart, err := u.cartReopository.GetAllItemsFromCart(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}
	GrandTotal, err := u.cartReopository.GetTotalPrice(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}
	return models.CheckoutDetails{
		Address:        alluseraddress,
		Payment_Method: paymentdetails,
		Cart:           Cart,
		Total_Price:    GrandTotal.FinalPrice,
	}, nil
}
