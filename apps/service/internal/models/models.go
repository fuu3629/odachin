package models

type Family struct {
	FamilyID   uint `gorm:"primaryKey"`
	FamilyName string
	Email      string `gorm:"unique"`
	Password   string
	CreatedAt  int64
	UpdatedAt  int64
	Users      []User `gorm:"foreignKey:FamilyID"`
}

type User struct {
	UserID    uint `gorm:"primaryKey"`
	FamilyID  uint `gorm:"index"`
	Role      string
	UserName  string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt int64
	UpdatedAt int64
	Wallet    Wallet `gorm:"foreignKey:UserID"`
}

type Wallet struct {
	WalletID  uint `gorm:"primaryKey"`
	UserID    uint `gorm:"uniqueIndex"`
	Balance   float64
	CreatedAt int64
	UpdatedAt int64
}

type Transaction struct {
	TransactionID uint `gorm:"primaryKey"`
	FromUserID    uint `gorm:"index"`
	ToUserID      uint `gorm:"index"`
	Amount        float64
	Type          string
	CreatedAt     int64
}

type Allowance struct {
	AllowanceID uint `gorm:"primaryKey"`
	FromUserID  uint `gorm:"index"`
	ToUserID    uint `gorm:"index"`
	Amount      float64
	Interval    string
	CreatedAt   int64
	UpdatedAt   int64
}

type Reward struct {
	RewardID  uint `gorm:"primaryKey"`
	ToUserID  uint `gorm:"index"`
	Amount    float64
	Reason    string
	CreatedAt int64
}

// func main() {
// 	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	db.AutoMigrate(&Family{}, &User{}, &Wallet{}, &Transaction{}, &Allowance{}, &Reward{})
// }
