package presentation

import (
	"context"

	"connectrpc.com/connect"
	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/pkg/presentation/dto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerStruct) CreateGroup(ctx context.Context, req *connect.Request[odachin.CreateGroupRequest]) (*connect.Response[emptypb.Empty], error) {
	err := s.familyUsecase.CreateGroup(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *ServerStruct) InviteUser(ctx context.Context, req *connect.Request[odachin.InviteUserRequest]) (*connect.Response[emptypb.Empty], error) {
	err := s.familyUsecase.InviteUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *ServerStruct) AcceptInvitation(ctx context.Context, req *connect.Request[odachin.AcceptInvitationRequest]) (*connect.Response[emptypb.Empty], error) {
	err := s.familyUsecase.AcceptInvitation(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *ServerStruct) GetFamilyInfo(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[odachin.GetFamilyInfoResponse], error) {
	member, family, err := s.familyUsecase.GetFamilyInfo(ctx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(dto.ToGetFamilyInfoResponse(member, family)), nil
}

func (s *ServerStruct) GetInvitationList(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[odachin.GetInvitationListResponse], error) {
	invitationList, err := s.familyUsecase.GetInvitationList(ctx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&odachin.GetInvitationListResponse{InvitationMembers: invitationList}), nil
}
