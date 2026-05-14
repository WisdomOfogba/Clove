package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// ==================== VENDOR CORE ====================

// Vendor represents a business/vendor in the system
type Business struct {
	BusinessID   int64     `json:"business_id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	TIN          string    `json:"tin"`
	BusinessType string    `json:"business_type"` // "restaurant", "ecommerce", "food_delivery"
	RCNumber     string    `json:"rc_number"`     // CAC/RC number
	NIN          string    `json:"nin"`           // Unique NIN for one store per vendor
	State        string    `json:"state"`
	Address      string    `json:"address"`
	Status       string    `json:"status"` // "pending_documents", "documents_received", "payment_pending", "processing", "approved", "restricted", "flagged", "suspended"
	VendorID     int64     `json:"vendor_id"`
	CreatedAt    time.Time `json:"created_at"`
}

// ==================== DOCUMENTS & KYC ====================

// VendorDocument represents documents submitted for verification
type VendorDocument struct {
	DocumentID         string     `json:"document_id"`
	VendorID           int64      `json:"vendor_id"`
	DocumentType       string     `json:"document_type"` // "cac_certificate", "valid_id", "bank_details", "business_photo", "product_photo", "exterior_photo", "selfie_photo"
	DocumentName       string     `json:"document_name"`
	FilePath           string     `json:"file_path"`
	VerificationStatus string     `json:"verification_status"` // "pending", "verified", "rejected"
	VerifiedBy         *string    `json:"verified_by,omitempty"`
	VerifiedAt         *time.Time `json:"verified_at,omitempty"`
	ExpiryDate         *time.Time `json:"expiry_date,omitempty"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

// VendorKYC stores personal KYC data for vendor
type VendorKYC struct {
	KYCID            string    `json:"kyc_id"`
	VendorID         int64     `json:"vendor_id"`
	NINNumber        string    `json:"nin_number"`
	BankName         string    `json:"bank_name"`
	AccountNumber    string    `json:"account_number"`
	AccountName      string    `json:"account_name"`
	Latitude         *float64  `json:"latitude,omitempty"`
	Longitude        *float64  `json:"longitude,omitempty"`
	LocationAccuracy *float64  `json:"location_accuracy,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// ==================== VERIFICATION & SCORING ====================

// VendorVerification represents the main verification record with AI scoring
type VendorVerification struct {
	VerificationID   string     `json:"verification_id"`
	VendorID         int64      `json:"vendor_id"`
	TrustScore       float64    `json:"trust_score"`    // 0-100, updated based on interactions
	InitialScore     float64    `json:"initial_score"`  // Initial score granted after registration
	Verdict          string     `json:"verdict"`        // "approved", "restricted", "flagged"
	BreakdownJSON    string     `json:"breakdown_json"` // JSON-encoded breakdown
	Flags            string     `json:"flags"`          // JSON-encoded array
	VerdictReason    string     `json:"verdict_reason"`
	ValidationStatus string     `json:"validation_status"` // "pending", "processing", "complete", "failed"
	JobID            string     `json:"job_id,omitempty"`
	ValidatedBy      *string    `json:"validated_by,omitempty"`
	ValidatedAt      *time.Time `json:"validated_at,omitempty"`
	AdminNote        *string    `json:"admin_note,omitempty"`
	AdminDecision    *string    `json:"admin_decision,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

// VerificationJob tracks async verification processing
type VerificationJob struct {
	JobID            string    `json:"job_id"`
	VendorID         int64     `json:"vendor_id"`
	Status           string    `json:"status"`       // "pending", "processing", "complete", "failed"
	CurrentStep      string    `json:"current_step"` // "cac_check", "nin_check", "image_analysis", "score_fusion"
	StepsJSON        string    `json:"steps_json"`   // JSON-encoded step statuses
	ErrorMessage     *string   `json:"error_message,omitempty"`
	Result           *string   `json:"result,omitempty"`
	EstimatedSeconds int       `json:"estimated_seconds"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// ==================== PAYMENT & TRANSACTIONS ====================

// PaymentTransaction tracks verification fee payments
type PaymentTransaction struct {
	TransactionID       string          `json:"transaction_id"`
	VendorID            int64           `json:"vendor_id"`
	TransactionRef      string          `json:"transaction_ref"`
	Amount              decimal.Decimal `json:"amount"`           // in kobo
	Currency            string          `json:"currency"`         // "NGN"
	Status              string          `json:"status"`           // "pending", "success", "failed"
	TransactionType     string          `json:"transaction_type"` // "verification_fee", "payout"
	SquadCheckoutURL    *string         `json:"squad_checkout_url,omitempty"`
	SquadTransactionRef *string         `json:"squad_transaction_ref,omitempty"`
	CreatedAt           time.Time       `json:"created_at"`
	UpdatedAt           time.Time       `json:"updated_at"`
}

// VendorVirtualAccount stores Squad virtual account details
type VendorVirtualAccount struct {
	VirtualAccountID     string    `json:"virtual_account_id"`
	VendorID             int64     `json:"vendor_id"`
	VirtualAccountNumber string    `json:"virtual_account_number"`
	BankName             string    `json:"bank_name"`
	CustomerIdentifier   string    `json:"customer_identifier"`
	Status               string    `json:"status"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

// ==================== ADMIN & USERS ====================

// AdminUser represents admin staff
type AdminUser struct {
	AdminID   int64     `json:"admin_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ==================== METRICS & ANALYTICS ====================

// VendorMetrics stores aggregated metrics for a vendor
type VendorMetrics struct {
	MetricsID            string    `json:"metrics_id"`
	VendorID             int64     `json:"vendor_id"`
	TotalRatings         int       `json:"total_ratings"`
	AverageRating        float64   `json:"average_rating"`
	TotalComments        int       `json:"total_comments"`
	PositiveComments     int       `json:"positive_comments"`
	NegativeComments     int       `json:"negative_comments"`
	PositiveCommentRatio float64   `json:"positive_comment_ratio"`
	ResponseRate         float64   `json:"response_rate"`
	AverageResponseTime  float64   `json:"average_response_time"`
	TotalOrders          int       `json:"total_orders"`
	OrderFulfillmentRate float64   `json:"order_fulfillment_rate"`
	LastUpdated          time.Time `json:"last_updated"`
}

// PlatformMetrics tracks overall platform analytics
type PlatformMetrics struct {
	MetricsID           string    `json:"metrics_id" gorm:"column:metrics_id;primaryKey"`
	TotalVendors        int       `json:"total_vendors" gorm:"column:total_vendors;not null;default:0"`
	ApprovedVendors     int       `json:"approved_vendors" gorm:"column:approved_vendors;not null;default:0"`
	RestrictedVendors   int       `json:"restricted_vendors" gorm:"column:restricted_vendors;not null;default:0"`
	FlaggedVendors      int       `json:"flagged_vendors" gorm:"column:flagged_vendors;not null;default:0"`
	PendingManualReview int       `json:"pending_manual_review" gorm:"column:pending_manual_review;not null;default:0"`
	AvgTrustScore       float64   `json:"avg_trust_score" gorm:"column:avg_trust_score;not null;default:0"`
	FraudAttempts       int       `json:"fraud_attempts" gorm:"column:fraud_attempts;not null;default:0"`
	TotalOrders         int       `json:"total_orders" gorm:"column:total_orders;not null;default:0"`
	TotalRevenue        int64     `json:"total_revenue" gorm:"column:total_revenue;not null;default:0"`
	LastUpdated         time.Time `json:"last_updated" gorm:"column:last_updated;not null"`
}
