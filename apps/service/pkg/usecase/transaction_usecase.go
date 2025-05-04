package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type TransactionUsecase interface {
	GetTransactionList(ctx context.Context, req *odachin.GetTransactionListRequest) ([]*models.Transaction, error)
}

type TransactionUsecaseImpl struct {
	db                    *gorm.DB
	TransactionRepository repository.TransactionRepository
}

func NewTransactionUsecase(db *gorm.DB) TransactionUsecase {
	return &TransactionUsecaseImpl{
		db:                    db,
		TransactionRepository: repository.NewTransactionRepository(),
	}
}

func (u *TransactionUsecaseImpl) GetTransactionList(ctx context.Context, req *odachin.GetTransactionListRequest) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := u.db.Transaction(func(tx *gorm.DB) error {
		var start, end time.Time
		user_id := ctx.Value("user_id").(string)
		start_date := req.StartDay
		end_date := req.EndDay
		if start_date == nil && end_date == nil {
			start = time.Date(int(req.StartYear), time.Month(req.StartMonth), 1, 0, 0, 0, 0, time.Local)
			end = time.Date(int(req.EndYear), time.Month(req.EndMonth), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, 0)
		} else {
			start = time.Date(int(req.StartYear), time.Month(req.StartMonth), int(*start_date), 0, 0, 0, 0, time.Local)
			end = time.Date(int(req.EndYear), time.Month(req.EndMonth), int(*end_date), 0, 0, 0, 0, time.Local)
		}
		var err error
		transactions, err = u.TransactionRepository.GetByCondition(tx, "to_user_id = ? AND created_at > ? AND created_at < ?", user_id, start, end)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "transaction error: %v", err)
	}
	return transactions, nil
}
