package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/fuu3629/odachin/apps/service/pkg/assets"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/domain"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type RewardUsecase interface {
	Reward() error
	RegisterReward(ctx context.Context, req *odachin.RegisterRewardRequest) error
	DeleteReward(ctx context.Context, req *odachin.DeleteRewardRequest) error
	GetRewardList(ctx context.Context, req *odachin.GetRewardListRequest) ([]models.RewardPeriod, error)
	GetChildRewardList(ctx context.Context, req *odachin.GetChildRewardListRequest) ([]models.Reward, error)
	GetUncompletedRewardCount(ctx context.Context) (*odachin.GetUncompletedRewardCountResponse, error)
	ReportReward(ctx context.Context, req *odachin.ReportRewardRequest) error
	GetReportedRewardList(ctx context.Context) ([]models.RewardPeriod, error)
	ApproveReward(ctx context.Context, req *odachin.ApproveRewardRequest) error
}

type RewardUsecaseImpl struct {
	db                     *gorm.DB
	rewardRepository       repository.RewardRepository
	rewardPeriodRepository repository.RewardPeriodRepository
	walletRepository       repository.WalletRepository
	transactionRepository  repository.TransactionRepository
}

func NewRewardUsecase(db *gorm.DB) RewardUsecase {
	return &RewardUsecaseImpl{
		db:                     db,
		rewardRepository:       repository.NewRewardRepository(),
		rewardPeriodRepository: repository.NewRewardPeriodRepository(),
		walletRepository:       repository.NewWalletRepository(),
		transactionRepository:  repository.NewTransactionRepository(),
	}
}

func (u *RewardUsecaseImpl) Reward() error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		reward, err := u.rewardRepository.GetByCondition(tx, "period_type = ?", "DAILY")
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		err = domain.RewardDomain(reward, tx)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		now := time.Now()
		if now.Weekday() == time.Monday {
			reward, err := u.rewardRepository.GetByCondition(tx, "period_type = ?", "WEEKLY")
			if err != nil {
				return status.Errorf(codes.Internal, "database error: %v", err)
			}
			err = domain.RewardDomain(reward, tx)
			if err != nil {
				return status.Errorf(codes.Internal, "database error: %v", err)
			}
		}
		if now.Day() == 1 {
			reward, err := u.rewardRepository.GetByCondition(tx, "period_type = ?", "MONTHLY")
			if err != nil {
				return status.Errorf(codes.Internal, "database error: %v", err)
			}
			err = domain.RewardDomain(reward, tx)
			if err != nil {
				return status.Errorf(codes.Internal, "database error: %v", err)
			}
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	return nil
}

func (u *RewardUsecaseImpl) RegisterReward(ctx context.Context, req *odachin.RegisterRewardRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
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
		now := time.Now()
		var err error
		rewards, err = u.rewardPeriodRepository.GetWithReward(tx, "rewards.to_user_id = ? AND rewards.period_type = ? AND start_date < ? AND end_date > ?", user_id, req.RewardType.String(), now, now)
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

func (u *RewardUsecaseImpl) GetChildRewardList(ctx context.Context, req *odachin.GetChildRewardListRequest) ([]models.Reward, error) {
	var rewards []models.Reward
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		var err error
		rewards, err = u.rewardRepository.GetByCondition(tx, "from_user_id = ? AND to_user_id = ? AND period_type = ?", user_id, req.ChildId, req.RewardType.String())
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
		enum := []string{"IN_PROGRESS", "REPORTED", "REJECTED"}
		daily_count, err := u.rewardPeriodRepository.Count(tx, "rewards.to_user_id = ? AND rewards.period_type = ? AND start_date < ? AND end_date > ? AND status IN ?", user_id, "DAILY", now, now, enum)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		weekly_count, err := u.rewardPeriodRepository.Count(tx, "rewards.to_user_id = ? AND rewards.period_type = ? AND start_date < ? AND end_date > ? AND status IN ?", user_id, "WEEKLY", now, now, enum)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		monthly_count, err := u.rewardPeriodRepository.Count(tx, "rewards.to_user_id = ? AND rewards.period_type = ? AND start_date < ? AND end_date > ? AND status IN ?", user_id, "MONTHLY", now, now, enum)
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

func (u *RewardUsecaseImpl) ReportReward(ctx context.Context, req *odachin.ReportRewardRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		rewardPeriodWithReward, err := u.rewardPeriodRepository.GetWithReward(tx, "reward_period_id = ?", req.RewardPeriodId)
		rewardPeriod := rewardPeriodWithReward[0]
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		if rewardPeriod.Reward.ToUserID != user_id {
			return status.Errorf(codes.Internal, "not your reward")
		}
		if rewardPeriod.Status == "COMPLETED" || rewardPeriod.Status == "REPORTED" {
			return status.Errorf(codes.Internal, "already completed or reported reward")
		}
		reward_period := make(map[string]interface{})
		reward_period["reward_period_id"] = rewardPeriod.RewardPeriodID
		reward_period["status"] = "REPORTED"
		err = u.rewardPeriodRepository.Update(tx, reward_period)
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

func (u *RewardUsecaseImpl) GetReportedRewardList(ctx context.Context) ([]models.RewardPeriod, error) {
	var rewards []models.RewardPeriod
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		var err error
		rewards, err = u.rewardPeriodRepository.GetWithReward(tx, "rewards.from_user_id = ? AND status = ?", user_id, "REPORTED")
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

func (u *RewardUsecaseImpl) ApproveReward(ctx context.Context, req *odachin.ApproveRewardRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		rewardPeriodWithReward, err := u.rewardPeriodRepository.GetWithReward(tx, "reward_period_id = ?", req.RewardPeriodId)
		rewardPeriod := rewardPeriodWithReward[0]
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		if rewardPeriod.Reward.FromUserID != user_id {
			return status.Errorf(codes.Internal, "not your reward")
		}
		if rewardPeriod.Status == "COMPLETED" {
			return status.Errorf(codes.Internal, "already completed or reported reward")
		}
		reward_period := make(map[string]interface{})
		reward_period["reward_period_id"] = rewardPeriod.RewardPeriodID
		reward_period["status"] = "COMPLETED"
		err = u.rewardPeriodRepository.Update(tx, reward_period)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		child_user_id := rewardPeriod.Reward.ToUserID
		wallet, err := u.walletRepository.GetByUserId(tx, child_user_id)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		wallet.Balance += rewardPeriod.Reward.Amount
		err = u.walletRepository.Update(tx, &wallet)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		transaction := &models.Transaction{
			FromUserID: user_id,
			ToUserID:   child_user_id,
			Amount:     rewardPeriod.Reward.Amount,
			Type:       "REWARD",
		}
		err = u.transactionRepository.Save(tx, transaction)
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
