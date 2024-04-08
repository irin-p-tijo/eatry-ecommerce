package repository

import (
	interfaces "eatry/pkg/repository/interfaces"
	"eatry/pkg/utils/models"
	"errors"

	"gorm.io/gorm"
)

type WishlistRepository struct {
	DB *gorm.DB
}

func NewWishlistRepository(DB *gorm.DB) interfaces.WishlistRepository {
	return &WishlistRepository{DB}
}
func (w *WishlistRepository) AddToWishList(userID int, productID int) error {

	err := w.DB.Exec("insert into wish_lists (user_id,product_id) values (?, ?)", userID, productID).Error
	if err != nil {
		return err
	}

	return nil
}

func (w *WishlistRepository) GetWishList(userID int) ([]models.WishListResponse, error) {

	var wishList []models.WishListResponse
	err := w.DB.Raw("select products.id as product_id, products.name as name,products.price as product_price from products inner join wish_lists on products.id = wish_lists.product_id where wish_lists.user_id = ? ", userID).Scan(&wishList).Error
	if err != nil {
		return []models.WishListResponse{}, err
	}

	return wishList, nil

}

func (w *WishlistRepository) RemoveFromWishList(userID int, productID int) error {

	err := w.DB.Exec("delete from wish_lists where user_id = ? and product_id = ?", userID, productID).Error
	if err != nil {
		return err
	}

	return nil

}
func (w *WishlistRepository) ProductExistInWishList(productID int, userID int) (bool, error) {

	var count int
	err := w.DB.Raw("select count(*) from wish_lists where user_id = ? and product_id = ? ", userID, productID).Scan(&count).Error
	if err != nil {
		return false, errors.New("error checking user product already present")
	}

	return count > 0, nil

}
func (w *WishlistRepository) DoesProductExist(productID int) (bool, error) {

	var count int
	err := w.DB.Raw("select count(*) from products where id = ?", productID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
