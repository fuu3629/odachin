package presentation

import (
	"context"

	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/pkg/presentation/dto"
	"github.com/fuu3629/odachin/apps/service/pkg/usecase"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type ServerStruct struct {
	useCase usecase.UseCase
	odachin.UnimplementedOdachinServiceServer
}

func NewServer(grpcServer *grpc.Server, db *gorm.DB) {
	userGrpc := &ServerStruct{useCase: usecase.New(db)}
	odachin.RegisterOdachinServiceServer(grpcServer, userGrpc)
}

// func (s *ServerStruct) GetUser(ctx context.Context, req *odachin.CreateUserRequest) (*odachin.CreateUserResponse, error) {
// 	// Implement the logic to get user details using the use case
// 	// For example:
// 	// user, err := s.useCase.GetUser(req.Id)
// 	// if err != nil {
// 	//     return nil, err
// 	// }
// 	// return &odachin.GetUserResponse{User: user}, nil

// 	return nil, nil // Placeholder return
// }

func (s *ServerStruct) CreateUser(ctx context.Context, req *odachin.CreateUserRequest) (*odachin.CreateUserResponse, error) {
	token, err := s.useCase.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &odachin.CreateUserResponse{Token: token}, nil
}

func (s *ServerStruct) UpdateUser(ctx context.Context, req *odachin.UpdateUserRequest) (*emptypb.Empty, error) {
	err := s.useCase.UpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *ServerStruct) Login(ctx context.Context, req *odachin.LoginRequest) (*odachin.LoginResponse, error) {
	token, err := s.useCase.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return &odachin.LoginResponse{Token: token}, nil
}

func (s *ServerStruct) CreateGroup(ctx context.Context, req *odachin.CreateGroupRequest) (*emptypb.Empty, error) {
	err := s.useCase.CreateGroup(ctx, req)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *ServerStruct) InviteUser(ctx context.Context, req *odachin.InviteUserRequest) (*emptypb.Empty, error) {
	err := s.useCase.InviteUser(ctx, req)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *ServerStruct) AcceptInvitation(ctx context.Context, req *odachin.AcceptInvitationRequest) (*emptypb.Empty, error) {
	err := s.useCase.AcceptInvitation(ctx, req)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *ServerStruct) RegisterReward(ctx context.Context, req *odachin.RegisterRewardRequest) (*emptypb.Empty, error) {
	err := s.useCase.RegisterReward(ctx, req)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *ServerStruct) DeleteReward(ctx context.Context, req *odachin.DeleteRewardRequest) (*emptypb.Empty, error) {
	err := s.useCase.DeleteReward(ctx, req)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *ServerStruct) RegisterAllowance(ctx context.Context, req *odachin.RegisterAllowanceRequest) (*emptypb.Empty, error) {
	err := s.useCase.RegisterAllowance(ctx, req)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *ServerStruct) UpdateAllowance(ctx context.Context, req *odachin.UpdateAllowanceRequest) (*emptypb.Empty, error) {
	err := s.useCase.UpdateAllowance(ctx, req)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *ServerStruct) GetUserInfo(ctx context.Context, req *odachin.GetUserInfoRequest) (*odachin.GetUserInfoResponse, error) {
	userInfo, err := s.useCase.GetUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return dto.ToUserInfoResponse(userInfo), nil
}
