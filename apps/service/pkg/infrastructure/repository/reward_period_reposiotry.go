package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type RewardPeriodRepository interface {
	Get(tx *gorm.DB, id string) (models.RewardPeriod, error)
	Save(tx *gorm.DB, param *models.RewardPeriod) error
	Update(tx *gorm.DB, reward_period map[string]interface{}) error
}

type RewardPeriodRepositoryImpl struct{}

func NewRewardPeriodRepository() RewardPeriodRepository {
	return &RewardPeriodRepositoryImpl{}
}

func (r *RewardPeriodRepositoryImpl) Get(tx *gorm.DB, id string) (models.RewardPeriod, error) {
	var reward_period models.RewardPeriod
	if err := tx.Where("reward_period_id = ?", id).First(&reward_period).Error; err != nil {
		return models.RewardPeriod{}, err
	}
	return reward_period, nil
}

func (r *RewardPeriodRepositoryImpl) Save(tx *gorm.DB, reward_period *models.RewardPeriod) error {
	if err := tx.Create(&reward_period).Error; err != nil {
		return err
	}
	return nil
}

func (r *RewardPeriodRepositoryImpl) Update(tx *gorm.DB, reward_period map[string]interface{}) error {
	id := reward_period["reward_period_id"]
	delete(reward_period, "reward_period_id")
	if err := tx.Model(&models.RewardPeriod{}).Where("reward_period_id = ?", id).Updates(&reward_period).Error; err != nil {
		return err
	}
	return nil
}
