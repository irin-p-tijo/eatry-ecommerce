package interfaces

import (
	"eatry/pkg/domain"
	"eatry/pkg/utils/models"
)

type CouponRepository interface {
	AddCoupon(coup models.AddCoupon) error
	MakeCouponInvalid(id int) error
	GetAllCoupons() ([]domain.Coupon, error)
	FindCouponDetails(couponID int) (domain.Coupon, error)
	CouponExist(couponName string) (bool, error)
	GetCouponMinimumAmount(coupon string) (float64, error)
	DidUserAlreadyUsedThisCoupon(coupon string, userID int) (bool, error)
	UpdateUsedCoupon(coupon string, userID int) (bool, error)
}
