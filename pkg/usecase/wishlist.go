package usecase

import (
	interfaces "eatry/pkg/repository/interfaces"
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/models"
	"errors"
)

type wishlistUseCase struct {
	wishlistRepo interfaces.WishlistRepository
}

func NewWishlistUseCase(repo interfaces.WishlistRepository) services.WishlistUseCase {
	return &wishlistUseCase{
		wishlistRepo: repo,
	}
}

func (w *wishlistUseCase) AddToWishList(productID int, userID int) error {

	productExist, err := w.wishlistRepo.DoesProductExist(productID)
	if err != nil {
		return err
	}

	if !productExist {
		return errors.New("product does not exist")
	}

	productExistInWishList, err := w.wishlistRepo.ProductExistInWishList(productID, userID)
	if err != nil {
		return err
	}
	if productExistInWishList {
		return errors.New("product already exist in wishlist")
	}

	err = w.wishlistRepo.AddToWishList(userID, productID)
	if err != nil {
		return err
	}

	return nil
}
func (w *wishlistUseCase) GetWishList(userID int) ([]models.WishListResponse, error) {

	wishList, err := w.wishlistRepo.GetWishList(userID)
	if err != nil {
		return []models.WishListResponse{}, err
	}

	return wishList, err
}

func (w *wishlistUseCase) RemoveFromWishList(productID int, userID int) error {

	productExistInWishList, err := w.wishlistRepo.ProductExistInWishList(productID, userID)
	if err != nil {
		return err
	}
	if !productExistInWishList {
		return errors.New("product does not exist in wishlist")
	}

	err = w.wishlistRepo.RemoveFromWishList(userID, productID)
	if err != nil {
		return err
	}

	return nil
}
