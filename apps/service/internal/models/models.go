package models

import "time"

type Family struct {
	FamilyID   uint `gorm:"primaryKey;autoIncrement"`
	FamilyName string
	Email      string `gorm:"unique"`
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Users      []User `gorm:"foreignKey:FamilyID"`
}

type User struct {
	UserID    string `gorm:"primaryKey"`
	FamilyID  *uint  `gorm:"index"`
	Role      string
	UserName  string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Wallet    Wallet `gorm:"foreignKey:UserID"`
}

type Wallet struct {
	WalletID  uint   `gorm:"primaryKey;autoIncrement"`
	UserID    string `gorm:"uniqueIndex"`
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Transaction struct {
	TransactionID uint `gorm:"primaryKey"`
	FromUserID    uint `gorm:"index"`
	ToUserID      uint `gorm:"index"`
	Amount        float64
	Type          string
	CreatedAt     time.Time
}

type Allowance struct {
	AllowanceID uint `gorm:"primaryKey"`
	FromUserID  uint `gorm:"index"`
	ToUserID    uint `gorm:"index"`
	Amount      float64
	Interval    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Reward struct {
	RewardID  uint `gorm:"primaryKey"`
	ToUserID  uint `gorm:"index"`
	Amount    float64
	Reason    string
	CreatedAt time.Time
}

// func main() {
// 	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	db.AutoMigrate(&Family{}, &User{}, &Wallet{}, &Transaction{}, &Allowance{}, &Reward{})
// }
