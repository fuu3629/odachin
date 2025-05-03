package presentation

import (
	"context"

	"connectrpc.com/connect"
	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/pkg/presentation/dto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerStruct) RegisterAllowance(ctx context.Context, req *connect.Request[odachin.RegisterAllowanceRequest]) (*connect.Response[emptypb.Empty], error) {
	err := s.allowanceUsecase.RegisterAllowance(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *ServerStruct) UpdateAllowance(ctx context.Context, req *connect.Request[odachin.UpdateAllowanceRequest]) (*connect.Response[emptypb.Empty], error) {
	err := s.allowanceUsecase.UpdateAllowance(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *ServerStruct) GetAllowanceByFromUserId(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[odachin.GetAllowanceByFromUserIdResponse], error) {
	allowanceList, userList, err := s.allowanceUsecase.GetAllowanceByFromUserId(ctx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(dto.ToGetAllowanceByFromUserIdResponse(allowanceList, userList)), nil
}

func (s *ServerStruct) Allowance(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	err := s.allowanceUsecase.Allowance()
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
