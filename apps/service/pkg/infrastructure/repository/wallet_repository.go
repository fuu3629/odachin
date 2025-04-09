package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type WalletRepository interface {
	Get(tx *gorm.DB, id string) (models.Wallet, error)
	Save(tx *gorm.DB, param *models.Wallet) error
	Update(tx *gorm.DB, param *models.Wallet) error
}

type WalletRepositoryImpl struct {
}

func NewWalletRepository() WalletRepository {
	return &WalletRepositoryImpl{}
}

func (r *WalletRepositoryImpl) Get(tx *gorm.DB, id string) (models.Wallet, error) {
	var wallet models.Wallet
	if err := tx.Where("wallet_id = ?", id).First(&wallet).Error; err != nil {
		return models.Wallet{}, err
	}
	return wallet, nil
}

func (r *WalletRepositoryImpl) Save(tx *gorm.DB, wallet *models.Wallet) error {
	if err := tx.Create(&wallet).Error; err != nil {
		return err
	}
	return nil
}

func (r *WalletRepositoryImpl) Update(tx *gorm.DB, wallet *models.Wallet) error {
	if err := tx.Updates(&wallet).Error; err != nil {
		return err
	}
	return nil
}
