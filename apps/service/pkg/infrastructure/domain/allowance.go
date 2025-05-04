package domain

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func AllowanceDomain(users []models.User, tx *gorm.DB) error {
	walletRepository := repository.NewWalletRepository()
	transactions := []*models.Transaction{}
	for _, us := range users {
		wallet := us.Wallet
		wallet.Balance += us.Allowance.Amount
		err := walletRepository.Update(tx, &wallet)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		transaction := &models.Transaction{
			FromUserID:  us.Allowance.FromUserID,
			ToUserID:    us.Allowance.ToUserID,
			Amount:      us.Allowance.Amount,
			Type:        "ALLOWANCE",
			Title:       "お小遣い",
			Description: "お小遣いを送金しました",
		}
		transactions = append(transactions, transaction)
	}
	transactionRepository := repository.NewTransactionRepository()
	err := transactionRepository.SaveList(tx, transactions)
	if err != nil {
		return status.Errorf(codes.Internal, "database error: %v", err)
	}
	return nil
}
