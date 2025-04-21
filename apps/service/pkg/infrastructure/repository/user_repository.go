package repository

import (
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Get(tx *gorm.DB, id string) (models.User, error)
	GetByConditions(tx *gorm.DB, condition string, args ...interface{}) ([]models.User, error)
	Save(tx *gorm.DB, param *models.User) error
	Update(tx *gorm.DB, user map[string]interface{}) error
}

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Get(tx *gorm.DB, id string) (models.User, error) {
	var user models.User
	if err := tx.Where("user_id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetByConditions(tx *gorm.DB, condition string, args ...interface{}) ([]models.User, error) {
	var users []models.User
	if err := tx.Where(condition, args...).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepositoryImpl) Save(tx *gorm.DB, user *models.User) error {
	if err := tx.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Update(tx *gorm.DB, user map[string]interface{}) error {
	id := user["user_id"]
	delete(user, "user_id")
	if err := tx.Model(&models.User{}).Where("user_id = ?", id).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}
