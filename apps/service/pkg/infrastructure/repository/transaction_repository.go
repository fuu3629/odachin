package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetByCondition(tx *gorm.DB, condition string, args ...interface{}) ([]*models.Transaction, error)
	Save(tx *gorm.DB, transaction *models.Transaction) error
	SaveList(tx *gorm.DB, transactions []*models.Transaction) error
}

type TransactionRepositoryImpl struct{}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (r *TransactionRepositoryImpl) GetByCondition(tx *gorm.DB, condition string, args ...interface{}) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	if err := tx.Where(condition, args...).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
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
