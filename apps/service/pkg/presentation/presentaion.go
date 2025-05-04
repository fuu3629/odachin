package presentation

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin/odachinconnect"
	"github.com/fuu3629/odachin/apps/service/pkg/middleware"
	"github.com/fuu3629/odachin/apps/service/pkg/usecase"
	"gorm.io/gorm"
)

type ServerStruct struct {
	authUsecase        usecase.AuthUsecase
	familyUsecase      usecase.FamilyUsecase
	allowanceUsecase   usecase.AllowanceUsecase
	rewardUsecase      usecase.RewardUsecase
	transactionUsecase usecase.TransactionUsecase
	usageUsecase       usecase.UsageUsecase
}

func NewServer(mux *http.ServeMux, db *gorm.DB) {
	server := &ServerStruct{
		authUsecase:        usecase.NewAuthUsecase(db),
		familyUsecase:      usecase.NewFamilyUsecase(db),
		allowanceUsecase:   usecase.NewAllowanceUsecase(db),
		rewardUsecase:      usecase.NewRewardUsecase(db),
		transactionUsecase: usecase.NewTransactionUsecase(db),
		usageUsecase:       usecase.NewUsageUsecase(db),
	}

	authInterceptor := middleware.NewAuthInterceptor()
	recoveryInterceptor := middleware.NewRecoveryInterceptor()
	loggerInterceptor := middleware.NewLoggerInterceptor()
	validateInterceptor := middleware.NewValidateInterceptor()

	intercepters := connect.WithInterceptors(
		loggerInterceptor,
		recoveryInterceptor,
		validateInterceptor,
		authInterceptor,
	)
	mux.Handle(odachinconnect.NewAuthServiceHandler(server, intercepters))
	mux.Handle(odachinconnect.NewAllowanceServiceHandler(server, intercepters))
	mux.Handle(odachinconnect.NewFamilyServiceHandler(server, intercepters))
	mux.Handle(odachinconnect.NewRewardServiceHandler(server, intercepters))
	mux.Handle(odachinconnect.NewTransactionServiceHandler(server, intercepters))
	mux.Handle(odachinconnect.NewUsageServiceHandler(server, intercepters))
}
