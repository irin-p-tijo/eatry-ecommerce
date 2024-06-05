package repository

import (
	"eatry/pkg/helper"
	interfaces "eatry/pkg/repository/interfaces"
	"eatry/pkg/utils/models"
	"errors"

	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.CartRepository {
	return &cartRepository{
		DB: db,
	}
}
func (car *cartRepository) QuantityofProductInCart(userID, productID int) (float64, error) {
	var qty float64
	err := car.DB.Raw("select quantity from carts where user_id = ? and product_id = ?", userID, productID).Scan(&qty).Error

	if err != nil {
		return 0.0, err
	}
	return qty, nil
}
func (car *cartRepository) AddToCart(userID int, productID int, Quantity float64, productprice float64) error {
	if err := car.DB.Exec("insert into carts(user_id,product_id,quantity,total_price) values (?,?,?,?)", userID, productID, Quantity, productprice).Error; err != nil {
		return err
	}
	return nil
}
func (car *cartRepository) TotalPriceIncrementInCart(userID int, productID int) (float64, error) {
	var TotalPrice float64
	if err := car.DB.Raw("select sum(total_price) as total_price from carts where user_id=? and product_id=?", userID, productID).Scan(&TotalPrice).Error; err != nil {
		return 0.0, err
	}
	return TotalPrice, nil

}
func (car *cartRepository) DisplayCart(userID int) ([]models.Cart, error) {
	var count int
	if err := car.DB.Raw("select count(*) from carts where user_id = ? ", userID).First(&count).Error; err != nil {
		return []models.Cart{}, err
	}

	if count == 0 {
		return []models.Cart{}, nil
	}
	var cartResponse []models.Cart
	if err := car.DB.Raw("select carts.user_id,users.name as user_name,carts.product_id,products.name as name,carts.quantity,carts.total_price from carts join users on carts.user_id=users.id join products on carts.product_id=products.id where user_id=?", userID).First(&cartResponse).Error; err != nil {
		return []models.Cart{}, err
	}
	return cartResponse, nil
}
func (car *cartRepository) GetTotalPrice(userID int) (models.CartTotal, error) {
	var cartTotal models.CartTotal
	err := car.DB.Raw("select COALESCE(SUM(total_price), 0) from carts where user_id = ?", userID).Scan(&cartTotal.TotalPrice).Error
	if err != nil {
		return models.CartTotal{}, err
	}
	err = car.DB.Raw("select name as user_name from users where id = ?", userID).Scan(&cartTotal.UserName).Error
	if err != nil {
		return models.CartTotal{}, err
	}
	var discount_price float64
	discount_price, err = helper.GetCouponDiscountPrice(userID, cartTotal.TotalPrice, car.DB)
	if err != nil {
		return models.CartTotal{}, err
	}

	cartTotal.FinalPrice = cartTotal.TotalPrice - discount_price
	return cartTotal, nil
}
func (car *cartRepository) UpdateCart(userID int, productID int, Quantity float64, TotalPrice float64) error {

	if err := car.DB.Exec("update carts set quantity=?,total_price=? where user_id=? and product_id=?", Quantity, TotalPrice, userID, productID).Error; err != nil {

		return err
	}
	return nil
}
func (car *cartRepository) ProductExists(userID int, productID int) (bool, error) {
	var count int
	if err := car.DB.Raw("select count(*)from carts where user_id = ? and product_id=?", userID, productID).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
func (car *cartRepository) GetQuantityAndProductDetails(userID int, productID int, cartdetails interfaces.CartDetails) (interfaces.CartDetails, error) {
	if err := car.DB.Raw("select quantity,total_price from carts where user_id = ? and product_id = ?", userID, productID).Scan(&cartdetails).Error; err != nil {
		return struct {
			Quantity   int
			TotalPrice float64
		}{}, err
	}
	return cartdetails, nil
}
func (car *cartRepository) RemoveProductFromCart(userID int, productID int) error {

	if err := car.DB.Exec("delete from carts where user_id = ? and product_id = ?", userID, productID).Error; err != nil {
		return err
	}

	return nil
}
func (car *cartRepository) UpdateCartDetails(cartdetails interfaces.CartDetails, userId int, productId int) error {
	if err := car.DB.Raw("update carts set quantity = ? , total_price = ? where user_id = ? and product_id = ? ", cartdetails.Quantity, cartdetails.TotalPrice, userId, productId).Scan(&cartdetails).Error; err != nil {
		return err
	}
	return nil

}
func (cr *cartRepository) RemoveFromCart(userID int) ([]models.Cart, error) {

	var cartResponse []models.Cart
	if err := cr.DB.Raw("select carts.product_id,products.name as name,carts.quantity,carts.total_price from carts inner join products on carts.product_id = products.id where carts.user_id = ?", userID).First(&cartResponse).Error; err != nil {
		return []models.Cart{}, err
	}

	return cartResponse, nil

}
func (cr *cartRepository) GetAllItemsFromCart(userID int) ([]models.Cart, error) {
	var count int
	var cartResponse []models.Cart

	err := cr.DB.Raw("select count(*) from carts where user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return []models.Cart{}, err
	}

	if count == 0 {
		return []models.Cart{}, nil
	}
	err = cr.DB.Raw("select carts.user_id,users.name as user_name,carts.product_id,products.name as name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userID).First(&cartResponse).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if len(cartResponse) == 0 {
				return []models.Cart{}, nil
			}
			return []models.Cart{}, err
		}
		return []models.Cart{}, err
	}

	return cartResponse, nil
}
