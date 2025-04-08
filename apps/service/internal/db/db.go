package database

import (
	"fmt"
	"log"
	"os"

	"github.com/fuu3629/odachin/apps/service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}
	return db
}

func Migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.Family{}, &models.User{}, &models.Wallet{}, &models.Transaction{}, &models.Allowance{}, &models.Reward{}, &models.Invitation{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("------------------------------")
		fmt.Println("")
		fmt.Println("Database migrated successfully🏃🏿‍♀️")
		fmt.Println("")
		fmt.Println("------------------------------")
	}
}
