package interfaces

import (
	"eatry/pkg/domain"
	"eatry/pkg/utils/models"
)

type CouponUseCase interface {
	AddCoupon(coupon models.AddCoupon) error
	MakeCouponInvalid(id int) error
	GetAllCoupons() ([]domain.Coupon, error)
}
