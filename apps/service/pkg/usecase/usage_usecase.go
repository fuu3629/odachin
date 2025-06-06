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
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type UsageUsecase interface {
	ApplicateUsage(ctx context.Context, req *odachin.ApplicateUsageRequest) error
	GetUsageCategories(ctx context.Context) ([]string, error)
	ApproveUsage(ctx context.Context, req *odachin.ApproveUsageRequest) error
	GetUsageApplication(ctx context.Context, req *odachin.GetUsageApplicationRequest) ([]*odachin.UsageApplication, error)
	GetUsageSummary(ctx context.Context) ([]*odachin.UsageSummary, []*odachin.UsageSummary, error)
	RejectUsage(ctx context.Context, req *odachin.RejectUsageRequest) error
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
		fmt.Println(user_id)
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
		fmt.Println(categories_count)
		categories = make([]string, 0)
		for _, c := range categories_count {
			if c.Count > 0 {
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
			return status.Errorf(codes.InvalidArgument, "usage already approved or rejected")
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
				allUsages = append(allUsages, &odachin.UsageApplication{
					UsageId:     uint64(usage.UsageID),
					Title:       usage.Title,
					Description: usage.Description,
					Amount:      usage.Amount,
					Category:    usage.Category,
					Status:      usage.Status,
					CreatedAt:   timestamppb.New(usage.CreatedAt),
				})
			}
		}

		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "transaction error: %v", err)
	}
	return allUsages, nil
}

func (u *UsageUsecaseImpl) GetUsageSummary(ctx context.Context) ([]*odachin.UsageSummary, []*odachin.UsageSummary, error) {
	var allUsages []*odachin.UsageSummary
	var allUsagesMonthly []*odachin.UsageSummary
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		usage_map := make(map[string]int32)
		usage_map_monthly := make(map[string]int32)
		now := time.Now()
		month_start := time.Date(
			now.Year(),
			now.Month(),
			1,
			0, 0, 0, 0,
			now.Location(),
		)
		month_end := time.Date(
			now.Year(),
			now.Month()+1,
			1,
			0, 0, 0, 0,
			now.Location(),
		)
		usages, err := u.usageRepository.GetByUserId(tx, user_id)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		for _, usage := range usages {
			if usage.Status != "APPROVED" {
				continue
			}
			if _, ok := usage_map[usage.Category]; !ok {
				usage_map[usage.Category] = 0
			}
			usage_map[usage.Category] += usage.Amount
		}

		for _, usage := range usages {
			if usage.Status != "APPROVED" {
				continue
			}
			if usage.CreatedAt.Before(month_start) || !usage.CreatedAt.Before(month_end) {
				continue
			}
			if _, ok := usage_map_monthly[usage.Category]; !ok {
				usage_map_monthly[usage.Category] = 0
			}
			usage_map_monthly[usage.Category] += usage.Amount
		}

		for category, amount := range usage_map {
			allUsages = append(allUsages, &odachin.UsageSummary{
				Category: category,
				Amount:   amount,
			})
		}
		for category, amount := range usage_map_monthly {
			allUsagesMonthly = append(allUsagesMonthly, &odachin.UsageSummary{
				Category: category,
				Amount:   amount,
			})
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, nil, status.Errorf(codes.Internal, "transaction error: %v", err)
	}
	return allUsages, allUsagesMonthly, nil
}

func (u *UsageUsecaseImpl) RejectUsage(ctx context.Context, req *odachin.RejectUsageRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		usage, err := u.usageRepository.GetById(tx, uint(req.UsageId))
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		if usage == nil {
			return status.Errorf(codes.NotFound, "usage not found")
		}
		if usage.Status != "APPLICATED" {
			return status.Errorf(codes.InvalidArgument, "usage already approved or rejected")
		}
		usage.Status = "REJECTED"
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
