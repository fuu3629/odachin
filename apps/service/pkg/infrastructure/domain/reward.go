package domain

import (
	"time"

	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func RewardDomain(reward_list []models.Reward, tx *gorm.DB) error {
	if len(reward_list) == 0 {
		return nil
	}
	reward_period_repository := repository.NewRewardPeriodRepository()
	reward_period_list := []models.RewardPeriod{}
	now := time.Now()
	start := time.Date(
		now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)
	var end time.Time
	if reward_list[0].PeriodType == "DAILY" {
		end = start.AddDate(0, 0, 1)
	}
	if reward_list[0].PeriodType == "WEEKLY" {
		end = start.AddDate(0, 0, 7)
	}
	if reward_list[0].PeriodType == "MONTHLY" {
		end = start.AddDate(0, 1, 0)
	}
	for _, reward := range reward_list {
		reward_period := models.RewardPeriod{
			RewardID:  reward.RewardID,
			StartDate: start,
			EndDate:   end,
			Status:    "IN_PROGRESS",
		}
		reward_period_list = append(reward_period_list, reward_period)
	}
	err := reward_period_repository.SaveList(tx, reward_period_list)
	if err != nil {
		return status.Errorf(codes.Internal, "database error: %v", err)
	}
	return nil
}
