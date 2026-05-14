package response

import "time"

// ==================== VERIFICATION RESPONSES ====================

// VerdictResponse represents the final AI verification verdict
type VerdictResponse struct {
	TrustScore    int                           `json:"trust_score"`
	Verdict       string                        `json:"verdict"`
	Breakdown     VerificationBreakdownResponse `json:"breakdown"`
	Flags         []string                      `json:"flags"`
	VerdictReason string                        `json:"verdict_reason"`
}

// VerificationBreakdownResponse details the scoring breakdown for each check
type VerificationBreakdownResponse struct {
	CACMatch            int `json:"cac_match"`
	NINMatch            int `json:"nin_match"`
	ImageAuthenticity   int `json:"image_authenticity"`
	ProductRecognition  int `json:"product_recognition"`
	LocationConsistency int `json:"location_consistency"`
	ReverseImageRisk    int `json:"reverse_image_risk"`
	SafeSearchScore     int `json:"safe_search_score"`
	FaceDetection       int `json:"face_detection"`
}

// PaymentInitiateResponse represents Squad payment initiation response
type PaymentInitiateResponse struct {
	CheckoutURL    string `json:"checkout_url"`
	TransactionRef string `json:"transaction_ref"`
	Amount         int64  `json:"amount"`
	Currency       string `json:"currency"`
}

// VerificationStatusResponse represents current verification processing status
type VerificationStatusResponse struct {
	JobID            string            `json:"job_id"`
	Status           string            `json:"status"`
	CurrentStep      string            `json:"current_step"`
	Steps            map[string]string `json:"steps"`
	EstimatedSeconds int               `json:"estimated_seconds"`
}

// DocumentUploadResponse represents successful document upload
type DocumentUploadResponse struct {
	Status        string `json:"status"`
	NextStep      string `json:"next_step"`
	UploadedCount int    `json:"uploaded_count"`
}

// ==================== VENDOR RESPONSES ====================

// VendorProfileResponse represents complete vendor profile with verification details
type VendorProfileResponse struct {
	VendorID             string `json:"vendor_id"`
	FullName             string `json:"full_name"`
	BusinessName         string `json:"business_name"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	BusinessType         string `json:"business_type"`
	Address              string `json:"address"`
	State                string `json:"state"`
	Status               string `json:"status"`
	TrustScore           int    `json:"trust_score"`
	Verdict              string `json:"verdict"`
	VerifiedBadge        bool   `json:"verified_badge"`
	VirtualAccountNumber string `json:"virtual_account_number"`
	Balance              int    `json:"balance"`
	TotalOrders          int    `json:"total_orders"`
}

// TrustSummary represents vendor trust information for customers
type TrustSummary struct {
	TrustScore int      `json:"trust_score"`
	Verdict    string   `json:"verdict"`
	Highlights []string `json:"highlights"`
}

// ==================== ADMIN RESPONSES ====================

// AdminVendorListResponse represents paginated vendor list for admin dashboard
type AdminVendorListResponse struct {
	Vendors []*AdminVendorItemResponse `json:"vendors"`
	Total   int                        `json:"total"`
	Page    int                        `json:"page"`
}

// AdminVendorItemResponse represents single vendor in admin list
type AdminVendorItemResponse struct {
	VendorID             string    `json:"vendor_id"`
	BusinessName         string    `json:"business_name"`
	Email                string    `json:"email"`
	TrustScore           int       `json:"trust_score"`
	Verdict              string    `json:"verdict"`
	Status               string    `json:"status"`
	SubmittedAt          time.Time `json:"submitted_at"`
	RequiresManualReview bool      `json:"requires_manual_review"`
}

// AdminMetricsResponse represents platform-wide metrics for admin dashboard
type AdminMetricsResponse struct {
	TotalVendors        int     `json:"total_vendors"`
	Approved            int     `json:"approved"`
	Restricted          int     `json:"restricted"`
	Flagged             int     `json:"flagged"`
	PendingManualReview int     `json:"pending_manual_review"`
	AvgTrustScore       float64 `json:"avg_trust_score"`
	FraudAttempts       int     `json:"fraud_attempts"`
}

// AdminVendorDetailResponse represents full vendor details for admin review
type AdminVendorDetailResponse struct {
	Vendor       VendorProfileResponse      `json:"vendor"`
	Verification AdminVerificationDetail    `json:"verification"`
	Documents    map[string]*DocumentDetail `json:"documents"`
}

// AdminVerificationDetail represents detailed verification info for admin
type AdminVerificationDetail struct {
	VerificationID   int                           `json:"verification_id"`
	TrustScore       int                           `json:"trust_score"`
	Verdict          string                        `json:"verdict"`
	Breakdown        VerificationBreakdownResponse `json:"breakdown"`
	Flags            []string                      `json:"flags"`
	ValidationStatus string                        `json:"validation_status"`
	ValidatedAt      *time.Time                    `json:"validated_at,omitempty"`
	AdminNote        *string                       `json:"admin_note,omitempty"`
}

// DocumentDetail represents document information for admin review
type DocumentDetail struct {
	DocumentID         string     `json:"document_id"`
	DocumentType       string     `json:"document_type"`
	VerificationStatus string     `json:"verification_status"`
	SignedURL          string     `json:"signed_url"`
	VerifiedAt         *time.Time `json:"verified_at,omitempty"`
}
