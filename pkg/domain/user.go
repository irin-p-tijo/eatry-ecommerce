package domain

type Users struct {
	ID       int    `json:"id" gorm:"primary key,not null"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8,max=20"`
	Phone    string `json:"phone"`
	Blocked  bool   `json:"blocked" gorm:"default:false"`
}
type Address struct {
	ID        int    `json:"id" gorm:"primary key;not null"`
	UserID    int    `json:"user_id"`
	Users     Users  `json:"-" gorm:"foreignkey:UserID"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	District  string `json:"district" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin_code" validate:"required"`
}
