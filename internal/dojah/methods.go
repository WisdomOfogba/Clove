package dojah

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURLSandbox = "https://sandbox.dojah.io"
	baseURLLive    = "https://api.dojah.io"
)

// SetEnvironment sets the API environment (sandbox or live)
// If useProduction is true, live API is used; otherwise sandbox is used
func (dc *DojahClient) SetEnvironment(useProduction bool) {
	dc.isProduction = useProduction
	if useProduction {
		dc.baseURL = baseURLLive
		return
	}
	dc.baseURL = baseURLSandbox
}

func (dc *DojahClient) IsProduction() bool {
	return dc.isProduction
}

func (dc *DojahClient) Secret() string {
	return dc.secretKey
}

func (dc *DojahClient) AppID() string {
	return dc.appId
}

// ==================== INTERNAL HELPER METHODS ====================

// buildURL constructs the full URL for an endpoint
func (dc *DojahClient) buildURL(baseURL, endpoint string) string {
	return fmt.Sprintf("%s%s", baseURL, endpoint)
}

// sendRequest sends an HTTP request to the Dojah API
func (dc *DojahClient) sendRequest(method, endpoint string, payload any) (*http.Response, error) {
	requestURL := dc.buildURL(dc.baseURL, endpoint)

	var body io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, requestURL, body)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("AppId", dc.appId)
	req.Header.Set("Authorization", dc.secretKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := dc.httpclient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// sendQueryRequest sends an HTTP GET request with query parameters
func (dc *DojahClient) sendQueryRequest(endpoint string, params map[string]string) (*http.Response, error) {
	requestURL := dc.buildURL(dc.baseURL, endpoint)

	// Add query parameters
	if len(params) > 0 {
		q := url.Values{}
		for key, value := range params {
			q.Add(key, value)
		}
		requestURL = fmt.Sprintf("%s?%s", requestURL, q.Encode())
	}

	// Parse URL for the query parameters properly
	parsedURL, err := url.Parse(requestURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("AppId", dc.appId)
	req.Header.Set("Authorization", dc.secretKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := dc.httpclient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// decodeResponse decodes the response body into the given value
func decodeResponse(resp *http.Response, v any) error {
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, v); err != nil {
		return err
	}

	return nil
}

// ==================== LIVENESS CHECK ====================

// LivenessCheck performs liveness detection on a selfie image
func (dc *DojahClient) LivenessCheck(req *LivenessCheckRequest) (*LivenessCheckResponse, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if len(req.Image) == 0 {
		return nil, errors.New("image data is required")
	}

	// Convert image to base64
	base64Image, err := ConvertToBase64(req.Image)
	if err != nil {
		return nil, err
	}

	// Create payload with base64 string
	payload := map[string]string{
		"image": base64Image,
	}

	resp, err := dc.sendRequest("POST", "/api/v1/ml/liveness/", payload)
	if err != nil {
		return nil, err
	}

	var result LivenessCheckResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("liveness check failed with status code: %d", resp.StatusCode)
	}

	return &result, nil
}

// ==================== NIN LOOKUP ====================

// LookupNIN performs a basic NIN lookup
func (dc *DojahClient) LookupNIN(nin string) (*NINLookupResponse, error) {
	if nin == "" {
		return nil, errors.New("nin is required")
	}

	// In sandbox, use test NIN if needed
	if !dc.isProduction && nin == "" {
		nin = "70123456789" // Sandbox test NIN
	}

	params := map[string]string{
		"nin": nin,
	}

	resp, err := dc.sendQueryRequest("/api/v1/kyc/nin", params)
	if err != nil {
		return nil, err
	}

	var result NINLookupResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("nin lookup failed with status code: %d", resp.StatusCode)
	}

	return &result, nil
}

// LookupNINAdvanced performs an advanced NIN lookup with detailed information
func (dc *DojahClient) LookupNINAdvanced(nin string) (*NINAdvancedLookupResponse, error) {
	if nin == "" {
		return nil, errors.New("nin is required")
	}

	// In sandbox, use test NIN if needed
	if !dc.isProduction && nin == "" {
		nin = "70123456789" // Sandbox test NIN
	}

	params := map[string]string{
		"nin": nin,
	}

	resp, err := dc.sendQueryRequest("/api/v1/kyc/nin/advance", params)
	if err != nil {
		return nil, err
	}

	var result NINAdvancedLookupResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("advanced nin lookup failed with status code: %d", resp.StatusCode)
	}

	return &result, nil
}

// ==================== NIN VERIFICATION WITH SELFIE ====================

