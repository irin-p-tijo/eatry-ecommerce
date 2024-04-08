package repository

import (
	"eatry/pkg/domain"
	interfaces "eatry/pkg/repository/interfaces"
	"eatry/pkg/utils/models"

	"gorm.io/gorm"
)

type couponRepository struct {
	DB *gorm.DB
}

func NewCouponRepository(db *gorm.DB) interfaces.CouponRepository {
	return &couponRepository{
		DB: db,
	}
}
func (repo *couponRepository) AddCoupon(coup models.AddCoupon) error {
	if err := repo.DB.Exec("insert into coupons (coupon,discount_percentage,minimum_price,validity ) values (?,?,?,?)", coup.Coupon, coup.DiscountPercentage, coup.MinimumPrice, true).Error; err != nil {
		return err
	}

	return nil
}

func (repo *couponRepository) MakeCouponInvalid(id int) error {
	if err := repo.DB.Exec("update coupons set validity=false where id=$1", id).Error; err != nil {
		return err
	}

	return nil
}
func (c *couponRepository) GetAllCoupons() ([]domain.Coupon, error) {
	var model []domain.Coupon
	err := c.DB.Raw("select * from coupons").Scan(&model).Error
	if err != nil {
		return []domain.Coupon{}, err
	}

	return model, nil
}
func (repo *couponRepository) FindCouponDetails(couponID int) (domain.Coupon, error) {
	var coupon domain.Coupon
	err := repo.DB.Raw("select * from coupons where id=$1", couponID).Scan(&coupon).Error
	if err != nil {
		return domain.Coupon{}, err
	}

	return domain.Coupon{}, nil
}
func (repo *couponRepository) CouponExist(couponName string) (bool, error) {

	var count int
	err := repo.DB.Raw("select count(*) from coupons where coupon = ?", couponName).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil

}
func (repo *couponRepository) GetCouponMinimumAmount(coupon string) (float64, error) {

	var MinDiscountPrice float64
	err := repo.DB.Raw("select minimum_price from coupons where coupon = ?", coupon).Scan(&MinDiscountPrice).Error
	if err != nil {
		return 0.0, err
	}
	return MinDiscountPrice, nil
}
func (repo *couponRepository) DidUserAlreadyUsedThisCoupon(coupon string, userID int) (bool, error) {

	var count int
	err := repo.DB.Raw("select count(*) from used_coupons where coupon_id = (select id from coupons where coupon = ?) and user_id = ?", coupon, userID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil

}
func (repo *couponRepository) UpdateUsedCoupon(coupon string, userID int) (bool, error) {

	var couponID uint
	err := repo.DB.Raw("select id from coupons where coupon = ?", coupon).Scan(&couponID).Error
	if err != nil {
		return false, err
	}

	var count int
	// if a coupon have already been added, replace the order with current coupon and delete the existing coupon
	err = repo.DB.Raw("select count(*) from used_coupons where user_id = ? and used = false", userID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		err = repo.DB.Exec("delete from used_coupons where user_id = ? and used = false", userID).Error
		if err != nil {
			return false, err
		}
	}

	err = repo.DB.Exec("insert into used_coupons (coupon_id,user_id,used) values (?, ?, false)", couponID, userID).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
