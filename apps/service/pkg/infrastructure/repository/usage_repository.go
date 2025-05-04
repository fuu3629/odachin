package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type UsageRepository interface {
	Save(tx *gorm.DB, usage *models.Usage) error
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
