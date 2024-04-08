package domain

type Wallet struct {
	ID           int     `json:"id" gorm:"unique;not null"`
	UserID       int     `json:"user_id"`
	Users        Users   `json:"-" gorm:"foreignkey:UserID"`
	WalletAmount float64 `json:"wallet_amount"`
}
type WalletHistory struct {
	ID           int     `json:"id"  gorm:"unique;not null"`
	WalletID     int     `json:"wallet_id" gorm:"not null"`
	Wallet       Wallet  `json:"-" gorm:"foreignkey:WalletID"`
	OrderID      string  `json:"order_id" gorm:"not null"`
	WalletAmount float64 `json:"wallet_amount" gorm:"not null"`
	Status       string  `json:"status" gorm:"status:2;default:'CREDITED';check:status IN ('CREDITED','DEBITED')"`
}
type NewWalletHistory struct {
	ID           int     `json:"id"  gorm:"unique;not null"`
	WalletID     int     `json:"wallet_id" gorm:"not null"`
	Wallet       Wallet  `json:"-" gorm:"foreignkey:WalletID"`
	OrderID      string  `json:"order_id" gorm:"not null"`
	WalletAmount float64 `json:"wallet_amount" gorm:"not null"`
	Status       string  `json:"status" gorm:"status:2;default:'CREDITED';check:status IN ('CREDITED','DEBITED')"`
}
