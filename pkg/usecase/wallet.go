package usecase

import (
	"eatry/pkg/domain"
	interfaces "eatry/pkg/repository/interfaces"
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/models"
)

type WalletUseCase struct {
	walletRepository interfaces.WalletRepository
}

func NewWalletUseCase(walletrepository interfaces.WalletRepository) services.WalletUseCase {
	return &WalletUseCase{
		walletRepository: walletrepository,
	}
}
func (wa *WalletUseCase) GetWallet(userID int) (domain.Wallet, error) {
	wallet, err := wa.walletRepository.GetWallet(userID)
	if err != nil {
		return domain.Wallet{}, err
	}
	return wallet, nil
}
func (wa *WalletUseCase) WalletHistory(userID int) ([]models.WalletHistoryResp, error) {
	wallet, err := wa.walletRepository.GetWalletData(userID)
	if err != nil {
		return []models.WalletHistoryResp{}, err
	}
	walletResp, err := wa.walletRepository.GetWalletHistory(int(wallet.WalletID))
	if err != nil {
		return []models.WalletHistoryResp{}, err
	}

	return walletResp, nil
}
