package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type WalletRepository interface {
	Get(tx *gorm.DB, id string) (models.Wallet, error)
	Save(tx *gorm.DB, param *models.Wallet) error
	Update(tx *gorm.DB, param *models.Wallet) error
	GetByUserId(tx *gorm.DB, userId string) (models.Wallet, error)
	GetByConditions(tx *gorm.DB, conditions string, args ...interface{}) ([]models.Wallet, error)
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

func (r *WalletRepositoryImpl) GetByUserId(tx *gorm.DB, userId string) (models.Wallet, error) {
	var wallet models.Wallet
	if err := tx.Where("user_id = ?", userId).First(&wallet).Error; err != nil {
		return models.Wallet{}, err
	}
	return wallet, nil
}

func (r *WalletRepositoryImpl) GetByConditions(tx *gorm.DB, conditions string, args ...interface{}) ([]models.Wallet, error) {
	var wallets []models.Wallet
	if err := tx.Where(conditions, args...).Find(&wallets).Error; err != nil {
		return nil, err
	}
	return wallets, nil
}
