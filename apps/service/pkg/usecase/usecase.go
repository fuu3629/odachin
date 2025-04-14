package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/client"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/domain"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/repository"
	"github.com/iancoleman/strcase"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

// trait
type UseCase interface {
	CreateUser(ctx context.Context, req *odachin.CreateUserRequest) (string, error)
	UpdateUser(ctx context.Context, req *odachin.UpdateUserRequest) error
	Login(ctx context.Context, req *odachin.LoginRequest) (string, error)
	CreateGroup(ctx context.Context, req *odachin.CreateGroupRequest) error
	InviteUser(ctx context.Context, req *odachin.InviteUserRequest) error
	AcceptInvitation(ctx context.Context, req *odachin.AcceptInvitationRequest) error
	RegisterReward(ctx context.Context, req *odachin.RegisterRewardRequest) error
	DeleteReward(ctx context.Context, req *odachin.DeleteRewardRequest) error
	RegisterAllowance(ctx context.Context, req *odachin.RegisterAllowanceRequest) error
	UpdateAllowance(ctx context.Context, req *odachin.UpdateAllowanceRequest) error
	GetUserInfo(ctx context.Context, req *odachin.GetUserInfoRequest) (*models.User, error)
}

type UseCaseImpl struct {
	userRepository       repository.UserRepository
	familyRepository     repository.FamilyRepository
	invitationRepository repository.InvitationRepository
	walletRepository     repository.WalletRepository
	rewardRepository     repository.RewardRepository
	allowanceRepository  repository.AllowanceRepository
	db                   *gorm.DB
	s3Client             client.AwsS3Client
}

func New(db *gorm.DB) UseCase {
	return &UseCaseImpl{
		userRepository:       repository.NewUserRepository(),
		familyRepository:     repository.NewFamilyRepository(),
		invitationRepository: repository.NewInvitationRepository(),
		walletRepository:     repository.NewWalletRepository(),
		rewardRepository:     repository.NewRewardRepository(),
		allowanceRepository:  repository.NewAllowanceRepository(),
		db:                   db,
		s3Client:             client.NewAwsS3Client(),
	}
}

func (u *UseCaseImpl) CreateUser(ctx context.Context, req *odachin.CreateUserRequest) (string, error) {
	var token string
	err := u.db.Transaction(func(tx *gorm.DB) error {
		fmt.Println(req.Password)

		hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		// req を models.User に変換
		fmt.Println(req.Role)
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

func (u *UseCaseImpl) UpdateUser(ctx context.Context, req *odachin.UpdateUserRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		user := make(map[string]interface{})
		user["user_id"] = user_id
		user["user_name"] = req.Name
		user["email"] = req.Email
		if req.ProfileImage != nil {
			avaterImageUrl, err := u.s3Client.PutObject(ctx, "odachin-dev", "avaters", req.ProfileImage)
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

func (u *UseCaseImpl) Login(ctx context.Context, req *odachin.LoginRequest) (string, error) {
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

func (u *UseCaseImpl) CreateGroup(ctx context.Context, req *odachin.CreateGroupRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		family := &models.Family{
			FamilyName: req.FamilyName,
		}

		family, err := u.familyRepository.Save(tx, family)

		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}

		// user := &models.User{
		// 	UserID:   user_id,
		// 	FamilyID: &family.FamilyID,
		// }
		user := make(map[string]interface{})
		user["user_id"] = user_id
		user["family_id"] = family.FamilyID
		u.userRepository.Update(tx, user)

		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCaseImpl) InviteUser(ctx context.Context, req *odachin.InviteUserRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		user, err := u.userRepository.Get(tx, user_id)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
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
		return nil

	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCaseImpl) AcceptInvitation(ctx context.Context, req *odachin.AcceptInvitationRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		invitation, err := u.invitationRepository.Get(tx, uint(req.InvitationId))
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}

		if invitation.IsAccepted {
			return status.Errorf(codes.Internal, "already accepted invitation")
		}

		user := make(map[string]interface{})
		user["user_id"] = user_id
		user["family_id"] = invitation.FamilyID
		err = u.userRepository.Update(tx, user)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		invitation.IsAccepted = true

		err = u.invitationRepository.Update(tx, &invitation)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCaseImpl) RegisterReward(ctx context.Context, req *odachin.RegisterRewardRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		reward := &models.Reward{
			ToUserID: req.ToUserId,
			Amount:   float64(req.Amount),
			Reason:   req.Reason,
		}

		err := u.rewardRepository.Save(tx, reward)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCaseImpl) DeleteReward(ctx context.Context, req *odachin.DeleteRewardRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		err := u.rewardRepository.Delete(tx, uint(req.RewardId))
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCaseImpl) RegisterAllowance(ctx context.Context, req *odachin.RegisterAllowanceRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		var dayOfWeek *string
		if req.DayOfWeek == nil {
			dayOfWeek = nil
		} else {
			tmp := req.DayOfWeek.String()
			dayOfWeek = &tmp
		}
		allowance := &models.Allowance{
			FromUserID:   user_id,
			ToUserID:     req.ToUserId,
			Amount:       req.Amount,
			IntervalType: req.IntervalType.String(),
			Interval:     req.Interval,
			Date:         req.Date,
			DayOfWeek:    dayOfWeek,
		}
		err := u.allowanceRepository.Save(tx, allowance)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCaseImpl) UpdateAllowance(ctx context.Context, req *odachin.UpdateAllowanceRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		updateFields, err := ProtoToMap(req)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to convert request to map: %v", err)
		}
		err = u.allowanceRepository.Update(tx, updateFields)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCaseImpl) GetUserInfo(ctx context.Context, req *odachin.GetUserInfoRequest) (*models.User, error) {
	var userInfo models.User
	err := u.db.Transaction(func(tx *gorm.DB) error {
		// _, err := domain.ExtractTokenMetadata(ctx)
		// if err != nil {
		// 	return status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		// }
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

func ProtoToMap(msg proto.Message) (map[string]interface{}, error) {
	jsonBytes, err := protojson.Marshal(msg)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, err
	}
	// Convert keys to snake_case
	result = ToSnakeCaseMap(result)
	return result, nil
}
func ToSnakeCaseMap(m map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m {
		snakeKey := strcase.ToSnake(k)
		result[snakeKey] = v
	}
	return result
}
