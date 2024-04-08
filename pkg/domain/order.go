package domain

import "time"

type Order struct {
	OrderId         string        `json:"order_id" gorm:"primaryKey;not null"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	DeletedAt       *time.Time    `json:"deleted_at"`
	UserID          int           `json:"user_id" gorm:"not null"`
	Users           Users         `json:"-" gorm:"foreignkey:UserID"`
	AddressID       int           `json:"address_id" gorm:"not null"`
	Address         Address       `json:"-" gorm:"foreignkey:AddressID"`
	PaymentMethodID int           `json:"paymentmethod_id"`
	PaymentMethod   PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
	FinalPrice      float64       `json:"final_price"`
	GrandTotal      float64       `json:"grand_total"`
	ShipmentStatus  string        `json:"status"`
	PaymentStatus   string        `json:"payment_status"`
	Approval        bool          `json:"approval"`
}

type UserOrderItem struct {
	ID         int      `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID    string   `json:"order_id"`
	Orders     Order    `json:"-" gorm:"foreignkey:OrderID;constraint:OnDelete:CASCADE"`
	ProductID  int      `json:"product_id"`
	Products   Products `json:"-" gorm:"foreignkey:ProductID"`
	Quantity   int      `json:"quantity"`
	TotalPrice float64  `json:"total_price"`
}
type PaymentMethod struct {
	ID           int    `json:"id" gorm:"primarykey"`
	Payment_Name string `json:"payment_name"`
}
type OrderSuccessResponse struct {
	OrderID        string `json:"order_id"`
	ShipmentStatus string `json:"order_status"`
}
