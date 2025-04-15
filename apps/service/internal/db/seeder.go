package database

import (
	"fmt"

	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/fuu3629/odachin/apps/service/pkg/assets"
	"golang.org/x/crypto/bcrypt"
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

	// Create a default family
	family := models.Family{
		FamilyName: "example",
	}

	if err := db.Create(&family).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	// Create a default user
	hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	user := models.User{
		UserID:   "parent1",
		UserName: "parent_Name1",
		Email:    "example1@xxx.com",
		Password: string(hashed),
		Role:     "PARENT",
	}
	if err := db.Create(&user).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	user_belog_family := models.User{
		UserID:   "parent2",
		FamilyID: &family.FamilyID,
		UserName: "parent_Name2",
		Email:    "example2@xxx.com",
		Password: string(hashed),
		Role:     "PARENT",
	}
	if err := db.Create(&user_belog_family).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	// Create a default child
	child := models.User{
		UserID:   "child1",
		UserName: "child_Name1",
		Email:    "example3@xxx.com",
		Password: string(hashed),
		Role:     "CHILD",
	}
	if err := db.Create(&child).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	child_belong := models.User{
		UserID:   "child2",
		FamilyID: &family.FamilyID,
		UserName: "child_Name2",
		Email:    "example4@xxx.com",
		Password: string(hashed),
		Role:     "CHILD",
	}
	if err := db.Create(&child_belong).Error; err != nil {
		fmt.Printf("%+v", err)
	}
	assets.Log("Database seeded successfully🍀")
	return nil
}
