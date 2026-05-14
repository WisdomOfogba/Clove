package squadco

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	baseURLSandbox = "https://sandbox-api-d.squadco.com"
	baseURLLive    = "https://api-d.squadco.com"
)

// SetEnvironment sets the API environment (sandbox or live)
// If useProduction is true, live API is used; otherwise sandbox is used
func (sc *Client) SetEnvironment(useProduction bool) {
	if useProduction {
		sc.baseURL = baseURLLive
		return
	}
	sc.baseURL = baseURLSandbox
}

// ==================== INTERNAL HELPER METHODS ====================

// buildURL constructs the full URL for an endpoint
func (sc *Client) buildURL(baseURL, endpoint string) string {
	return fmt.Sprintf("%s%s", baseURL, endpoint)
}

// sendRequest sends an HTTP request to the Squad API
func (sc *Client) sendRequest(method, endpoint string, payload any) (*http.Response, error) {
	url := sc.buildURL(sc.baseURL, endpoint)

	var body io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request: %w", err)
		}
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	// They didn't use "Bearer " prefix in their docs :)
	req.Header.Set("Authorization", sc.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := sc.httpclient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}

// sendQueryRequest sends an HTTP GET request with query parameters
func (sc *Client) sendQueryRequest(endpoint string, params map[string]string) (*http.Response, error) {
	requestUrl := sc.buildURL(sc.baseURL, endpoint)

	// Add query parameters
	if len(params) > 0 {
		q, err := url.Parse(requestUrl)
		if err != nil {
			log.Println("[SquadCo]: Invalid URL scheme")
			return nil, err
		}
		if q.Query() == nil {
			q.RawQuery = ""
		}
		query := q.Query()
		for k, v := range params {
			query.Add(k, v)
		}
		q.RawQuery = query.Encode()
		requestUrl = q.String()
	}

	// Parse URL for the query parameters properly
	parsedURL, err := url.Parse(requestUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", sc.apiKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := sc.httpclient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
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
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, v); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

// ==================== PAYMENT INITIATION METHODS ====================

// InitiatePayment initiates a payment transaction and returns a checkout URL
func (sc *Client) InitiatePayment(req *InitiatePaymentRequest) (*InitiatePaymentResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.Email == "" {
		return nil, fmt.Errorf("email is required")
	}

	if req.Amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than 0")
	}

	if req.Currency == "" {
		req.Currency = "NGN"
	}

	if req.InitiateType == "" {
		req.InitiateType = "inline"
	}

	resp, err := sc.sendRequest("POST", "/transaction/initiate", req)
	if err != nil {
		return nil, err
	}

	var result InitiatePaymentResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to initiate payment: %s", result.Message)
	}

	return &result, nil
}

// ChargeCard charges a previously tokenized card
func (sc *Client) ChargeCard(req *ChargeCardRequest) (*ChargeCardResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.TokenID == "" {
		return nil, fmt.Errorf("token_id is required")
	}

	if req.Amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than 0")
	}

	resp, err := sc.sendRequest("POST", "/transaction/charge_card", req)
	if err != nil {
		return nil, err
	}

	var result ChargeCardResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to charge card: %s", result.Message)
	}

	return &result, nil
}

// SimulateTransferPayment simulates a payment into a virtual account (for testing)
func (sc *Client) SimulateTransferPayment(req *SimulateTransferPaymentRequest) (*SimulateTransferPaymentResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.VirtualAccountNumber == "" {
		return nil, fmt.Errorf("virtual_account_number is required")
	}

	if req.Amount == "" {
		return nil, fmt.Errorf("amount is required")
	}

	resp, err := sc.sendRequest("POST", "/virtual-account/simulate/payment", req)
	if err != nil {
		return nil, err
	}

	var result SimulateTransferPaymentResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to simulate transfer payment: %s", result.Message)
	}

	return &result, nil
}

// CancelRecurringCharge cancels previously tokenized card(s)
func (sc *Client) CancelRecurringCharge(req *CancelRecurringRequest) (*CancelRecurringResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if len(req.AuthCodes) == 0 {
		return nil, fmt.Errorf("at least one auth_code is required")
	}

	resp, err := sc.sendRequest("PATCH", "/transaction/cancel/recurring", req)
	if err != nil {
		return nil, err
	}

	var result CancelRecurringResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to cancel recurring charge: %s", result.Message)
	}

	return &result, nil
}

// ==================== TRANSACTION VERIFICATION METHODS ====================

// VerifyTransaction verifies the status of a transaction using its reference
func (sc *Client) VerifyTransaction(transactionRef string) (*VerifyTransactionResponse, error) {
	if transactionRef == "" {
		return nil, fmt.Errorf("transaction_ref is required")
	}

	endpoint := fmt.Sprintf("/transaction/verify/%s", transactionRef)

	resp, err := sc.sendRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result VerifyTransactionResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to verify transaction: %s", result.Message)
	}

	return &result, nil
}

