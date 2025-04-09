package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type InvitationRepository interface {
	Get(tx *gorm.DB, id uint) (models.Invitation, error)
	Save(tx *gorm.DB, param *models.Invitation) error
}

type InvitationRepositoryImpl struct {
}

func NewInvitationRepository() InvitationRepository {
	return &InvitationRepositoryImpl{}
}

func (r *InvitationRepositoryImpl) Get(tx *gorm.DB, id uint) (models.Invitation, error) {
	var invitation models.Invitation
	if err := tx.Where("invitation_id = ?", id).First(&invitation).Error; err != nil {
		return models.Invitation{}, err
	}
	return invitation, nil
}

func (r *InvitationRepositoryImpl) Save(tx *gorm.DB, invitation *models.Invitation) error {
	if err := tx.Create(&invitation).Error; err != nil {
		return err
	}
	return nil
}
