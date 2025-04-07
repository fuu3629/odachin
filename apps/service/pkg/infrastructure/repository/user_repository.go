package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Get(id string) (models.User, error)
	Save(param *models.User) error
	UpdateUser(param *models.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) Get(id string) (models.User, error) {
	var user models.User
	if err := r.db.Where("user_id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Save(user *models.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) UpdateUser(user *models.User) error {
	if err := r.db.Updates(&user).Error; err != nil {
		return err
	}
	return nil
}
