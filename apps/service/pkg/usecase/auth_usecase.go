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

type AuthUsecase interface {
	CreateUser(ctx context.Context, req *odachin.CreateUserRequest) (string, error)
	UpdateUser(ctx context.Context, req *odachin.UpdateUserRequest) error
	Login(ctx context.Context, req *odachin.LoginRequest) (string, error)
	GetUserInfo(ctx context.Context, req *odachin.GetUserInfoRequest) (*models.User, error)
	GetOwnInfo(ctx context.Context) (*models.User, error)
}

type AuthUsecaseImpl struct {
	userRepository   repository.UserRepository
	walletRepository repository.WalletRepository
	db               *gorm.DB
	s3Client         client.AwsS3Client
}

func NewAuthUsecase(db *gorm.DB) AuthUsecase {
	return &AuthUsecaseImpl{
		userRepository:   repository.NewUserRepository(),
		walletRepository: repository.NewWalletRepository(),
		db:               db,
		s3Client:         client.NewAwsS3Client(),
	}
}

func (u *AuthUsecaseImpl) CreateUser(ctx context.Context, req *odachin.CreateUserRequest) (string, error) {
	var token string
	err := u.db.Transaction(func(tx *gorm.DB) error {

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
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *AuthUsecaseImpl) UpdateUser(ctx context.Context, req *odachin.UpdateUserRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		user := make(map[string]interface{})
		user["user_id"] = user_id
		user["user_name"] = req.Name
		user["email"] = req.Email
		if req.ProfileImage != nil {
			avaterImageUrl, err := u.s3Client.PutObject(ctx, "odachin-dev", "avatars", req.ProfileImage)
			if err != nil {
				return status.Errorf(codes.Internal, "s3 upload error: %v", err)
			}
			user["avatar_image_url"] = avaterImageUrl
		}
		u.userRepository.Update(tx, user)
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	return nil
}

func (u *AuthUsecaseImpl) Login(ctx context.Context, req *odachin.LoginRequest) (string, error) {
	var token string
	err := u.db.Transaction(func(tx *gorm.DB) error {
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
	if err != nil {
		return "", err
	}
	return token, nil
}

// TODO Family内じゃないと取得できないようにする
func (u *AuthUsecaseImpl) GetUserInfo(ctx context.Context, req *odachin.GetUserInfoRequest) (*models.User, error) {
	var userInfo models.User
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user, err := u.userRepository.Get(tx, req.UserId)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		userInfo = models.User{
			UserID:         req.UserId,
			UserName:       user.UserName,
			Role:           user.Role,
			AvatarImageUrl: user.AvatarImageUrl,
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (u *AuthUsecaseImpl) GetOwnInfo(ctx context.Context) (*models.User, error) {
	var userInfo models.User
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		var err error
		userInfo, err = u.userRepository.Get(tx, user_id)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}
