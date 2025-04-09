package usecase

import (
	"context"

	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/domain"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/repository"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// trait
type UseCaseImpl interface {
	CreateUser(ctx context.Context, req *odachin.CreateUserRequest) (string, error)
	Login(ctx context.Context, req *odachin.LoginRequest) (string, error)
	CreateGroup(ctx context.Context, req *odachin.CreateGroupRequest) error
}

type UseCase struct {
	userRepository   repository.UserRepository
	familyRepository repository.FamilyRepository
	walletRepository repository.WalletRepository
}

func New(db *gorm.DB) UseCaseImpl {
	return &UseCase{userRepository: repository.NewUserRepository(db), familyRepository: repository.NewFamilyRepository(db), walletRepository: repository.NewWalletRepository(db)}
}

func (u *UseCase) CreateUser(ctx context.Context, req *odachin.CreateUserRequest) (string, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	// req を models.User に変換
	user := &models.User{
		UserID:   req.UserId,
		UserName: req.Name,
		Email:    req.Email,
		Password: string(hashed),
		Role:     req.Role.String(),
	}
	err := u.userRepository.Save(user)
	if err != nil {
		return "", status.Errorf(codes.Internal, "database error: %v", err)
	}
	if user.Role == "CHILD" {
		wallet := &models.Wallet{
			UserID: req.UserId,
		}
		err := u.walletRepository.Save(wallet)
		if err != nil {
			return "", status.Errorf(codes.Internal, "database error: %v", err)
		}

	}
	token, err := domain.GenerateToken(user.UserID)
	if err != nil {
		return "", status.Errorf(codes.Internal, "token generation error: %v", err)
	}
	return token, nil

}

func (u *UseCase) Login(ctx context.Context, req *odachin.LoginRequest) (string, error) {
	user, err := u.userRepository.Get(req.UserId)
	if err != nil {
		return "", status.Errorf(codes.Internal, "database error: %v", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, "invalid password")
	}
	token, err := domain.GenerateToken(user.UserID)
	if err != nil {
		return "", status.Errorf(codes.Internal, "token generation error: %v", err)
	}

	return token, nil
}

func (u *UseCase) CreateGroup(ctx context.Context, req *odachin.CreateGroupRequest) error {
	user_id, err := domain.ExtractTokenMetadata(ctx)

	if err != nil {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}
	family := &models.Family{
		FamilyName: req.FamilyName,
	}

	family, err = u.familyRepository.Save(family)
	if err != nil {
		return status.Errorf(codes.Internal, "database error: %v", err)
	}

	user := &models.User{
		UserID:   user_id,
		FamilyID: &family.FamilyID,
	}
	u.userRepository.UpdateUser(user)

	return nil
}
