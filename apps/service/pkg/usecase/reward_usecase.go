package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/fuu3629/odachin/apps/service/pkg/assets"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type RewardUsecase interface {
	RegisterReward(ctx context.Context, req *odachin.RegisterRewardRequest) error
	DeleteReward(ctx context.Context, req *odachin.DeleteRewardRequest) error
	GetRewardList(ctx context.Context, req *odachin.GetRewardListRequest) ([]models.RewardPeriod, error)
	GetUncompletedRewardCount(ctx context.Context) (*odachin.GetUncompletedRewardCountResponse, error)
}

type RewardUsecaseImpl struct {
	db                     *gorm.DB
	rewardRepository       repository.RewardRepository
	rewardPeriodRepository repository.RewardPeriodRepository
}

func NewRewardUsecase(db *gorm.DB) RewardUsecase {
	return &RewardUsecaseImpl{
		db:                     db,
		rewardRepository:       repository.NewRewardRepository(),
		rewardPeriodRepository: repository.NewRewardPeriodRepository(),
	}
}

func (u *RewardUsecaseImpl) RegisterReward(ctx context.Context, req *odachin.RegisterRewardRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		fmt.Println("user_id: ", user_id)
		reward := &models.Reward{
			FromUserID:  user_id,
			ToUserID:    req.ToUserId,
			PeriodType:  req.RewardType.String(),
			Title:       req.Title,
			Description: req.Description,
			Amount:      float64(req.Amount),
		}
		err := u.rewardRepository.Save(tx, reward)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		reward_period := assets.MakePeriod(req, reward)
		err = u.rewardPeriodRepository.Save(tx, reward_period)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	return nil
}

func (u *RewardUsecaseImpl) DeleteReward(ctx context.Context, req *odachin.DeleteRewardRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		err := u.rewardRepository.Delete(tx, uint(req.RewardId))
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	return nil
}

// TODO 過去のものをGetできるようにする
func (u *RewardUsecaseImpl) GetRewardList(ctx context.Context, req *odachin.GetRewardListRequest) ([]models.RewardPeriod, error) {
	var rewards []models.RewardPeriod
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		var err error
		rewards, err = u.rewardPeriodRepository.GetWithReward(tx, "rewards.to_user_id = ? AND rewards.period_type = ? AND is_editable = ?", user_id, req.RewardType.String(), true)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	return rewards, nil
}

func (u *RewardUsecaseImpl) GetUncompletedRewardCount(ctx context.Context) (*odachin.GetUncompletedRewardCountResponse, error) {
	var rewardCount odachin.GetUncompletedRewardCountResponse
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		now := time.Now()
		daily_count, err := u.rewardPeriodRepository.Count(tx, "rewards.to_user_id = ? AND rewards.period_type = ? AND start_date < ? AND end_date > ? AND is_completed = ?", user_id, "DAILY", now, now, false)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		weekly_count, err := u.rewardPeriodRepository.Count(tx, "rewards.to_user_id = ? AND rewards.period_type = ? AND start_date < ? AND end_date > ? AND is_completed = ?", user_id, "WEEKLY", now, now, false)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		monthly_count, err := u.rewardPeriodRepository.Count(tx, "rewards.to_user_id = ? AND rewards.period_type = ? AND start_date < ? AND end_date > ? AND is_completed = ?", user_id, "MONTHLY", now, now, false)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		rewardCount = odachin.GetUncompletedRewardCountResponse{
			DailyCount:   daily_count,
			WeeklyCount:  weekly_count,
			MonthlyCount: monthly_count,
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	return &rewardCount, nil
}
