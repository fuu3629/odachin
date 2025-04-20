package assets

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Log(message string) {
	fmt.Println("------------------------------")
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
	fmt.Println("------------------------------")
}

func ProtoToMap(msg proto.Message) (map[string]interface{}, error) {
	jsonBytes, err := protojson.Marshal(msg)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, err
	}
	// Convert keys to snake_case
	result = ToSnakeCaseMap(result)
	return result, nil
}
func ToSnakeCaseMap(m map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m {
		snakeKey := strcase.ToSnake(k)
		result[snakeKey] = v
	}
	return result
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
			IsEditable:  true,
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
			IsEditable:  true,
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
			IsEditable:  true,
		}
	}
	return reward_period
}