// QueryAllTransactions retrieves all transactions within a date range
func (sc *Client) QueryAllTransactions(startDate, endDate string, options map[string]any) (*QueryTransactionsResponse, error) {
	if startDate == "" || endDate == "" {
		return nil, fmt.Errorf("start_date and end_date are required")
	}

	params := make(map[string]string)
	params["start_date"] = startDate
	params["end_date"] = endDate

	if currency, ok := options["currency"].(string); ok && currency != "" {
		params["currency"] = currency
	}

	if page, ok := options["page"].(int); ok && page > 0 {
		params["page"] = fmt.Sprintf("%d", page)
	}

	if perpage, ok := options["perpage"].(int); ok && perpage > 0 {
		params["perpage"] = fmt.Sprintf("%d", perpage)
	}

	if reference, ok := options["reference"].(string); ok && reference != "" {
		params["reference"] = reference
	}

	resp, err := sc.sendQueryRequest("/transaction", params)
	if err != nil {
		return nil, err
	}

	var result QueryTransactionsResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to query transactions: %s", result.Message)
	}

	return &result, nil
}

// ==================== TRANSFER API METHODS ====================

// LookupBankAccount verifies that a bank account belongs to the recipient
func (sc *Client) LookupBankAccount(req *LookupBankAccountRequest) (*LookupBankAccountResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.BankCode == "" {
		return nil, fmt.Errorf("bank_code is required")
	}

	if req.AccountNumber == "" {
		return nil, fmt.Errorf("account_number is required")
	}

	resp, err := sc.sendRequest("POST", "/payout/account/lookup", req)
	if err != nil {
		return nil, err
	}

	var result LookupBankAccountResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to lookup bank account: %s", result.Message)
	}

	return &result, nil
}

// InitiateTransfer initiates a transfer from Squad wallet to a bank account
func (sc *Client) InitiateTransfer(req *InitiateTransferRequest) (*InitiateTransferResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.TransactionReference == "" {
		return nil, fmt.Errorf("transaction_reference is required")
	}

	if req.Amount == "" {
		return nil, fmt.Errorf("amount is required")
	}

	if req.BankCode == "" {
		return nil, fmt.Errorf("bank_code is required")
	}

	if req.AccountNumber == "" {
		return nil, fmt.Errorf("account_number is required")
	}

	if req.AccountName == "" {
		return nil, fmt.Errorf("account_name is required")
	}

	if req.CurrencyID == "" {
		req.CurrencyID = "NGN"
	}

	resp, err := sc.sendRequest("POST", "/payout/transfer", req)
	if err != nil {
		return nil, err
	}

	var result InitiateTransferResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to initiate transfer: %s", result.Message)
	}

	return &result, nil
}

// RequeryTransfer checks the status of a transfer
func (sc *Client) RequeryTransfer(transactionRef string) (*RequeryTransferResponse, error) {
	if transactionRef == "" {
		return nil, fmt.Errorf("transaction_reference is required")
	}

	req := &RequeryTransferRequest{
		TransactionReference: transactionRef,
	}

	resp, err := sc.sendRequest("POST", "/payout/requery", req)
	if err != nil {
		return nil, err
	}

	var result RequeryTransferResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to requery transfer: %s", result.Message)
	}

	return &result, nil
}

// GetAllTransfers retrieves all transfers made from the Squad wallet
func (sc *Client) GetAllTransfers(page, perPage int, sortOrder string) (*GetAllTransfersResponse, error) {
	params := make(map[string]string)

	if page > 0 {
		params["page"] = fmt.Sprintf("%d", page)
	}

	if perPage > 0 {
		params["perPage"] = fmt.Sprintf("%d", perPage)
	}

	if sortOrder != "" && (sortOrder == "ASC" || sortOrder == "DESC") {
		params["dir"] = sortOrder
	}

	resp, err := sc.sendQueryRequest("/payout/list", params)
	if err != nil {
		return nil, err
	}

	var result GetAllTransfersResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get all transfers: %s", result.Message)
	}

	return &result, nil
}

// CreateDynamicVirtualAccount creates a dynamic virtual account within the merchant pool
func (sc *Client) CreateDynamicVirtualAccount(req *CreateDynamicVirtualAccountRequest) (*CreateDynamicVirtualAccountResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	resp, err := sc.sendRequest("POST", "/virtual-account/create-dynamic-virtual-account", req)
	if err != nil {
		return nil, err
	}

	var result CreateDynamicVirtualAccountResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create dynamic virtual account: %s", result.Message)
	}

	return &result, nil
}

// InitiateDynamicVirtualAccount initiates a dynamic virtual account transaction
func (sc *Client) InitiateDynamicVirtualAccount(req *InitiateDynamicVirtualAccountRequest) (*InitiateDynamicVirtualAccountResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.Amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than 0")
	}

	if req.Duration <= 0 {
		return nil, fmt.Errorf("duration must be greater than 0")
	}

	if req.Email == "" {
		return nil, fmt.Errorf("email is required")
	}

	if req.TransactionRef == "" {
		return nil, fmt.Errorf("transaction_ref is required")
	}

	resp, err := sc.sendRequest("POST", "/virtual-account/initiate-dynamic-virtual-account", req)
	if err != nil {
		return nil, err
	}

	var result InitiateDynamicVirtualAccountResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to initiate dynamic virtual account: %s", result.Message)
	}

	return &result, nil
}

