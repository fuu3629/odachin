package usecase

import (
	"context"
	"database/sql"
	"strings"
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

type AllowanceUsecase interface {
	RegisterAllowance(ctx context.Context, req *odachin.RegisterAllowanceRequest) error
	UpdateAllowance(ctx context.Context, req *odachin.UpdateAllowanceRequest) error
	GetAllowanceByFromUserId(ctx context.Context) ([]models.Allowance, []models.User, error)
	Allowance() error
}

type AllowanceUsecaseImpl struct {
	db                  *gorm.DB
	userRepository      repository.UserRepository
	allowanceRepository repository.AllowanceRepository
	walletRepository    repository.WalletRepository
}

func NewAllowanceUsecase(db *gorm.DB) AllowanceUsecase {
	return &AllowanceUsecaseImpl{
		db:                  db,
		userRepository:      repository.NewUserRepository(),
		allowanceRepository: repository.NewAllowanceRepository(),
		walletRepository:    repository.NewWalletRepository(),
	}
}

func (u *AllowanceUsecaseImpl) Allowance() error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user, err := u.userRepository.GetWithAll(tx, "allowances.interval_type = ?", "DAILY")
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		err = domain.AllowanceDomain(user, tx)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		nowWeek := strings.ToUpper(time.Now().Weekday().String())
		user, err = u.userRepository.GetWithAll(tx, "allowances.interval_type = ? AND allowances.day_of_week = ?", "WEEKLY", nowWeek)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		err = domain.AllowanceDomain(user, tx)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		now := time.Now().Day()
		user, err = u.userRepository.GetWithAll(tx, "allowances.interval_type = ? AND allowances.date = ?", "MONTHLY", now)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		err = domain.AllowanceDomain(user, tx)
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

func (u *AllowanceUsecaseImpl) RegisterAllowance(ctx context.Context, req *odachin.RegisterAllowanceRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		var dayOfWeek *string
		if req.DayOfWeek == nil {
			dayOfWeek = nil
		} else {
			tmp := req.DayOfWeek.String()
			dayOfWeek = &tmp
		}
		allowance := &models.Allowance{
			FromUserID:   user_id,
			ToUserID:     req.ToUserId,
			Amount:       req.Amount,
			IntervalType: req.IntervalType.String(),
			Date:         req.Date,
			DayOfWeek:    dayOfWeek,
		}
		err := u.allowanceRepository.Save(tx, allowance)
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

// TODO ROLEの確認を行う
// TODO 要改修
func (u *AllowanceUsecaseImpl) UpdateAllowance(ctx context.Context, req *odachin.UpdateAllowanceRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		updateFields, err := assets.ProtoToMap(req)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to convert request to map: %v", err)
		}
		err = u.allowanceRepository.Update(tx, updateFields)
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

func (u *AllowanceUsecaseImpl) GetAllowanceByFromUserId(ctx context.Context) ([]models.Allowance, []models.User, error) {
	var allowanceList []models.Allowance
	var to_user []models.User
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		var err error
		allowanceList, err = u.allowanceRepository.GetAllowanceByCondition(tx, "from_user_id = ?", user_id)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		to_user = make([]models.User, len(allowanceList))
		for i, allowance := range allowanceList {
			to_user[i], err = u.userRepository.Get(tx, allowance.ToUserID)
			if err != nil {
				return status.Errorf(codes.Internal, "database error: %v", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	return allowanceList, to_user, nil
}
