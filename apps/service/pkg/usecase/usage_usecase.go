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

type UsageUsecase interface {
	ApplicateUsage(ctx context.Context, req *odachin.ApplicateUsageRequest) error
	GetUsageCategories(ctx context.Context) ([]string, error)
	ApproveUsage(ctx context.Context, req *odachin.ApproveUsageRequest) error
	GetUsageApplication(ctx context.Context, req *odachin.GetUsageApplicationRequest) ([]*odachin.UsageApplication, error)
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

func (u *UsageUsecaseImpl) GetUsageCategories(ctx context.Context) ([]string, error) {
	var categories []string
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		usage, err := u.usageRepository.GetByUserId(tx, user_id)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		categories = make([]string, 0)
		for _, u := range usage {
			if u.Category != "" {
				categories = append(categories, u.Category)
			}
		}
		categories_count := assets.CountAndSortByFrequency(categories)
		categories = make([]string, 0)
		for _, c := range categories_count {
			if c.Count > 1 {
				categories = append(categories, c.Value)
			}
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "transaction error: %v", err)
	}
	return categories, nil
}

func (u *UsageUsecaseImpl) ApproveUsage(ctx context.Context, req *odachin.ApproveUsageRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		usage, err := u.usageRepository.GetById(tx, uint(req.UsageId))
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		if usage == nil {
			return status.Errorf(codes.NotFound, "usage not found")
		}
		if usage.Status != "APPLICATED" {
			return status.Errorf(codes.InvalidArgument, "usage already approved")
		}
		usage.Status = "APPROVED"
		err = u.usageRepository.Update(tx, usage)
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

func (u *UsageUsecaseImpl) GetUsageApplication(ctx context.Context, req *odachin.GetUsageApplicationRequest) ([]*odachin.UsageApplication, error) {
	var allUsages []*odachin.UsageApplication
	err := u.db.Transaction(func(tx *gorm.DB) error {
		for _, userId := range req.UserId {
			usages, err := u.usageRepository.GetByUserId(tx, userId)
			if err != nil {
				return status.Errorf(codes.Internal, "database error: %v", err)
			}
			for _, usage := range usages {
				if usage.Status != "APPROVED" {
					allUsages = append(allUsages, &odachin.UsageApplication{
						UsageId:     uint64(usage.UsageID),
						Title:       usage.Title,
						Description: usage.Description,
						Amount:      usage.Amount,
						Category:    usage.Category,
					})

				}
			}
		}

		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "transaction error: %v", err)
	}
	return allUsages, nil
}
