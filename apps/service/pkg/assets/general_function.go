package assets

import (
	"fmt"
	"time"

	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/internal/models"
)

func Log(message string) {
	fmt.Println("------------------------------")
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
	fmt.Println("------------------------------")
}

func MakePeriod(req *odachin.RegisterRewardRequest, reward *models.Reward) *models.RewardPeriod {
	var reward_period *models.RewardPeriod
	now := time.Now()

	if req.RewardType == odachin.Reward_DAILY {
		tomorrowMidnight := time.Date(
			now.Year(),
			now.Month(),
			now.Day()+1,
			0, 0, 0, 0,
			now.Location(),
		)
		reward_period = &models.RewardPeriod{
			RewardID:    reward.RewardID,
			StartDate:   now,
			EndDate:     tomorrowMidnight,
			IsCompleted: false,
		}
	} else if req.RewardType == odachin.Reward_WEEKLY {
		weekday := int(now.Weekday())
		daysUntilNextMonday := (8 - weekday) % 7
		if daysUntilNextMonday == 0 {
			daysUntilNextMonday = 7
		}
		nextMonday := now.AddDate(0, 0, daysUntilNextMonday)
		nextMondayMidnight := time.Date(
			nextMonday.Year(),
			nextMonday.Month(),
			nextMonday.Day(),
			0, 0, 0, 0,
			now.Location(),
		)
		reward_period = &models.RewardPeriod{
			RewardID:    reward.RewardID,
			StartDate:   now,
			EndDate:     nextMondayMidnight,
			IsCompleted: false,
		}
	} else if req.RewardType == odachin.Reward_MONTHLY {
		firstOfNextMonth := time.Date(
			now.Year(),
			now.Month()+1, // time.Month は繰り上がりに対応している（12月+1 → 翌年1月）
			1,
			0, 0, 0, 0,
			now.Location(),
		)
		reward_period = &models.RewardPeriod{
			RewardID:    reward.RewardID,
			StartDate:   now,
			EndDate:     firstOfNextMonth,
			IsCompleted: false,
		}
	}
	return reward_period
}
