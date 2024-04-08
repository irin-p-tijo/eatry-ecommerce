package domain

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID     int      `json:"user_id" gorm:"uniquekey; not null"`
	Users      Users    `json:"-" gorm:"foreignkey:UserID"`
	ProductID  int      `json:"product_id"`
	Products   Products `json:"-" gorm:"foreignkey:ProductID"`
	Quantity   float64  `json:"quantity"`
	TotalPrice float64  `json:"total_price"`
}
