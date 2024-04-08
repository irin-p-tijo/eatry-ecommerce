package models

type Wallet struct {
	Amount float64 `json:"wallet_amount"`
}

type WalletHistory struct {
	WalletID     int     `json:"wallet_id"  gorm:"not null"`
	OrderID      string  `json:"order_id" gorm:"not null"`
	WalletAmount float64 `json:"wallet_amount" gorm:"not null"`
	Status       string  `json:"status" gorm:"not null"`
}
type WalletHistoryResp struct {
	OrderID int     `json:"order_id" gorm:"not null"`
	Amount  float64 `json:"amount" gorm:"not null"`
	Status  string  `json:"status" gorm:"not null"`
}
