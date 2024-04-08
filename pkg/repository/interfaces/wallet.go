package interfaces

import (
	"eatry/pkg/domain"
	"eatry/pkg/utils/models"
)

type WalletRepository interface {
	GetWallet(userID int) (domain.Wallet, error)
	GetWalletData(userID int) (models.WalletHistory, error)
	AddtoWallet(userID int, amount float64) error
	AddToWalletHistory(wallet models.WalletHistory) error
	GetWalletHistory(walletID int) ([]models.WalletHistoryResp, error)
}
