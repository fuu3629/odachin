package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(tx *gorm.DB, transaction *models.Transaction) error
	SaveList(tx *gorm.DB, transactions []*models.Transaction) error
}

type TransactionRepositoryImpl struct{}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (r *TransactionRepositoryImpl) Save(tx *gorm.DB, transaction *models.Transaction) error {
	if err := tx.Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepositoryImpl) SaveList(tx *gorm.DB, transactions []*models.Transaction) error {
	if len(transactions) == 0 {
		return nil
	}
	if err := tx.Create(&transactions).Error; err != nil {
		return err
	}
	return nil
}
