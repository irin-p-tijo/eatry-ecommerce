package repository

import (
	"eatry/pkg/domain"
	interfaces "eatry/pkg/repository/interfaces"
	"eatry/pkg/utils/models"

	"gorm.io/gorm"
)

type WalletRepository struct {
	DB *gorm.DB
}

func NewWalletRepository(DB *gorm.DB) interfaces.WalletRepository {
	return &WalletRepository{
		DB: DB,
	}
}
func (wa *WalletRepository) GetWallet(userID int) (domain.Wallet, error) {
	err := wa.DB.Exec("select * from wallets where user_id=?", userID).Error
	if err != nil {

		return domain.Wallet{}, err
	}
	return domain.Wallet{}, nil
}
func (wa *WalletRepository) GetWalletData(userID int) (models.WalletHistory, error) {
	err := wa.DB.Exec("select id,user_id,wallet_amount from wallets where user_id=?", userID).Error
	if err != nil {
		return models.WalletHistory{}, err
	}
	return models.WalletHistory{}, nil
}
func (wa *WalletRepository) AddtoWallet(userID int, amount float64) error {

	err := wa.DB.Exec("update wallets  set  wallet_amount=wallet_amount+? where user_id=? returning wallet_amount", amount, userID).Error
	if err != nil {
		return err
	}
	return nil

}
func (wa *WalletRepository) AddToWalletHistory(wallet models.WalletHistory) error {

	query := `
	insert into wallet_histories
	(wallet_id,order_id,wallet_amount,status)  
	values (?,?,?,?)
	`
	err := wa.DB.Exec(query, wallet.WalletID, wallet.OrderID, wallet.WalletAmount, wallet.Status).Error
	if err != nil {
		return err
	}
	return nil
}
func (wa *WalletRepository) GetWalletHistory(walletID int) ([]models.WalletHistoryResp, error) {
	var wallet []models.WalletHistoryResp
	err := wa.DB.Raw("select * from wallet_histories where wallet_id = ? ", walletID).Scan(&wallet).Error
	if err != nil {
		return []models.WalletHistoryResp{}, err
	}
	return wallet, nil
}
