package request

// ==================== VENDOR PROFILE ====================

// UpdateVendorProfileRequest represents vendor profile update
type UpdateVendorProfileRequest struct {
	BusinessName string `form:"business_name" json:"business_name" name:"Business Name"`
	Phone        string `form:"phone" json:"phone" name:"Phone Number"`
	Address      string `form:"address" json:"address" name:"Address"`
	State        string `form:"state" json:"state" name:"State"`
}

// ==================== DOCUMENTS & KYC ====================

// UploadDocumentsRequest represents CAC, ID, and bank details upload
type UploadDocumentsRequest struct {
	DocumentType string `form:"document_type" json:"document_type" name:"Document Type"`
	// cacCertificate, validId, bankDetails files handled via multipart/form-data
	NINNumber     string `form:"nin_number" json:"nin_number" name:"NIN Number"`
	BankName      string `form:"bank_name" json:"bank_name" name:"Bank Name"`
	AccountNumber string `form:"account_number" json:"account_number" name:"Account Number"`
	AccountName   string `form:"account_name" json:"account_name" name:"Account Name"`
}

// ProofOfLifeRequest represents kitchen/meal photos and GPS location submission
type ProofOfLifeRequest struct {
	NINNumber        string  `form:"nin_number" json:"nin_number" name:"NIN Number"`
	BankName         string  `form:"bank_name" json:"bank_name" name:"Bank Name"`
	AccountNumber    string  `form:"account_number" json:"account_number" name:"Account Number"`
	AccountName      string  `form:"account_name" json:"account_name" name:"Account Name"`
	Latitude         float64 `form:"latitude" json:"latitude" name:"Latitude"`
	Longitude        float64 `form:"longitude" json:"longitude" name:"Longitude"`
	LocationAccuracy float64 `form:"location_accuracy" json:"location_accuracy" name:"Location Accuracy"`
	// productPhotos, exteriorPhoto, selfiePhoto handled as multipart files
}

// ==================== PAYMENT ====================

// InitiateVerificationFeeRequest initiates ₦1,000 Squad verification fee payment
type InitiateVerificationFeeRequest struct {
	VendorID int64 `form:"vendor_id" json:"vendor_id" name:"Vendor ID"`
}

// ==================== VERIFICATION PROCESSING ====================

// TriggerVerificationRequest triggers async AI verification pipeline
type TriggerVerificationRequest struct {
	VendorID int64 `form:"vendor_id" json:"vendor_id" name:"Vendor ID"`
}

// ==================== ADMIN ====================

// AdminLoginRequest represents admin staff login
type AdminLoginRequest struct {
	Email    string `form:"email" json:"email" name:"Email Address"`
	Password string `form:"password" json:"password" name:"Password"`
}

// AdminDecisionRequest represents manual verification decision by admin
type AdminDecisionRequest struct {
	Decision   string `form:"decision" json:"decision" name:"Decision"` // "approved", "restricted", "flagged", "suspended"
	AdminNote  string `form:"admin_note" json:"admin_note" name:"Admin Notes"`
	ReviewedBy string `form:"reviewed_by" json:"reviewed_by" name:"Reviewed By"`
}

// ListVendorsRequest represents vendor list query parameters for admin
type ListVendorsRequest struct {
	Status string `form:"status" json:"status" name:"Status"` // "all|pending|approved|restricted|flagged"
	Page   int    `form:"page" json:"page" name:"Page"`
	Limit  int    `form:"limit" json:"limit" name:"Limit"`
}