// VerifyNINWithSelfie verifies a NIN against a selfie image
func (dc *DojahClient) VerifyNINWithSelfie(req *VerifyNINWithSelfieRequest) (*VerifyNINWithSelfieResponse, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if req.NIN == "" {
		return nil, errors.New("nin is required")
	}

	if len(req.SelfieImage) == 0 {
		return nil, errors.New("selfie image is required")
	}

	// Convert selfie to base64
	base64Selfie, err := ConvertToBase64(req.SelfieImage)
	if err != nil {
		return nil, err
	}

	dataStr := base64Selfie
	if strings.HasPrefix(dataStr, "data:image/") {
		// Find the comma that separates the metadata from the base64 data
		_, after, ok := strings.Cut(dataStr, ",")
		if ok {
			// Return everything after the comma (the base64 data)
			base64Selfie = after
		}
	}

	payload := map[string]any{
		"nin":          req.NIN,
		"selfie_image": base64Selfie,
	}

	if req.FirstName != "" {
		payload["first_name"] = req.FirstName
	}
	if req.LastName != "" {
		payload["last_name"] = req.LastName
	}

	resp, err := dc.sendRequest("POST", "/api/v1/kyc/nin/verify", payload)
	if err != nil {
		return nil, err
	}

	var result VerifyNINWithSelfieResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("nin selfie verification failed with status code: %d", resp.StatusCode)
	}

	return &result, nil
}

// ==================== CAC VERIFICATION ====================

// VerifyCAC verifies a business using Corporate Affairs Commission number
func (dc *DojahClient) VerifyCAC(req *CACVerificationRequest) (*CACVerificationResponse, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if req.CACNumber == "" {
		return nil, errors.New("cac_number is required")
	}

	payload := map[string]any{
		"cac_number": req.CACNumber,
	}

	if req.BusinessName != "" {
		payload["business_name"] = req.BusinessName
	}

	resp, err := dc.sendRequest("POST", "/api/v1/kyc/cac", payload)
	if err != nil {
		return nil, err
	}

	var result CACVerificationResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cac verification failed with status code: %d", resp.StatusCode)
	}

	return &result, nil
}

// ==================== DOCUMENT ANALYSIS ====================

// AnalyzeDocument performs document analysis on ID documents
func (dc *DojahClient) AnalyzeDocument(req *DocumentAnalysisRequest) (*DocumentAnalysisResponse, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if len(req.ImageFrontSide) == 0 {
		return nil, errors.New("front side image is required")
	}

	// Convert images to base64
	base64FrontSide, err := ConvertToBase64(req.ImageFrontSide)
	if err != nil {
		return nil, err
	}

	payload := map[string]any{
		"input_type":     "base64",
		"imagefrontside": base64FrontSide,
	}

	if len(req.ImageBackSide) > 0 {
		base64BackSide, err := ConvertToBase64(req.ImageBackSide)
		if err != nil {
			return nil, err
		}
		payload["imagebackside"] = base64BackSide
	}

	if len(req.Images) > 0 {
		base64Images, err := ConvertToBase64(req.Images)
		if err != nil {
			return nil, err
		}
		payload["images"] = base64Images
	}

	resp, err := dc.sendRequest("POST", "/api/v1/document/analysis", payload)
	if err != nil {
		return nil, err
	}

	var result DocumentAnalysisResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("document analysis failed with status code: %d", resp.StatusCode)
	}

	return &result, nil
}

// ==================== GET VERIFICATION DETAILS ====================

// GetVerificationDetails retrieves verification details by reference ID
func (dc *DojahClient) GetVerificationDetails(referenceID string) (*GetVerificationDetailsResponse, error) {
	if referenceID == "" {
		return nil, errors.New("reference_id is required")
	}

	params := map[string]string{
		"reference_id": referenceID,
	}

	resp, err := dc.sendQueryRequest("/api/v1/kyc/verification", params)
	if err != nil {
		return nil, err
	}

	var result GetVerificationDetailsResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get verification details failed with status code: %d", resp.StatusCode)
	}

	return &result, nil
}

// ==================== LIST ALL VERIFICATIONS ====================

// ListAllVerifications retrieves all verifications with optional filters
func (dc *DojahClient) ListAllVerifications(searchTerm, startDate, endDate, status string) (*ListAllVerificationsResponse, error) {
	params := make(map[string]string)

	if searchTerm != "" {
		params["term"] = searchTerm
	}
	if startDate != "" {
		params["start"] = startDate
	}
	if endDate != "" {
		params["end"] = endDate
	}
	if status != "" {
		params["status"] = status
	}

	resp, err := dc.sendQueryRequest("/api/v1/kyc/verifications", params)
	if err != nil {
		return nil, err
	}

	var result ListAllVerificationsResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list verifications failed with status code: %d", resp.StatusCode)
	}

	return &result, nil
}
