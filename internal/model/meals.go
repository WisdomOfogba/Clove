package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// ==================== MEALS & MENU ====================

// Meal represents a meal offered by a vendor
type Meal struct {
	ID          int64     `json:"meal_id"`
	VendorID    int64     `json:"vendor_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`    // in kobo
	Category    string    `json:"category"` // "main_course", "appetizer", "dessert", "beverage"
	Status      string    `json:"status"`   // "active", "inactive"
	Score       float64   `json:"score"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MealPicture represents pictures for a meal (multiple per meal)
type MealPicture struct {
	ID        int64     `json:"picture_id"`
	MealID    int64     `json:"meal_id"`
	ImageURL  string    `json:"image_url"`
	PublicID  string    `json:"-"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
}

// ==================== ORDERS ====================

// Order represents a customer order
type Order struct {
	ID              string          `json:"order_id"`
	CustomerID      int64           `json:"customer_id"`
	VendorID        int64           `json:"vendor_id"`
	Status          string          `json:"status"`       // "pending", "confirmed", "preparing", "ready", "delivered", "cancelled"
	TotalAmount     decimal.Decimal `json:"total_amount"` // in kobo
	DeliveryFee     decimal.Decimal `json:"delivery_fee"` // in kobo
	DeliveryAddress string          `json:"delivery_address"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

// OrderItem represents items in an order
type OrderItem struct {
	ID        int64           `json:"id"`
	OrderID   string          `json:"order_id"`
	MealID    int64           `json:"meal_id"`
	Quantity  int             `json:"quantity"`
	Price     decimal.Decimal `json:"price"` // per item in kobo
	CreatedAt time.Time       `json:"created_at"`
}

// ==================== REVIEWS & RATINGS ====================

// Review represents customer review for a meal/order (only buyers can review)
type Review struct {
	ID             int64     `json:"id"`
	CustomerID     int64     `json:"customer_id"`
	VendorID       int64     `json:"vendor_id"`
	MealID         int64     `json:"meal_id"`
	Rating         int64     `json:"rating"` // 1-5
	Comment        string    `json:"comment"`
	Edits          int16     `json:"edits"`
	Sentiment      *string   `json:"sentiment"`
	SentimentScore *float64  `json:"sentiment_score"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
