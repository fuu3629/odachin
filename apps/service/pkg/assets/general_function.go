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
			RewardID:  reward.RewardID,
			StartDate: now,
			EndDate:   tomorrowMidnight,
			Status:    "IN_PROGRESS",
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
			RewardID:  reward.RewardID,
			StartDate: now,
			EndDate:   nextMondayMidnight,
			Status:    "IN_PROGRESS",
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
			RewardID:  reward.RewardID,
			StartDate: now,
			EndDate:   firstOfNextMonth,
			Status:    "IN_PROGRESS",
		}
	}
	return reward_period
}

func Map[T any, R any](input []T, f func(T) R) []R {
	result := make([]R, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

var weekdayToInt = map[string]int{
	"SUNDAY":    0,
	"MONDAY":    1,
	"TUESDAY":   2,
	"WEDNESDAY": 3,
	"THURSDAY":  4,
	"FRIDAY":    5,
	"SATURDAY":  6,
}

var IntToWeekday = map[int]string{
	0: "SUNDAY",
	1: "MONDAY",
	2: "TUESDAY",
	3: "WEDNESDAY",
	4: "THURSDAY",
	5: "FRIDAY",
	6: "SATURDAY",
}
