package dojah

// ==================== LIVENESS CHECK ====================

// LivenessCheckRequest represents the request for liveness check API
type LivenessCheckRequest struct {
	Image []byte `json:"image"`
}

// LivenessCheckResponse represents the response from liveness check endpoint
type LivenessCheckResponse struct {
	Entity LivenessCheckEntity `json:"entity"`
}

// LivenessCheckEntity contains the liveness check data
type LivenessCheckEntity struct {
	Liveness LivenessData `json:"liveness"`
	Face     FaceData     `json:"face"`
}

// LivenessData contains liveness detection results
type LivenessData struct {
	LivenessCheck       bool    `json:"liveness_check"`
	LivenessProbability float64 `json:"liveness_probability"`
}

// FaceData contains face detection results
type FaceData struct {
	FaceDetected      bool        `json:"face_detected"`
	Message           string      `json:"message"`
	MultifaceDetected bool        `json:"multiface_detected"`
	Details           FaceDetails `json:"details"`
	Quality           FaceQuality `json:"quality"`
	Confidence        float64     `json:"confidence"`
	BoundingBox       BoundingBox `json:"bounding_box"`
}

// FaceDetails contains detailed face attributes
type FaceDetails struct {
	AgeRange   AgeRange      `json:"age_range"`
	Smile      Attribute     `json:"smile"`
	Gender     Attribute     `json:"gender"`
	Eyeglasses Attribute     `json:"eyeglasses"`
	Sunglasses Attribute     `json:"sunglasses"`
	Beard      Attribute     `json:"beard"`
	Mustache   Attribute     `json:"mustache"`
	EyesOpen   Attribute     `json:"eyes_open"`
	MouthOpen  Attribute     `json:"mouth_open"`
	Emotions   []EmotionData `json:"emotions"`
}

// AgeRange represents estimated age range
type AgeRange struct {
	Low  int `json:"low"`
	High int `json:"high"`
}

// Attribute represents a face attribute with value and confidence
type Attribute struct {
	Value      any     `json:"value"`
	Confidence float64 `json:"confidence"`
}

// EmotionData represents emotion detection result
type EmotionData struct {
	Type       string  `json:"type"`
	Confidence float64 `json:"confidence"`
}

// FaceQuality contains image quality metrics
type FaceQuality struct {
	Brightness float64 `json:"brightness"`
	Sharpness  float64 `json:"sharpness"`
}

// BoundingBox represents face bounding box coordinates
type BoundingBox struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Left   float64 `json:"left"`
	Top    float64 `json:"top"`
}

// ==================== NIN LOOKUP ====================

// NINLookupResponse represents the response from NIN lookup endpoint
type NINLookupResponse struct {
	Entity NINData `json:"entity"`
}

// NINData contains basic NIN holder information
type NINData struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	MiddleName       string `json:"middle_name,omitempty"`
	Gender           string `json:"gender"`
	Photo            string `json:"photo"`
	DateOfBirth      string `json:"date_of_birth"`
	PhoneNumber      string `json:"phone_number"`
	EmploymentStatus string `json:"employment_status"`
	MaritalStatus    string `json:"marital_status"`
}

// NINAdvancedLookupResponse represents the response from NIN advanced lookup
type NINAdvancedLookupResponse struct {
	Entity NINAdvancedData `json:"entity"`
}

// NINAdvancedData contains detailed NIN holder information
type NINAdvancedData struct {
	NIN                   string `json:"nin"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	MiddleName            string `json:"middle_name"`
	DateOfBirth           string `json:"date_of_birth"`
	PhoneNumber           string `json:"phone_number"`
	Photo                 string `json:"photo"`
	Gender                string `json:"gender"`
	EmploymentStatus      string `json:"employment_status"`
	MaritalStatus         string `json:"marital_status"`
	BirthCountry          string `json:"birth_country"`
	BirthLGA              string `json:"birth_lga"`
	BirthState            string `json:"birth_state"`
	EducationalLevel      string `json:"educational_level"`
	MaidenName            string `json:"maiden_name"`
	NspokenLang           string `json:"nspoken_lang"`
	Profession            string `json:"profession"`
	Religion              string `json:"religion"`
	ResidenceAddressLine1 string `json:"residence_address_line_1"`
	ResidenceAddressLine2 string `json:"residence_address_line_2"`
	ResidenceStatus       string `json:"residence_status"`
	ResidenceTown         string `json:"residence_town"`
	ResidenceLGA          string `json:"residence_lga"`
	ResidenceState        string `json:"residence_state"`
	OspokenLang           string `json:"ospoken_lang"`
	OriginLGA             string `json:"origin_lga"`
	OriginPlace           string `json:"origin_place"`
	OriginState           string `json:"origin_state"`
	Height                string `json:"height"`
	PFirstName            string `json:"p_first_name"`
	PMiddleName           string `json:"p_middle_name"`
	PLastName             string `json:"p_last_name"`
	TaxID                 string `json:"tax_id"`
	TaxResidency          string `json:"tax_residency"`
}

// ==================== NIN VERIFICATION WITH SELFIE ====================

// VerifyNINWithSelfieRequest represents the request to verify NIN with selfie
type VerifyNINWithSelfieRequest struct {
	NIN         string `json:"nin"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	SelfieImage []byte `json:"selfie_image"`
}

