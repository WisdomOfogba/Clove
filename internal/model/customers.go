package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// User represents a customer/user in the system
type User struct {
	ID              int64     `json:"user_id"`
	FullName        string    `json:"full_name"`
	Email           string    `json:"email"`
	PhoneNumber     string    `json:"phone"`
	Password        string    `json:"password"`
	Status          string    `json:"status"` // "active", "suspended"
	IsEmailVerified bool      `json:"is_email_verified"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type UserSession struct {
	UserId int64 `gorm:"not null" redis:"user_id"`
	// UserId           uint         `gorm:"not null" redis:"user_id"`
	RefreshTokenHash string    `gorm:"not null" redis:"-"`
	LastIP           string    `gorm:"" redis:"last_ip"`
	DeviceId         uuid.UUID `gorm:"" redis:"device_id"`
	UserAgent        string    `gorm:"not null" redis:"user_agent"`
	CreatedAt        time.Time `gorm:"not null" redis:"created_at"`
	ExpiresAt        time.Time `gorm:"not null" redis:"created_at"`
}

type UserAddress struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Address string `json:"address"`
}

type UserWallet struct {
	UserID       int64           `json:"user_id"`
	Amount       decimal.Decimal `json:"amount"`
	LastFundedAt *time.Time      `json:"last_funded_at"`
}

// On insert do nothing (for conflict)
type UserPurchasedMeal struct {
	UserID int64 `json:"user_id"`
	MealID int64 `json:"meal_id"`
}
