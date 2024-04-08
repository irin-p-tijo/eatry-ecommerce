package models

type AddCoupon struct {
	Coupon             string  `json:"coupon" binding:"required"`
	DiscountPercentage int     `json:"discount_percentage" binding:"required"`
	MinimumPrice       float64 `json:"minimum_price" binding:"required"`
	Validity           bool    `json:"validity" binding:"required"`
}
type CouponAddUser struct {
	CouponName string `json:"coupon_name" binding:"required"`
}
type Coupons struct {
	Coupon       string `json:"coupon" gorm:"unique;not null"`
	DiscountRate int    `json:"discount_rate" gorm:"not null"`
	Valid        bool   `json:"valid" gorm:"default:true"`
}
