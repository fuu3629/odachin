package usecase

import (
	"github.com/fuu3629/odachin/apps/service/pkg/infrastructure/repository"
	"gorm.io/gorm"
)

// trait
type UseCaseImpl interface {
	CreateUser()
}

type UseCase struct {
	userRepository repository.UserRepository
	// Add other repositories or services as needed
}

func New(db *gorm.DB) UseCaseImpl {
	return &UseCase{userRepository: repository.NewUserRepository(db)}
}

func (u *UseCase) CreateUser() {
	// hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

}
