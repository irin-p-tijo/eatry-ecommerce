package usecase

import (
	"eatry/pkg/domain"
	interfaces "eatry/pkg/repository/interfaces"
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/models"
)

type couponUseCase struct {
	repository     interfaces.CouponRepository
	cartRepository interfaces.CartRepository
}

func NewCouponUseCase(repo interfaces.CouponRepository, cartrepository interfaces.CartRepository) services.CouponUseCase {
	return &couponUseCase{
		repository:     repo,
		cartRepository: cartrepository,
	}
}
func (coup *couponUseCase) AddCoupon(coupon models.AddCoupon) error {
	if err := coup.repository.AddCoupon(coupon); err != nil {
		return err
	}

	return nil
}

func (coup *couponUseCase) MakeCouponInvalid(id int) error {
	if err := coup.repository.MakeCouponInvalid(id); err != nil {
		return err
	}

	return nil
}

func (Coup *couponUseCase) GetAllCoupons() ([]domain.Coupon, error) {

	coupons, err := Coup.repository.GetAllCoupons()
	if err != nil {
		return []domain.Coupon{}, err
	}
	return coupons, nil

}
