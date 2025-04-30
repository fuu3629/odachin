package database

import (
	"context"
	"fmt"

	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/fuu3629/odachin/apps/service/pkg/assets"
	"github.com/fuu3629/odachin/apps/service/pkg/usecase"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	assets.Log("Seeding database...🌱")

	var count int64
	err := db.Model(&models.User{}).Count(&count)
	if err.Error != nil {
		return fmt.Errorf("failed to count users: %v", err)
	}
	if count > 0 {
		assets.Log("Database already seeded")
		return nil
	}
	ctx := context.Background()

	// Create a default family
	family := models.Family{
		FamilyName: "example",
	}

	if err := db.Create(&family).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	auth_usecase := usecase.NewAuthUsecase(db)

	req_create_user1 := &odachin.CreateUserRequest{
		UserId:   "parent1",
		Name:     "parent_Name1",
		Email:    "example1@xxx.com",
		Password: "password",
		Role:     odachin.Role_PARENT,
	}
	_, _ = auth_usecase.CreateUser(ctx, req_create_user1)

	req_create_user2 := &odachin.CreateUserRequest{
		UserId:   "parent2",
		Name:     "parent_Name2",
		Email:    "example2@xxx.com",
		Password: "password",
		Role:     odachin.Role_PARENT,
	}
	_, _ = auth_usecase.CreateUser(ctx, req_create_user2)

	req_create_user3 := &odachin.CreateUserRequest{
		UserId:   "child1",
		Name:     "child_Name1",
		Email:    "example3@xxx.com",
		Password: "password",
		Role:     odachin.Role_CHILD,
	}
	_, _ = auth_usecase.CreateUser(ctx, req_create_user3)

	req_create_user4 := &odachin.CreateUserRequest{
		UserId:   "child2",
		Name:     "child_Name2",
		Email:    "example4@xxx.com",
		Password: "password",
		Role:     odachin.Role_CHILD,
	}
	_, _ = auth_usecase.CreateUser(ctx, req_create_user4)

	// Create a default allowance
	allowance := models.Allowance{
		FromUserID:   "parent2",
		ToUserID:     "child2",
		Amount:       100,
		IntervalType: "DAILY",
	}
	if err := db.Create(&allowance).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	user := make(map[string]interface{})
	user["family_id"] = family.FamilyID
	db.Model(&models.User{}).Where("user_id = ?", "parent2").Updates(&user)
	db.Model(&models.User{}).Where("user_id = ?", "child2").Updates(&user)

	ctx = context.WithValue(ctx, "user_id", "parent2")

	// Create a default reward
	req := &odachin.RegisterRewardRequest{
		ToUserId:    "child2",
		Amount:      100,
		RewardType:  0,
		Title:       "宿題をする",
		Description: "宿題をすることができたら100円あげるよ",
	}
	reward_usecase := usecase.NewRewardUsecase(db)
	err2 := reward_usecase.RegisterReward(ctx, req)
	if err2 != nil {
		fmt.Printf("Failed to register reward: %+v\n", err2)
	}

	req2 := &odachin.RegisterRewardRequest{
		ToUserId:    "child2",
		Amount:      1000,
		RewardType:  1,
		Title:       "宿題をする",
		Description: "宿題をすること",
	}
	err2 = reward_usecase.RegisterReward(ctx, req2)
	if err2 != nil {
		fmt.Printf("Failed to register reward: %+v\n", err2)
	}

	req3 := &odachin.RegisterRewardRequest{
		ToUserId:    "child2",
		Amount:      10000,
		RewardType:  2,
		Title:       "宿題をする",
		Description: "宿題をすること",
	}
	err2 = reward_usecase.RegisterReward(ctx, req3)
	if err2 != nil {
		fmt.Printf("Failed to register reward: %+v\n", err2)
	}

	assets.Log("Database seeded successfully🍀")
	return nil
}
