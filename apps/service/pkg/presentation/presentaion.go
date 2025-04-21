package presentation

import (
	"context"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/pkg/presentation/dto"
	"github.com/fuu3629/odachin/apps/service/pkg/usecase"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type ServerStruct struct {
	authUsecase      usecase.AuthUsecase
	familyUsecase    usecase.FamilyUsecase
	allowanceUsecase usecase.AllowanceUsecase
	rewardUsecase    usecase.RewardUsecase
	odachin.UnimplementedAuthServiceServer
	odachin.UnimplementedFamilyServiceServer
	odachin.UnimplementedAllowanceServiceServer
	odachin.UnimplementedRewardServiceServer
}

// TODO Roleによる認可を実装する
func (s *ServerStruct) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	fmt.Println("AuthFuncOverride", fullMethodName)
	if fullMethodName == "/odachin.auth.AuthService/CreateUser" || fullMethodName == "/odachin.auth.AuthService/Login" {
		return ctx, nil
	}
	tokenString, err := auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, err
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		newCtx := context.WithValue(ctx, "user_id", claims["user_id"])
		return newCtx, nil
	} else {
		return nil, fmt.Errorf("no userId")
	}
}

func NewServer(grpcServer *grpc.Server, db *gorm.DB) {
	userGrpc := &ServerStruct{
		authUsecase:      usecase.NewAuthUsecase(db),
		familyUsecase:    usecase.NewFamilyUsecase(db),
		allowanceUsecase: usecase.NewAllowanceUsecase(db),
		rewardUsecase:    usecase.NewRewardUsecase(db),
	}
	// odachin.RegisterOdachinServiceServer(grpcServer, userGrpc)
	odachin.RegisterAuthServiceServer(grpcServer, userGrpc)
	odachin.RegisterFamilyServiceServer(grpcServer, userGrpc)
	odachin.RegisterAllowanceServiceServer(grpcServer, userGrpc)
	odachin.RegisterRewardServiceServer(grpcServer, userGrpc)
}

func (s *ServerStruct) CreateUser(ctx context.Context, req *odachin.CreateUserRequest) (*odachin.CreateUserResponse, error) {
	token, err := s.authUsecase.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &odachin.CreateUserResponse{Token: token}, nil
}

func (s *ServerStruct) UpdateUser(ctx context.Context, req *odachin.UpdateUserRequest) (*emptypb.Empty, error) {
	err := s.authUsecase.UpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *ServerStruct) Login(ctx context.Context, req *odachin.LoginRequest) (*odachin.LoginResponse, error) {
	token, err := s.authUsecase.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return &odachin.LoginResponse{Token: token}, nil
}

func (s *ServerStruct) CreateGroup(ctx context.Context, req *odachin.CreateGroupRequest) (*emptypb.Empty, error) {
	err := s.familyUsecase.CreateGroup(ctx, req)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *ServerStruct) InviteUser(ctx context.Context, req *odachin.InviteUserRequest) (*emptypb.Empty, error) {
	err := s.familyUsecase.InviteUser(ctx, req)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *ServerStruct) AcceptInvitation(ctx context.Context, req *odachin.AcceptInvitationRequest) (*emptypb.Empty, error) {
	err := s.familyUsecase.AcceptInvitation(ctx, req)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *ServerStruct) RegisterReward(ctx context.Context, req *odachin.RegisterRewardRequest) (*emptypb.Empty, error) {
	err := s.rewardUsecase.RegisterReward(ctx, req)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *ServerStruct) DeleteReward(ctx context.Context, req *odachin.DeleteRewardRequest) (*emptypb.Empty, error) {
	err := s.rewardUsecase.DeleteReward(ctx, req)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *ServerStruct) RegisterAllowance(ctx context.Context, req *odachin.RegisterAllowanceRequest) (*emptypb.Empty, error) {
	err := s.allowanceUsecase.RegisterAllowance(ctx, req)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *ServerStruct) UpdateAllowance(ctx context.Context, req *odachin.UpdateAllowanceRequest) (*emptypb.Empty, error) {
	err := s.allowanceUsecase.UpdateAllowance(ctx, req)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *ServerStruct) GetUserInfo(ctx context.Context, req *odachin.GetUserInfoRequest) (*odachin.GetUserInfoResponse, error) {
	userInfo, err := s.authUsecase.GetUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return dto.ToUserInfoResponse(userInfo), nil
}

func (s *ServerStruct) GetOwnInfo(ctx context.Context, req *emptypb.Empty) (*odachin.GetOwnInfoResponse, error) {
	userInfo, err := s.authUsecase.GetOwnInfo(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ToOwnInfoResponse(userInfo), nil
}

func (s *ServerStruct) GetRewardList(ctx context.Context, req *odachin.GetRewardListRequest) (*odachin.GetRewardListResponse, error) {
	rewardList, err := s.rewardUsecase.GetRewardList(ctx, req)
	if err != nil {
		return nil, err
	}
	return dto.ToGetRewardListResponse(rewardList), nil
}

func (s *ServerStruct) GetUncompletedRewardCount(ctx context.Context, req *emptypb.Empty) (*odachin.GetUncompletedRewardCountResponse, error) {
	rewardCount, err := s.rewardUsecase.GetUncompletedRewardCount(ctx)
	if err != nil {
		return nil, err
	}
	return rewardCount, nil
}

func (s *ServerStruct) GetFamilyInfo(ctx context.Context, req *emptypb.Empty) (*odachin.GetFamilyInfoResponse, error) {
	member, family, err := s.familyUsecase.GetFamilyInfo(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ToGetFamilyInfoResponse(member, family), nil
}
