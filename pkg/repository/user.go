package repository

import (
	interfaces "eatry/pkg/repository/interfaces"
	"eatry/pkg/utils/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &UserRepository{DB}
}
func (c *UserRepository) CheckUserAvailability(email string) bool {

	var count int
	query := fmt.Sprintf("select count(*) from users where email='%s'", email)
	if err := c.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0

}
func (c *UserRepository) FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error) {

	var userDetails models.UserSignInResponse

	err := c.DB.Raw(`
		select * from users where email = ? and blocked = false	`, user.Email).Scan(&userDetails).Error

	if err != nil {
		return models.UserSignInResponse{}, errors.New("error checking user details")
	}

	return userDetails, nil

}
func (c *UserRepository) UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error) {

	var userDetails models.UserDetailsResponse
	err := c.DB.Raw(`INSERT INTO users (name, email, phone, password) VALUES ($1, $2, $3, $4) RETURNING id, name, email, phone`, user.Name, user.Email, user.Phone, user.Password).Scan(&userDetails).Error

	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return userDetails, nil
}
func (c *UserRepository) LoginHandler(user models.UserDetails) (models.UserDetailsResponse, error) {

	var userResponse models.UserDetailsResponse
	err := c.DB.Save(&userResponse).Error
	return userResponse, err

}
func (c *UserRepository) AddAddress(userID int, address models.AddAddress) error {
	err := c.DB.Exec("insert into addresses (user_id,name,house_name,street,city,district,state,pin) values (?,?,?,?,?,?,?,?)", address.UserID, address.Name, address.HouseName, address.Street, address.City, address.District, address.State, address.Pin).Error
	if err != nil {
		return err
	}
	return nil

}
func (c *UserRepository) DeleteAddress(userID, addressID int) error {

	err := c.DB.Exec("delete from addresses where user_id =? and id =? ", userID, addressID)
	if err.RowsAffected < 1 {
		return errors.New("the  address is not present")
	}
	return nil

}

func (c *UserRepository) GetAddresses(id int) ([]models.Address, error) {

	var addresses []models.Address

	if err := c.DB.Raw("SELECT * FROM addresses WHERE user_id=$1", id).Scan(&addresses).Error; err != nil {
		return []models.Address{}, err
	}

	return addresses, nil

}
func (c *UserRepository) UserProfile(id int) (models.UserProfile, error) {
	var userprofile models.UserProfile
	err := c.DB.Raw("select users.id,users.name,users.email,users.phone,addresses.house_name,addresses.street,addresses.city,addresses.district,addresses.state,addresses.pin from users  join addresses  on users.id=addresses.user_id where users.id = $1 ", id).Scan(&userprofile).Error
	if err != nil {
		return models.UserProfile{}, err
	}
	return userprofile, nil
}
func (cr *UserRepository) GetAllPaymentOptions() ([]models.PaymentDetails, error) {

	var paymentMethods []models.PaymentDetails
	err := cr.DB.Raw("select * from payment_methods").Scan(&paymentMethods).Error
	if err != nil {
		return []models.PaymentDetails{}, err
	}

	return paymentMethods, nil

}
