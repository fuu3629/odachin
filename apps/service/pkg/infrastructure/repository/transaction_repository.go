package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(tx *gorm.DB, transaction *models.Transaction) error
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