// VerifyNINWithSelfieResponse represents the response from NIN selfie verification
type VerifyNINWithSelfieResponse struct {
	Entity VerifyNINWithSelfieData `json:"entity"`
}

// VerifyNINWithSelfieData contains NIN verification result with selfie matching
type VerifyNINWithSelfieData struct {
	FirstName          string                  `json:"first_name"`
	LastName           string                  `json:"last_name"`
	MiddleName         string                  `json:"middle_name"`
	Gender             string                  `json:"gender"`
	Image              string                  `json:"image"`
	PhoneNumber        string                  `json:"phone_number"`
	DateOfBirth        string                  `json:"date_of_birth"`
	NIN                string                  `json:"nin"`
	SelfieVerification SelfieVerificationMatch `json:"selfie_verification"`
}

// SelfieVerificationMatch contains selfie verification results
type SelfieVerificationMatch struct {
	ConfidenceValue float64 `json:"confidence_value"`
	Match           bool    `json:"match"`
}

// ==================== CAC VERIFICATION ====================

// CACVerificationRequest represents the request to verify CAC (Corporate Affairs Commission)
type CACVerificationRequest struct {
	CACNumber    string `json:"cac_number"`
	BusinessName string `json:"business_name,omitempty"`
}

// CACVerificationResponse represents the response from CAC verification
type CACVerificationResponse struct {
	Entity CACVerificationData `json:"entity"`
}

// CACVerificationData contains CAC verification results
type CACVerificationData struct {
	Status             string                   `json:"status"`
	Message            string                   `json:"message"`
	BusinessInfo       CACBusinessInfo          `json:"business_info"`
	VerificationResult VerificationResultStatus `json:"verification_result,omitempty"`
}

// CACBusinessInfo contains CAC business information
type CACBusinessInfo struct {
	BusinessName       string   `json:"business_name"`
	RegistrationNumber string   `json:"registration_number"`
	BusinessType       string   `json:"business_type"`
	RegistrationDate   string   `json:"registration_date"`
	State              string   `json:"state"`
	LGA                string   `json:"lga"`
	BusinessAddress    string   `json:"business_address"`
	Directors          []string `json:"directors,omitempty"`
	Shareholders       []string `json:"shareholders,omitempty"`
}

// VerificationResultStatus contains verification status
type VerificationResultStatus struct {
	IsValid bool   `json:"is_valid"`
	Reason  string `json:"reason,omitempty"`
}

// ==================== DOCUMENT ANALYSIS ====================

// DocumentAnalysisRequest represents the request for document analysis
type DocumentAnalysisRequest struct {
	InputType      string `json:"input_type"`
	ImageFrontSide []byte `json:"imagefrontside"`
	ImageBackSide  []byte `json:"imagebackside,omitempty"`
	Images         []byte `json:"images,omitempty"`
}

// DocumentAnalysisResponse represents the response from document analysis endpoint
type DocumentAnalysisResponse struct {
	Entity DocumentAnalysisEntity `json:"entity"`
}

// DocumentAnalysisEntity contains document analysis data
type DocumentAnalysisEntity struct {
	Status         DocumentStatus   `json:"status"`
	DocumentType   DocumentTypeInfo `json:"document_type"`
	DocumentImages DocumentImages   `json:"document_images"`
	TextData       []TextDataField  `json:"text_data"`
}

// DocumentStatus contains document validation status
type DocumentStatus struct {
	OverallStatus  int    `json:"overall_status"`
	Reason         string `json:"reason"`
	DocumentImages string `json:"document_images"`
	Text           string `json:"text"`
	DocumentType   string `json:"document_type"`
	Expiry         string `json:"expiry"`
}

