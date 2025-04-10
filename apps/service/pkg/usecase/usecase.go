package usecase

import (
	"context"
	"database/sql"

	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/client"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/domain"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/repository"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// trait
type UseCase interface {
	CreateUser(ctx context.Context, req *odachin.CreateUserRequest) (string, error)
	UpdateUser(ctx context.Context, req *odachin.UpdateUserRequest) error
	Login(ctx context.Context, req *odachin.LoginRequest) (string, error)
	CreateGroup(ctx context.Context, req *odachin.CreateGroupRequest) error
	InviteUser(ctx context.Context, req *odachin.InviteUserRequest) error
}

type UseCaseImpl struct {
	userRepository       repository.UserRepository
	familyRepository     repository.FamilyRepository
	invitationRepository repository.InvitationRepository
	walletRepository     repository.WalletRepository
	db                   *gorm.DB
	s3Client             client.AwsS3Client
}

func New(db *gorm.DB) UseCase {
	return &UseCaseImpl{
		userRepository:       repository.NewUserRepository(),
		familyRepository:     repository.NewFamilyRepository(),
		invitationRepository: repository.NewInvitationRepository(),
		walletRepository:     repository.NewWalletRepository(),
		db:                   db,
		s3Client:             client.NewAwsS3Client(),
	}
}

func (u *UseCaseImpl) CreateUser(ctx context.Context, req *odachin.CreateUserRequest) (string, error) {
	var token string
	u.db.Transaction(func(tx *gorm.DB) error {

		hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		// req を models.User に変換
		user := &models.User{
			UserID:   req.UserId,
			UserName: req.Name,
			Email:    req.Email,
			Password: string(hashed),
			Role:     req.Role.String(),
		}
		err := u.userRepository.Save(tx, user)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		if user.Role == "CHILD" {
			wallet := &models.Wallet{
				UserID: req.UserId,
			}
			err := u.walletRepository.Save(tx, wallet)
			if err != nil {
				return status.Errorf(codes.Internal, "database error: %v", err)
			}

		}
		token, err = domain.GenerateToken(user.UserID)
		if err != nil {
			return status.Errorf(codes.Internal, "token generation error: %v", err)
		}
		// トランザクションをコミット
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	return token, nil
}

func (u *UseCaseImpl) UpdateUser(ctx context.Context, req *odachin.UpdateUserRequest) error {
	u.db.Transaction(func(tx *gorm.DB) error {
		user_id, err := domain.ExtractTokenMetadata(ctx)
		if err != nil {
			return status.Errorf(codes.Unauthenticated, "invalid token")
		}
		var avaterImageUrl *string
		if req.ProfileImage != nil {
			avaterImageUrl, err = u.s3Client.PutObject(ctx, "odachin-dev", "avaters", req.ProfileImage)
			if err != nil {
				return status.Errorf(codes.Internal, "s3 upload error: %v", err)
			}
		} else {
			avaterImageUrl = nil
		}

		user := &models.User{
			UserID:         user_id,
			UserName:       *req.Name,
			Email:          *req.Email,
			Password:       *req.Password,
			AvatarImageUrl: avaterImageUrl,
		}
		u.userRepository.Update(tx, user)
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	return nil
}

func (u *UseCaseImpl) Login(ctx context.Context, req *odachin.LoginRequest) (string, error) {
	var token string
	u.db.Transaction(func(tx *gorm.DB) error {
		user, err := u.userRepository.Get(tx, req.UserId)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			return status.Errorf(codes.Unauthenticated, "invalid password")
		}
		token, err = domain.GenerateToken(user.UserID)
		if err != nil {
			return status.Errorf(codes.Internal, "token generation error: %v", err)
		}
		// トランザクションをコミット
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	return token, nil
}

func (u *UseCaseImpl) CreateGroup(ctx context.Context, req *odachin.CreateGroupRequest) error {
	u.db.Transaction(func(tx *gorm.DB) error {
		user_id, err := domain.ExtractTokenMetadata(ctx)

		if err != nil {
			return status.Errorf(codes.Unauthenticated, "invalid token")
		}
		family := &models.Family{
			FamilyName: req.FamilyName,
		}

		family, err = u.familyRepository.Save(tx, family)

		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}

		user := &models.User{
			UserID:   user_id,
			FamilyID: &family.FamilyID,
		}
		u.userRepository.Update(tx, user)

		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	return nil
}

func (u *UseCaseImpl) InviteUser(ctx context.Context, req *odachin.InviteUserRequest) error {
	u.db.Transaction(func(tx *gorm.DB) error {
		user_id, err := domain.ExtractTokenMetadata(ctx)
		if err != nil {
			return status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}
		user, err := u.userRepository.Get(tx, user_id)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		// req を models.Invitation に変換
		invitation := &models.Invitation{
			FamilyID:   user.FamilyID,
			FromUserID: user_id,
			ToUserID:   req.ToUserId,
			IsAccepted: false,
		}
		err = u.invitationRepository.Save(tx, invitation)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		// トランザクションをコミット
		return nil

	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	return nil
}
