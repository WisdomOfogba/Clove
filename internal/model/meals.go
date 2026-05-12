package model

import "time"

// ==================== MEALS & MENU ====================

// Meal represents a meal offered by a vendor
type Meal struct {
	MealID      string    `json:"meal_id"`
	VendorID    string    `json:"vendor_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`       // in kobo
	Category    string    `json:"category"`    // "main_course", "appetizer", "dessert", "beverage"
	Status      string    `json:"status"`      // "active", "inactive"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MealPicture represents pictures for a meal (multiple per meal)
type MealPicture struct {
	PictureID string    `json:"picture_id"`
	MealID    string    `json:"meal_id"`
	ImageURL  string    `json:"image_url"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
}

// ==================== CUSTOMERS ====================

// Customer represents a customer/user in the system
type Customer struct {
	CustomerID string    `json:"customer_id"`
	FullName   string    `json:"full_name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Password   string    `json:"password"`
	Status     string    `json:"status"` // "active", "suspended"
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ==================== ORDERS ====================

// Order represents a customer order
type Order struct {
	OrderID         string    `json:"order_id"`
	CustomerID      string    `json:"customer_id"`
	VendorID        string    `json:"vendor_id"`
	Status          string    `json:"status"`       // "pending", "confirmed", "preparing", "ready", "delivered", "cancelled"
	TotalAmount     int64     `json:"total_amount"` // in kobo
	DeliveryFee     int64     `json:"delivery_fee"` // in kobo
	DeliveryAddress string    `json:"delivery_address"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// OrderItem represents items in an order
type OrderItem struct {
	OrderItemID string    `json:"order_item_id"`
	OrderID     string    `json:"order_id"`
	MealID      string    `json:"meal_id"`
	Quantity    int       `json:"quantity"`
	Price       int64     `json:"price"` // per item in kobo
	CreatedAt   time.Time `json:"created_at"`
}

// ==================== REVIEWS & RATINGS ====================

// Review represents customer review for a meal/order (only buyers can review)
type Review struct {
	ReviewID   string    `json:"review_id"`
	OrderID    string    `json:"order_id"`
	CustomerID string    `json:"customer_id"`
	VendorID   string    `json:"vendor_id"`
	MealID     string    `json:"meal_id"`
	Rating     int       `json:"rating"`  // 1-5
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
}