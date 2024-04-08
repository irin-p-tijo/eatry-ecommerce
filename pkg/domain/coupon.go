package domain

type Coupon struct {
	ID                 int     `json:"id" gorm:"uniquekey; not null"`
	Coupon             string  `json:"coupon" gorm:"coupon"`
	DiscountPercentage int     `json:"discount_percentage"`
	Validity           bool    `json:"validity"`
	MinimumPrice       float64 `json:"minimum_price"`
}
type UsedCoupon struct {
	ID       int    `json:"id" gorm:"uniquekey not null"`
	CouponID int    `json:"coupon_id"`
	Coupon   Coupon `json:"-" gorm:"foreignkey:CouponID"`
	UserID   int    `json:"user_id"`
	Users    Users  `json:"-" gorm:"foreignkey:UserID"`
	Used     bool   `json:"used"`
}
