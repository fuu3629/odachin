package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type UsageRepository interface {
	Save(tx *gorm.DB, usage *models.Usage) error
	Get(tx *gorm.DB, userID string) ([]*models.Usage, error)
}

type UsageRepositoryImpl struct{}

func NewUsageRepository() UsageRepository {
	return &UsageRepositoryImpl{}
}

func (r *UsageRepositoryImpl) Save(tx *gorm.DB, usage *models.Usage) error {
	if err := tx.Create(&usage).Error; err != nil {
		return err
	}
	return nil
}

func (r *UsageRepositoryImpl) Get(tx *gorm.DB, userID string) ([]*models.Usage, error) {
	var usages []*models.Usage
	if err := tx.Where("user_id = ?", userID).Find(&usages).Error; err != nil {
		return nil, err
	}
	return usages, nil
}