// DocumentTypeInfo contains identified document type details
type DocumentTypeInfo struct {
	DocumentName        string `json:"document_name"`
	DocumentCountryName string `json:"document_country_name"`
	DocumentCountryCode string `json:"document_country_code"`
}

// DocumentImages contains extracted images from the document
type DocumentImages struct {
	Portrait          string `json:"portrait"`
	DocumentFrontSide string `json:"document_front_side"`
	DocumentBackSide  string `json:"document_back_side"`
}

// TextDataField represents an extracted text field
type TextDataField struct {
	FieldName string `json:"field_name"`
	FieldKey  string `json:"field_key"`
	Status    int    `json:"status"`
	Value     string `json:"value"`
}

// ==================== GET VERIFICATION DETAILS ====================

// GetVerificationDetailsResponse represents the response for getting verification details
type GetVerificationDetailsResponse struct {
	AML                any                 `json:"aml"`
	Data               VerificationData    `json:"data"`
	Value              string              `json:"value"`
	IDUrl              string              `json:"id_url"`
	Status             bool                `json:"status"`
	IDType             string              `json:"id_type"`
	Message            string              `json:"message"`
	BackUrl            string              `json:"back_url"`
	Metadata           VerificationMetadata `json:"metadata"`
	SelfieUrl          string              `json:"selfie_url"`
	ReferenceID        string              `json:"reference_id"`
	VerificationUrl    string              `json:"verification_url"`
	VerificationMode   string              `json:"verification_mode"`
	VerificationType   string              `json:"verification_type"`
	VerificationValue  string              `json:"verification_value"`
	VerificationStatus string              `json:"verification_status"`
}

// VerificationData contains all verification step results
type VerificationData struct {
	ID                 any   `json:"id"`
	Email              any   `json:"email"`
	Selfie             any   `json:"selfie"`
	Countries          any   `json:"countries"`
	UserData           any   `json:"user_data"`
	BusinessID         any   `json:"business_id"`
	PhoneNumber        any   `json:"phone_number"`
	BusinessData       any   `json:"business_data"`
	GovernmentData     any   `json:"government_data"`
	AdditionalDocument []any `json:"additional_document"`
}

// VerificationMetadata contains metadata about the verification session
type VerificationMetadata struct {
	IPInfo     IPInfo `json:"ipinfo"`
	DeviceInfo string `json:"device_info"`
}

// IPInfo contains IP geolocation data
type IPInfo struct {
	AS         string  `json:"as"`
	ISP        string  `json:"isp"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	Org        string  `json:"org"`
	Zip        string  `json:"zip"`
	City       string  `json:"city"`
	Proxy      bool    `json:"proxy"`
	Query      string  `json:"query"`
	Mobile     bool    `json:"mobile"`
	Status     string  `json:"status"`
	Country    string  `json:"country"`
	Hosting    bool    `json:"hosting"`
	District   string  `json:"district"`
	Timezone   string  `json:"timezone"`
	RegionName string  `json:"region_name"`
}

// ==================== LIST ALL VERIFICATIONS ====================

// ListAllVerificationsResponse represents the response for listing all verifications
type ListAllVerificationsResponse struct {
	Entity ListAllVerificationsEntity `json:"entity"`
}

// ListAllVerificationsEntity contains paginated verification data
type ListAllVerificationsEntity struct {
	Data []VerificationRecord `json:"data"`
	Meta PaginationMeta       `json:"meta"`
}

// VerificationRecord represents a single verification record
type VerificationRecord struct {
	ReferenceID        string `json:"reference_id"`
	AppID              string `json:"app_id"`
	VerificationStatus string `json:"verificationStatus"`
	DateTime           string `json:"datetime"`
	Environment        string `json:"environment"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	MiddleName         string `json:"middle_name"`
	FullName           string `json:"full_name"`
	BusinessName       string `json:"business_name"`
	SelfieUrl          string `json:"selfieUrl"`
	VerificationUrl    string `json:"verificationUrl"`
}

// PaginationMeta contains pagination information
type PaginationMeta struct {
	TotalCount  int `json:"total_count"`
	ItemPerPage int `json:"item_per_page"`
	CurrentPage int `json:"current_page"`
}

// ==================== WEBHOOK VERIFICATION ====================

// WebhookVerificationRequest represents webhook event verification
type WebhookVerificationRequest struct {
	SignatureHeader string
	SecretKey       string
	Payload         []byte
}

// WebhookVerificationResult contains the result of webhook verification
type WebhookVerificationResult struct {
	IsValid bool
	Error   error
}
