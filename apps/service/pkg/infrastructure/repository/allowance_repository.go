package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type AllowanceRepository interface {
	Get(tx *gorm.DB, id uint) (models.Allowance, error)
	Save(tx *gorm.DB, param *models.Allowance) error
	Update(tx *gorm.DB, param map[string]interface{}) error
	GetAllowanceByCondition(tx *gorm.DB, condition string, args ...interface{}) ([]models.Allowance, error)
}

type AllowanceRepositoryImpl struct {
}

func NewAllowanceRepository() AllowanceRepository {
	return &AllowanceRepositoryImpl{}
}

func (r *AllowanceRepositoryImpl) Get(tx *gorm.DB, id uint) (models.Allowance, error) {
	var allowance models.Allowance
	if err := tx.Where("allowance_id = ?", id).First(&allowance).Error; err != nil {
		return models.Allowance{}, err
	}
	return allowance, nil
}

func (r *AllowanceRepositoryImpl) GetAllowanceByCondition(tx *gorm.DB, condition string, args ...interface{}) ([]models.Allowance, error) {
	var allowances []models.Allowance
	if err := tx.Where(condition, args...).Find(&allowances).Error; err != nil {
		return nil, err
	}
	return allowances, nil
}

func (r *AllowanceRepositoryImpl) Save(tx *gorm.DB, allowance *models.Allowance) error {
	if err := tx.Create(&allowance).Error; err != nil {
		return err
	}
	return nil
}

func (r *AllowanceRepositoryImpl) Update(tx *gorm.DB, allowance map[string]interface{}) error {
	id := allowance["allowance_id"]
	delete(allowance, "allowance_id")
	if err := tx.Model(&models.Allowance{}).Where("allowance_id = ?", id).Updates(&allowance).Error; err != nil {
		return err
	}
	return nil
}
