package presentation

import (
	"context"

	"connectrpc.com/connect"
	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerStruct) ApplicateUsage(
	ctx context.Context,
	req *connect.Request[odachin.ApplicateUsageRequest],
) (*connect.Response[emptypb.Empty], error) {
	err := s.usageUsecase.ApplicateUsage(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
