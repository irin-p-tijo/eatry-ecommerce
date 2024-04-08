package models

type Orders struct {
	UserID          int `json:"user_id"`
	AddressID       int `json:"address_id"`
	PaymentMethodID int `json:"payment_method_id"`
	CouponID        int `json:"coupon_id"`
}
type OrderFromCart struct {
	PaymentID int `json:"payment_id" binding:"required"`
	AddressID int `json:"address_id" binding:"required"`
}
type OrderIncoming struct {
	UserID    int `json:"user_id"`
	PaymentID int `json:"payment_id"`
	AddressID int `json:"address_id"`
}
type CombinedOrderDetails struct {
	OrderId        string  `json:"order_id"`
	FinalPrice     float64 `json:"final_price"`
	ShipmentStatus string  `json:"shipment_status"`
	PaymentStatus  string  `json:"payment_status"`
	Name           string  `json:"firstname"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	HouseName      string  `json:"house_name" validate:"required"`
	State          string  `json:"state" validate:"required"`
	Pin            string  `json:"pin" validate:"required"`
	Street         string  `json:"street"`
	City           string  `json:"city"`
}
type OrderPaymentDetails struct {
	UserID     int     `json:"user_id"`
	Username   string  `json:"username"`
	Razor_id   string  `josn:"razor_id"`
	OrderID    int     `json:"order_id"`
	FinalPrice float64 `json:"final_price"`
}
type FullOrderDetails struct {
	OrderDetails        OrderDetails
	OrderProductDetails []OrderProductDetails
}
type OrderDetails struct {
	OrderId        string
	FinalPrice     float64
	ShipmentStatus string
	PaymentStatus  string
}
type OrderProductDetails struct {
	ProductID  int     `json:"product_id"`
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
type OrderProducts struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
type Invoice struct {
	OrderDetails OrderDetails
	AddressInfo  Address
	Cart         []Cart
}
