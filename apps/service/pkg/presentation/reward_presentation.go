package presentation

import (
	"context"

	"connectrpc.com/connect"
	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/pkg/presentation/dto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerStruct) RegisterReward(ctx context.Context, req *connect.Request[odachin.RegisterRewardRequest]) (*connect.Response[emptypb.Empty], error) {
	err := s.rewardUsecase.RegisterReward(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *ServerStruct) DeleteReward(ctx context.Context, req *connect.Request[odachin.DeleteRewardRequest]) (*connect.Response[emptypb.Empty], error) {
	err := s.rewardUsecase.DeleteReward(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *ServerStruct) GetRewardList(ctx context.Context, req *connect.Request[odachin.GetRewardListRequest]) (*connect.Response[odachin.GetRewardListResponse], error) {
	rewardList, err := s.rewardUsecase.GetRewardList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(dto.ToGetRewardListResponse(rewardList)), nil
}

func (s *ServerStruct) GetChildRewardList(ctx context.Context, req *connect.Request[odachin.GetChildRewardListRequest]) (*connect.Response[odachin.GetChildRewardListResponse], error) {
	rewardList, err := s.rewardUsecase.GetChildRewardList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(dto.ToGetChildRewardListResponse(rewardList)), nil
}

func (s *ServerStruct) GetUncompletedRewardCount(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[odachin.GetUncompletedRewardCountResponse], error) {
	rewardCount, err := s.rewardUsecase.GetUncompletedRewardCount(ctx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(rewardCount), nil
}
