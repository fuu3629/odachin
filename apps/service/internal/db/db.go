package database

import (
	"fmt"
	"log"
	"os"

	"github.com/fuu3629/odachin/apps/service/internal/models"
	"github.com/fuu3629/odachin/apps/service/pkg/assets"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Tokyo",
		host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}
	return db
}

func Migrations(db *gorm.DB) {
	db.Exec(`
DO
$$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
    create type role_enum AS ENUM ('CHILD', 'PARENT');
  END IF;
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'interval_enum') THEN
		create type interval_enum AS ENUM ('DAILY', 'MONTHLY', 'WEEKLY');
	END IF;
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'dayofweek_enum') THEN
		create type dayofweek_enum AS ENUM ('MONDAY', 'TUESDAY', 'WEDNESDAY', 'THURSDAY', 'FRIDAY', 'SATURDAY', 'SUNDAY');
	END IF;
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'period_enum') THEN
		create type period_enum AS ENUM ('DAILY', 'WEEKLY', 'MONTHLY');
	END IF;
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'reward_period_status_enum') THEN
		create type reward_period_status_enum AS ENUM ('IN_PROGRESS', 'REPORTED', 'COMPLETED', 'REJECTED');
	END IF;
END
$$;
`)
	err := db.AutoMigrate(&models.Family{}, &models.User{}, &models.Wallet{}, &models.Transaction{}, &models.Allowance{}, &models.Reward{}, &models.RewardPeriod{}, &models.Invitation{})
	if err != nil {
		fmt.Println(err)
	} else {
		assets.Log("Database migrated successfully🕊️")
	}
}
