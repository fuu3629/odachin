package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type WalletRepository interface {
	Get(id string) (models.Wallet, error)
	Save(param *models.Wallet) error
	Update(param *models.Wallet) error
}

type WalletRepositoryImpl struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &WalletRepositoryImpl{db}
}

func (r *WalletRepositoryImpl) Get(id string) (models.Wallet, error) {
	var wallet models.Wallet
	if err := r.db.Where("wallet_id = ?", id).First(&wallet).Error; err != nil {
		return models.Wallet{}, err
	}
	return wallet, nil
}

func (r *WalletRepositoryImpl) Save(wallet *models.Wallet) error {
	if err := r.db.Create(&wallet).Error; err != nil {
		return err
	}
	return nil
}

func (r *WalletRepositoryImpl) Update(wallet *models.Wallet) error {
	if err := r.db.Updates(&wallet).Error; err != nil {
		return err
	}
	return nil
}
