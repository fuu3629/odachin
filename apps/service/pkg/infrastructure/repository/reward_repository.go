package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type RewardRepository interface {
	Get(tx *gorm.DB, id uint) (models.Reward, error)
	GetWithPeriodByUserId(tx *gorm.DB, id string) ([]models.Reward, error)
	Save(tx *gorm.DB, param *models.Reward) error
	Update(tx *gorm.DB, param *models.Reward) error
	Delete(tx *gorm.DB, id uint) error
	GetByToUserId(tx *gorm.DB, id string) ([]models.Reward, error)
}

type RewardRepositoryImpl struct {
}

func NewRewardRepository() RewardRepository {
	return &RewardRepositoryImpl{}
}

func (r *RewardRepositoryImpl) Get(tx *gorm.DB, id uint) (models.Reward, error) {
	var reward models.Reward
	if err := tx.Where("reward_id = ?", id).First(&reward).Error; err != nil {
		return models.Reward{}, err
	}
	return reward, nil
}

func (r *RewardRepositoryImpl) GetWithPeriodByUserId(tx *gorm.DB, id string) ([]models.Reward, error) {
	var reward []models.Reward
	if err := tx.Preload("RewardPeriods").Where("to_user_id = ?", id).Find(&reward).Error; err != nil {
		return []models.Reward{}, err
	}
	return reward, nil
}

func (r *RewardRepositoryImpl) Save(tx *gorm.DB, reward *models.Reward) error {
	if err := tx.Create(&reward).Error; err != nil {
		return err
	}
	return nil
}

func (r *RewardRepositoryImpl) Update(tx *gorm.DB, reward *models.Reward) error {
	if err := tx.Updates(&reward).Error; err != nil {
		return err
	}
	return nil
}

func (r *RewardRepositoryImpl) GetByToUserId(tx *gorm.DB, id string) ([]models.Reward, error) {
	var Reward []models.Reward
	if err := tx.Where("to_user_id = ?", id).Find(&Reward).Error; err != nil {
		return []models.Reward{}, err
	}
	return Reward, nil
}

func (r *RewardRepositoryImpl) Delete(tx *gorm.DB, id uint) error {
	var reward models.Reward
	if err := tx.Where("reward_id = ?", id).First(&reward).Error; err != nil {
		return err
	}
	if err := tx.Delete(&reward).Error; err != nil {
		return err
	}
	return nil
}
