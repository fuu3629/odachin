package presentation

import (
	"context"

	"connectrpc.com/connect"
	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/pkg/presentation/dto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerStruct) CreateUser(ctx context.Context, req *connect.Request[odachin.CreateUserRequest]) (*connect.Response[odachin.CreateUserResponse], error) {
	token, err := s.authUsecase.CreateUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&odachin.CreateUserResponse{Token: token}), nil
}

func (s *ServerStruct) UpdateUser(ctx context.Context, req *connect.Request[odachin.UpdateUserRequest]) (*connect.Response[emptypb.Empty], error) {
	err := s.authUsecase.UpdateUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *ServerStruct) Login(ctx context.Context, req *connect.Request[odachin.LoginRequest]) (*connect.Response[odachin.LoginResponse], error) {
	token, err := s.authUsecase.Login(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&odachin.LoginResponse{Token: token}), nil
}

func (s *ServerStruct) GetUserInfo(ctx context.Context, req *connect.Request[odachin.GetUserInfoRequest]) (*connect.Response[odachin.GetUserInfoResponse], error) {
	userInfo, err := s.authUsecase.GetUserInfo(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(dto.ToUserInfoResponse(userInfo)), nil
}

func (s *ServerStruct) GetOwnInfo(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[odachin.GetOwnInfoResponse], error) {
	userInfo, err := s.authUsecase.GetOwnInfo(ctx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(dto.ToOwnInfoResponse(userInfo)), nil
}