// RequeryDynamicVirtualAccount retrieves transaction attempts for a dynamic virtual account reference
func (sc *Client) RequeryDynamicVirtualAccount(transactionRef string) (*DynamicVirtualAccountTransactionsResponse, error) {
	if transactionRef == "" {
		return nil, fmt.Errorf("transaction_reference is required")
	}

	endpoint := fmt.Sprintf("/virtual-account/get-dynamic-virtual-account-transactions/%s", transactionRef)
	resp, err := sc.sendRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result DynamicVirtualAccountTransactionsResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to requery dynamic virtual account: %s", result.Message)
	}

	return &result, nil
}

// UpdateDynamicVirtualAccount updates amount or duration for an existing dynamic virtual account transaction
func (sc *Client) UpdateDynamicVirtualAccount(req *UpdateDynamicVirtualAccountRequest) (*UpdateDynamicVirtualAccountResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.TransactionReference == "" {
		return nil, fmt.Errorf("transaction_reference is required")
	}

	if req.Amount <= 0 && req.Duration <= 0 {
		return nil, fmt.Errorf("either amount or duration must be provided")
	}

	resp, err := sc.sendRequest("PATCH", "/virtual-account/update-dynamic-virtual-account-time-and-amount", req)
	if err != nil {
		return nil, err
	}

	var result UpdateDynamicVirtualAccountResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to update dynamic virtual account: %s", result.Message)
	}

	return &result, nil
}

// SimulateDynamicVirtualAccountPayment simulates a dynamic virtual account payment in sandbox
func (sc *Client) SimulateDynamicVirtualAccountPayment(req *SimulateDynamicVirtualAccountPaymentRequest) (*SimulateTransferPaymentResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.VirtualAccountNumber == "" {
		return nil, fmt.Errorf("virtual_account_number is required")
	}

	if req.Amount == "" {
		return nil, fmt.Errorf("amount is required")
	}

	resp, err := sc.sendRequest("POST", "/virtual-account/simulate/payment", req)
	if err != nil {
		return nil, err
	}

	var result SimulateTransferPaymentResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to simulate dynamic virtual account payment: %s", result.Message)
	}

	return &result, nil
}

// VerifyPOSTransaction verifies a POS transaction by transaction reference
func (sc *Client) VerifyPOSTransaction(transactionRef string) (*VerifyTransactionResponse, error) {
	if transactionRef == "" {
		return nil, fmt.Errorf("transaction_ref is required")
	}

	endpoint := fmt.Sprintf("/softpos/transaction/verify/%s", transactionRef)
	resp, err := sc.sendRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result VerifyTransactionResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to verify POS transaction: %s", result.Message)
	}

	return &result, nil
}

// ==================== REFUND API METHODS ====================

// RefundTransaction initiates a refund for a successful transaction
func (sc *Client) RefundTransaction(req *RefundRequest) (*RefundResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.GatewayTransactionRef == "" {
		return nil, fmt.Errorf("gateway_transaction_ref is required")
	}

	if req.TransactionRef == "" {
		return nil, fmt.Errorf("transaction_ref is required")
	}

	if req.RefundType == "" {
		return nil, fmt.Errorf("refund_type is required (Full or Partial)")
	}

	if req.RefundType != "Full" && req.RefundType != "Partial" {
		return nil, fmt.Errorf("refund_type must be either 'Full' or 'Partial'")
	}

	if req.RefundType == "Partial" && req.RefundAmount == "" {
		return nil, fmt.Errorf("refund_amount is required for partial refunds")
	}

	if req.ReasonForRefund == "" {
		return nil, fmt.Errorf("reason_for_refund is required")
	}

	resp, err := sc.sendRequest("POST", "/transaction/refund", req)
	if err != nil {
		return nil, err
	}

	var result RefundResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to refund transaction: %s", result.Message)
	}

	return &result, nil
}

// ==================== BALANCE METHODS ====================

// GetLedgerBalance retrieves the current balance in the Squad wallet
func (sc *Client) GetLedgerBalance(currencyID string) (*LedgerBalanceResponse, error) {
	if currencyID == "" {
		currencyID = "NGN"
	}

	if currencyID != "NGN" {
		return nil, fmt.Errorf("only NGN currency is supported for ledger balance")
	}

	params := make(map[string]string)
	params["currency_id"] = currencyID

	resp, err := sc.sendQueryRequest("/merchant/balance", params)
	if err != nil {
		return nil, err
	}

	var result LedgerBalanceResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	if !result.Success && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get ledger balance: %s", result.Message)
	}

	return &result, nil
}
