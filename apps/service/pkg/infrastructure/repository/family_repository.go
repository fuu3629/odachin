package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type FamilyRepository interface {
	Get(tx *gorm.DB, id uint) (models.Family, error)
	Save(tx *gorm.DB, param *models.Family) (*models.Family, error)
}

type FamilyRepositoryImpl struct {
}

func NewFamilyRepository() FamilyRepository {
	return &FamilyRepositoryImpl{}
}

func (r *FamilyRepositoryImpl) Get(tx *gorm.DB, id uint) (models.Family, error) {
	var family models.Family
	if err := tx.Where("family_id = ?", id).First(&family).Error; err != nil {
		return models.Family{}, err
	}
	return family, nil
}

func (r *FamilyRepositoryImpl) Save(tx *gorm.DB, family *models.Family) (*models.Family, error) {
	if err := tx.Create(&family).Error; err != nil {
		return nil, err
	}
	return family, nil
}
