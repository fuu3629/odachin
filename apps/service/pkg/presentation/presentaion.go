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
	useCase usecase.UseCase
	odachin.UnimplementedOdachinServiceServer
}

// TODO Roleによる認可を実装する
func (s *ServerStruct) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	if fullMethodName == "/odachin.OdachinService/CreateUser" || fullMethodName == "/odachin.OdachinService/Login" {
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
	userGrpc := &ServerStruct{useCase: usecase.New(db)}
	odachin.RegisterOdachinServiceServer(grpcServer, userGrpc)
}

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

func (s *ServerStruct) GetOwnInfo(ctx context.Context, req *emptypb.Empty) (*odachin.GetOwnInfoResponse, error) {
	userInfo, err := s.useCase.GetOwnInfo(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ToOwnInfoResponse(userInfo), nil
}
