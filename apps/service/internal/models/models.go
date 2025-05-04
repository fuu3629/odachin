package models

import "time"

type Family struct {
	FamilyID   uint `gorm:"primaryKey;autoIncrement"`
	FamilyName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Users      []User `gorm:"foreignKey:FamilyID"`
}

type User struct {
	UserID         string `gorm:"primaryKey"`
	FamilyID       *uint  `gorm:"index"`
	Role           string `gorm:"type: role_enum;"`
	UserName       string
	Email          string `gorm:"unique"`
	Password       string
	AvatarImageUrl *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Wallet         Wallet    `gorm:"foreignKey:UserID"`
	Allowance      Allowance `gorm:"foreignKey:ToUserID;constraint:OnDelete:CASCADE;"`
}

type Wallet struct {
	WalletID  uint   `gorm:"primaryKey;autoIncrement"`
	UserID    string `gorm:"uniqueIndex"`
	Balance   int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TODO typeをenumにする
type Transaction struct {
	TransactionID uint   `gorm:"primaryKey;autoIncrement"`
	FromUserID    string `gorm:"index"`
	ToUserID      string `gorm:"index"`
	Amount        int32
	Type          string
	Title         string
	Description   string
	CreatedAt     time.Time
}

// Monthlyは毎月Dateに送金される
// Weeklyは毎週DayOfWeekに送金される
type Allowance struct {
	AllowanceID  uint   `gorm:"primaryKey;autoIncrement"`
	FromUserID   string `gorm:"index"`
	ToUserID     string `gorm:"index"`
	Amount       int32
	IntervalType string `gorm:"type: period_enum"`
	Date         *uint32
	DayOfWeek    *string `gorm:"type: dayofweek_enum"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Reward struct {
	RewardID      uint   `gorm:"primaryKey;autoIncrement"`
	FromUserID    string `gorm:"index"`
	ToUserID      string `gorm:"index"`
	PeriodType    string `gorm:"type: period_enum"`
	Title         string
	Description   string
	Amount        int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	RewardPeriods []RewardPeriod `gorm:"foreignKey:RewardID;constraint:OnDelete:CASCADE;"`
}

type RewardPeriod struct {
	RewardPeriodID uint `gorm:"primaryKey;autoIncrement"`
	RewardID       uint
	StartDate      time.Time
	EndDate        time.Time
	Status         string `gorm:"type: reward_period_status_enum"`
	Reward         Reward
}

type Invitation struct {
	InvitationID uint   `gorm:"primaryKey;autoIncrement"`
	FamilyID     *uint  `gorm:"index"`
	FromUserID   string `gorm:"index:idx_user_pair,unique"`
	ToUserID     string `gorm:"index:idx_user_pair,unique"`
	IsAccepted   bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Usage struct {
	UsageID     uint   `gorm:"primaryKey;autoIncrement"`
	UserID      string `gorm:"index"`
	Amount      int32
	Title       string
	Description string
	Category    string
	CreatedAt   time.Time
}
