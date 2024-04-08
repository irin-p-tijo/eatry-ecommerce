package interfaces

import (
	"eatry/pkg/domain"
	"eatry/pkg/utils/models"
)

type WalletUseCase interface {
	GetWallet(userID int) (domain.Wallet, error)
	WalletHistory(userID int) ([]models.WalletHistoryResp, error)
}
