package presentation

import (
	"context"

	"connectrpc.com/connect"
	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/pkg/presentation/dto"
)

func (s ServerStruct) GetTransactionList(ctx context.Context, req *connect.Request[odachin.GetTransactionListRequest]) (*connect.Response[odachin.GetTransactionListResponse], error) {
	transactionList, err := s.transactionUsecase.GetTransactionList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(dto.ToGetTransactionListResponse(transactionList)), nil
}
