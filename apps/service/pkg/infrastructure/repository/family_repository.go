package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type FamilyRepository interface {
	Get(id uint) (models.Family, error)
	Save(param *models.Family) (*models.Family, error)
}

type FamilyRepositoryImpl struct {
	db *gorm.DB
}

func NewFamilyRepository(db *gorm.DB) FamilyRepository {
	return &FamilyRepositoryImpl{db}
}

func (r *FamilyRepositoryImpl) Get(id uint) (models.Family, error) {
	var family models.Family
	if err := r.db.Where("family_id = ?", id).First(&family).Error; err != nil {
		return models.Family{}, err
	}
	return family, nil
}

func (r *FamilyRepositoryImpl) Save(family *models.Family) (*models.Family, error) {
	if err := r.db.Create(&family).Error; err != nil {
		return nil, err
	}
	return family, nil
}
