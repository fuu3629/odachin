package usecase

import (
	"context"
	"database/sql"

	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UsageUsecase interface {
	ApplicateUsage(ctx context.Context, req *odachin.ApplicateUsageRequest) error
}

type UsageUsecaseImpl struct {
	db              *gorm.DB
	usageRepository repository.UsageRepository
}

func NewUsageUsecase(db *gorm.DB) UsageUsecase {
	return &UsageUsecaseImpl{
		db:              db,
		usageRepository: repository.NewUsageRepository(),
	}
}

func (u *UsageUsecaseImpl) ApplicateUsage(ctx context.Context, req *odachin.ApplicateUsageRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		usage := &models.Usage{
			UserID:      user_id,
			Amount:      req.Amount,
			Title:       req.Title,
			Description: req.Description,
			Category:    req.Category,
			Status:      "APPLICATED",
		}
		err := u.usageRepository.Save(tx, usage)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return status.Errorf(codes.Internal, "transaction error: %v", err)
	}
	return nil
}
