package usecase

import (
	"context"
	"database/sql"

	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type FamilyUsecase interface {
	CreateGroup(ctx context.Context, req *odachin.CreateGroupRequest) error
	InviteUser(ctx context.Context, req *odachin.InviteUserRequest) error
	AcceptInvitation(ctx context.Context, req *odachin.AcceptInvitationRequest) error
	GetFamilyInfo(ctx context.Context) ([]models.User, *models.Family, error)
	GetInvitationList(ctx context.Context) ([]*odachin.InvitationMember, error)
}

type FamilyUsecaseImpl struct {
	db                   *gorm.DB
	userRepository       repository.UserRepository
	familyRepository     repository.FamilyRepository
	invitationRepository repository.InvitationRepository
}

func NewFamilyUsecase(db *gorm.DB) FamilyUsecase {
	return &FamilyUsecaseImpl{
		db:                   db,
		userRepository:       repository.NewUserRepository(),
		familyRepository:     repository.NewFamilyRepository(),
		invitationRepository: repository.NewInvitationRepository(),
	}
}

func (u *FamilyUsecaseImpl) CreateGroup(ctx context.Context, req *odachin.CreateGroupRequest) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		family := &models.Family{
			FamilyName: req.FamilyName,
		}
		family, err := u.familyRepository.Save(tx, family)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
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

func (u *FamilyUsecaseImpl) InviteUser(ctx context.Context, req *odachin.InviteUserRequest) error {
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

func (u *FamilyUsecaseImpl) AcceptInvitation(ctx context.Context, req *odachin.AcceptInvitationRequest) error {
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

func (u *FamilyUsecaseImpl) GetFamilyInfo(ctx context.Context) ([]models.User, *models.Family, error) {
	var members []models.User
	var family models.Family
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		var err error
		user, err := u.userRepository.Get(tx, user_id)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		if user.FamilyID == nil {
			return status.Errorf(codes.NotFound, "user is not in family")
		}
		family, err = u.familyRepository.Get(tx, *user.FamilyID)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		members, err = u.userRepository.GetByConditions(tx, "family_id = ?", family.FamilyID)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		return nil
	})

	return members, &family, err
}

func (u *FamilyUsecaseImpl) GetInvitationList(ctx context.Context) ([]*odachin.InvitationMember, error) {
	var invitationList []*odachin.InvitationMember
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user_id := ctx.Value("user_id").(string)
		user, err := u.userRepository.Get(tx, user_id)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		invitations, err := u.invitationRepository.GetByFamilyId(tx, *user.FamilyID)
		if err != nil {
			return status.Errorf(codes.Internal, "database error: %v", err)
		}
		for _, invitation := range invitations {
			user, err := u.userRepository.Get(tx, invitation.ToUserID)
			if err != nil {
				return status.Errorf(codes.Internal, "database error: %v", err)
			}
			//TODO 直す
			invitationList = append(invitationList, &odachin.InvitationMember{
				UserId:         user.UserID,
				Name:           user.UserName,
				AvatarImageUrl: user.AvatarImageUrl,
				InvitationId:   uint64(invitation.InvitationID),
			})
		}
		return nil
	})

	return invitationList, err
}
