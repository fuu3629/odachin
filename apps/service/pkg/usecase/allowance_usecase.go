package usecase

import (
	"context"
	"database/sql"

	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/fuu3629/odachin/apps/service/pkg/assets"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AllowanceUsecase interface {
	RegisterAllowance(ctx context.Context, req *odachin.RegisterAllowanceRequest) error
	UpdateAllowance(ctx context.Context, req *odachin.UpdateAllowanceRequest) error
}

type AllowanceUsecaseImpl struct {
	db                  *gorm.DB
	allowanceRepository repository.AllowanceRepository
}

func NewAllowanceUsecase(db *gorm.DB) AllowanceUsecase {
	return &AllowanceUsecaseImpl{
		db:                  db,
		allowanceRepository: repository.NewAllowanceRepository(),
	}
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
			Interval:     req.Interval,
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
