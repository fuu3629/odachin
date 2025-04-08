package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Get(tx *gorm.DB, id string) (models.User, error)
	Save(tx *gorm.DB, param *models.User) error
	Update(tx *gorm.DB, param *models.User) error
}

type UserRepositoryImpl struct {
	// db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Get(tx *gorm.DB, id string) (models.User, error) {
	var user models.User
	if err := tx.Where("user_id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Save(tx *gorm.DB, user *models.User) error {
	if err := tx.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Update(tx *gorm.DB, user *models.User) error {
	if err := tx.Updates(&user).Error; err != nil {
		return err
	}
	return nil
}
