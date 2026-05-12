package squadco

// ==================== PAYMENT INITIATION ====================

// InitiatePaymentRequest represents the request payload for initiating a payment
type InitiatePaymentRequest struct {
	Amount          int64                  `json:"amount"`
	Email           string                 `json:"email"`
	Currency        string                 `json:"currency"`
	InitiateType    string                 `json:"initiate_type"`
	TransactionRef  string                 `json:"transaction_ref,omitempty"`
	CustomerName    string                 `json:"customer_name,omitempty"`
	CallbackURL     string                 `json:"callback_url,omitempty"`
	PaymentChannels []string               `json:"payment_channels,omitempty"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
	PassCharge      bool                   `json:"pass_charge,omitempty"`
	SubMerchantID   string                 `json:"sub_merchant_id,omitempty"`
	IsRecurring     bool                   `json:"is_recurring,omitempty"`
}

// MerchantInfo contains merchant details
type MerchantInfo struct {
	MerchantResponse string `json:"merchant_response"`
	MerchantName     string `json:"merchant_name"`
	MerchantLogo     string `json:"merchant_logo"`
	MerchantID       string `json:"merchant_id"`
}

// RecurringInfo contains recurring payment details
type RecurringInfo struct {
	Frequency    *string `json:"frequency"`
	Duration     *string `json:"duration"`
	Type         int     `json:"type"`
	PlanCode     *string `json:"plan_code"`
	CustomerName *string `json:"customer_name"`
}

// InitiatePaymentResponse represents the response from initiate payment endpoint
type InitiatePaymentResponse struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Success bool                 `json:"success"`
	Data    *InitiatePaymentData `json:"data"`
}

// InitiatePaymentData contains the payment data
type InitiatePaymentData struct {
	AuthURL            *string       `json:"auth_url"`
	AccessToken        *string       `json:"access_token"`
	MerchantInfo       MerchantInfo  `json:"merchant_info"`
	Currency           string        `json:"currency"`
	Recurring          RecurringInfo `json:"recurring"`
	IsRecurring        bool          `json:"is_recurring"`
	PlanCode           *string       `json:"plan_code"`
	CallbackURL        string        `json:"callback_url"`
	TransactionRef     string        `json:"transaction_ref"`
	TransactionMemo    *string       `json:"transaction_memo"`
	TransactionAmount  int64         `json:"transaction_amount"`
	AuthorizedChannels []string      `json:"authorized_channels"`
	CheckoutURL        string        `json:"checkout_url"`
}

// ==================== CHARGE CARD ====================

// ChargeCardRequest represents the request to charge a tokenized card
type ChargeCardRequest struct {
	Amount         int64  `json:"amount"`
	TokenID        string `json:"token_id"`
	TransactionRef string `json:"transaction_ref,omitempty"`
}

// ChargeCardResponse represents the response from charge card endpoint
type ChargeCardResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Success bool                `json:"success"`
	Data    *TransactionDetails `json:"data"`
}

// ==================== CANCEL RECURRING CHARGE ====================

// CancelRecurringRequest represents the request to cancel recurring payment
type CancelRecurringRequest struct {
	AuthCodes []string `json:"auth_code"`
}

// CancelRecurringResponse represents the response from cancel recurring endpoint
type CancelRecurringResponse struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Success bool                 `json:"success"`
	Data    *CancelRecurringData `json:"data"`
}

// CancelRecurringData contains cancelled auth codes
type CancelRecurringData struct {
	AuthCodes []string `json:"auth_code"`
}

// ==================== SIMULATE TRANSFER PAYMENT ====================

// SimulateTransferPaymentRequest represents the request to simulate transfer payment
type SimulateTransferPaymentRequest struct {
	VirtualAccountNumber string `json:"virtual_account_number"`
	Amount               string `json:"amount"`
}

// SimulateTransferPaymentResponse represents the response from simulate transfer payment
type SimulateTransferPaymentResponse struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// ==================== VERIFY TRANSACTION ====================

// TransactionDetails represents detailed transaction information
type TransactionDetails struct {
	ID                    int64                  `json:"id"`
	TransactionAmount     int64                  `json:"transaction_amount"`
	TransactionRef        string                 `json:"transaction_ref"`
	Email                 string                 `json:"email"`
	MerchantID            string                 `json:"merchant_id"`
	MerchantAmount        int64                  `json:"merchant_amount"`
	MerchantName          string                 `json:"merchant_name"`
	MerchantBusinessName  *string                `json:"merchant_business_name"`
	MerchantEmail         string                 `json:"merchant_email"`
	CustomerEmail         string                 `json:"customer_email"`
	CustomerName          *string                `json:"customer_name"`
	Metadata              string                 `json:"meta_data"`
	Meta                  map[string]interface{} `json:"meta"`
	TransactionStatus     string                 `json:"transaction_status"`
	TransactionCharges    int64                  `json:"transaction_charges"`
	TransactionCurrencyID string                 `json:"transaction_currency_id"`
	TransactionGatewayID  string                 `json:"transaction_gateway_id"`
	TransactionType       string                 `json:"transaction_type"`
	FlatCharge            int64                  `json:"flat_charge"`
	IsSuspicious          bool                   `json:"is_suspicious"`
	IsRefund              bool                   `json:"is_refund"`
	CreatedAt             string                 `json:"created_at"`
	GatewayTransactionRef string                 `json:"gateway_transaction_ref"`
	Recurring             *RecurringInfo         `json:"recurring"`
	PlanCode              *string                `json:"plan_code"`
	PaymentInformation    *PaymentInformation    `json:"payment_information"`
	IsRecurring           bool                   `json:"is_recurring"`
}

// PaymentInformation contains payment method details
type PaymentInformation struct {
	PaymentType string `json:"payment_type"`
	PAN         string `json:"pan"`
	CardType    string `json:"card_type"`
	TokenID     string `json:"token_id"`
}

// VerifyTransactionResponse represents the response from verify transaction endpoint
type VerifyTransactionResponse struct {
	Status  int                 `json:"status"`
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    *TransactionDetails `json:"data"`
}

// ==================== QUERY ALL TRANSACTIONS ====================

// QueryTransactionsRequest represents the request to query transactions
type QueryTransactionsRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Currency  string `json:"currency,omitempty"`
	Page      int    `json:"page,omitempty"`
	PerPage   int    `json:"perpage,omitempty"`
	Reference string `json:"reference,omitempty"`
}

// QueryTransactionsResponse represents the response from query transactions endpoint
type QueryTransactionsResponse struct {
	Status  int                  `json:"status"`
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Data    []TransactionDetails `json:"data"`
}

// ==================== TRANSFER API ====================

// LookupBankAccountRequest represents the request to lookup a bank account
type LookupBankAccountRequest struct {
	BankCode      string `json:"bank_code"`
	AccountNumber string `json:"account_number"`
}

// LookupBankAccountResponse represents the response from lookup bank account endpoint
type LookupBankAccountResponse struct {
	Status  int              `json:"status"`
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    *BankAccountInfo `json:"data"`
}

// BankAccountInfo contains bank account information
type BankAccountInfo struct {
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
}

// InitiateTransferRequest represents the request to initiate a transfer
type InitiateTransferRequest struct {
	TransactionReference string `json:"transaction_reference"`
	Amount               string `json:"amount"`
	BankCode             string `json:"bank_code"`
	AccountNumber        string `json:"account_number"`
	AccountName          string `json:"account_name"`
	CurrencyID           string `json:"currency_id"`
	Remark               string `json:"remark"`
}

// TransferData contains transfer response details
type TransferData struct {
	TransactionReference       string `json:"transaction_reference"`
	ResponseDescription        string `json:"response_description"`
	CurrencyID                 string `json:"currency_id"`
	Amount                     string `json:"amount"`
	NipTransactionReference    string `json:"nip_transaction_reference"`
	AccountNumber              string `json:"account_number"`
	AccountName                string `json:"account_name"`
	DestinationInstitutionName string `json:"destination_institution_name"`
}

// InitiateTransferResponse represents the response from initiate transfer endpoint
type InitiateTransferResponse struct {
	Status  int           `json:"status"`
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    *TransferData `json:"data"`
}

// RequeryTransferRequest represents the request to requery transfer status
type RequeryTransferRequest struct {
	TransactionReference string `json:"transaction_reference"`
}

// RequeryTransferResponse represents the response from requery transfer endpoint
type RequeryTransferResponse struct {
	Remark               string `json:"remark"`
	BankCode             string `json:"bank_code"`
	CurrencyID           string `json:"currency_id"`
	Amount               string `json:"amount"`
	AccountNumber        string `json:"account_number"`
	TransactionReference string `json:"transaction_reference"`
	AccountName          string `json:"account_name"`
	Status               int    `json:"status"`
	Success              bool   `json:"success"`
	Message              string `json:"message"`
}

// TransferRecord represents a single transfer record
type TransferRecord struct {
	AccountNumberCredited string  `json:"account_number_credited"`
	AmountDebited         string  `json:"amount_debited"`
	TotalAmountDebited    string  `json:"total_amount_debited"`
	Success               bool    `json:"success"`
	Recipient             string  `json:"recipient"`
	BankCode              string  `json:"bank_code"`
	TransactionReference  string  `json:"transaction_reference"`
	TransactionStatus     string  `json:"transaction_status"`
	SwitchTransaction     *string `json:"switch_transaction"`
}

// GetAllTransfersResponse represents the response from get all transfers endpoint
type GetAllTransfersResponse struct {
	Status  int              `json:"status"`
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    []TransferRecord `json:"data"`
}

// ==================== DYNAMIC VIRTUAL ACCOUNT API ====================

// CreateDynamicVirtualAccountRequest represents the request to create a dynamic virtual account
type CreateDynamicVirtualAccountRequest struct {
	FirstName          string `json:"first_name,omitempty"`
	LastName           string `json:"last_name,omitempty"`
	BeneficiaryAccount string `json:"beneficiary_account,omitempty"`
}

// CreateDynamicVirtualAccountResponse represents the response from create dynamic virtual account endpoint
type CreateDynamicVirtualAccountResponse struct {
	Status  int            `json:"status"`
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    map[string]any `json:"data"`
}

// InitiateDynamicVirtualAccountRequest represents the request to initiate a dynamic virtual account transaction
type InitiateDynamicVirtualAccountRequest struct {
	Amount         int64  `json:"amount"`
	Duration       int64  `json:"duration"`
	Email          string `json:"email"`
	TransactionRef string `json:"transaction_ref"`
	Currency       string `json:"currency,omitempty"`
}

// InitiateDynamicVirtualAccountResponse represents the response from initiate dynamic virtual account transaction
type InitiateDynamicVirtualAccountResponse struct {
	Status  int                                `json:"status"`
	Success bool                               `json:"success"`
	Message string                             `json:"message"`
	Data    *InitiateDynamicVirtualAccountData `json:"data"`
}

// InitiateDynamicVirtualAccountData contains the dynamic virtual account details
type InitiateDynamicVirtualAccountData struct {
	IsBlocked            bool   `json:"is_blocked"`
	AccountName          string `json:"account_name"`
	AccountNumber        string `json:"account_number"`
	ExpectedAmount       string `json:"expected_amount"`
	ExpiresAt            string `json:"expires_at"`
	TransactionReference string `json:"transaction_reference"`
	Bank                 string `json:"bank"`
	Currency             string `json:"currency"`
}

// DynamicVirtualAccountTransactionRow represents a single dynamic virtual account transaction attempt
type DynamicVirtualAccountTransactionRow struct {
	TransactionStatus    string `json:"transaction_status"`
	TransactionReference string `json:"transaction_reference"`
	CreatedAt            string `json:"created_at"`
	Refund               *bool  `json:"refund"`
}

// DynamicVirtualAccountTransactionsData contains transaction attempt rows
type DynamicVirtualAccountTransactionsData struct {
	Count int                                   `json:"count"`
	Rows  []DynamicVirtualAccountTransactionRow `json:"rows"`
}

// DynamicVirtualAccountTransactionsResponse represents the response for DVA transaction history
type DynamicVirtualAccountTransactionsResponse struct {
	Status  int                                    `json:"status"`
	Success bool                                   `json:"success"`
	Message string                                 `json:"message"`
	Data    *DynamicVirtualAccountTransactionsData `json:"data"`
}

// UpdateDynamicVirtualAccountRequest represents the request to update amount or duration for a DVA transaction
type UpdateDynamicVirtualAccountRequest struct {
	TransactionReference string `json:"transaction_reference"`
	Amount               int64  `json:"amount,omitempty"`
	Duration             int64  `json:"duration,omitempty"`
}

// UpdateDynamicVirtualAccountResponse represents the response from updating a DVA transaction
type UpdateDynamicVirtualAccountResponse struct {
	Status  int                                `json:"status"`
	Success bool                               `json:"success"`
	Message string                             `json:"message"`
	Data    *InitiateDynamicVirtualAccountData `json:"data"`
}

// SimulateDynamicVirtualAccountPaymentRequest represents the request to simulate a DVA payment in sandbox
type SimulateDynamicVirtualAccountPaymentRequest struct {
	VirtualAccountNumber string `json:"virtual_account_number"`
	Amount               string `json:"amount"`
	DVA                  bool   `json:"dva"`
}

// ==================== REFUND API ====================

// RefundRequest represents the request to initiate a refund
type RefundRequest struct {
	GatewayTransactionRef string `json:"gateway_transaction_ref"`
	TransactionRef        string `json:"transaction_ref"`
	RefundType            string `json:"refund_type"`
	ReasonForRefund       string `json:"reason_for_refund"`
	RefundAmount          string `json:"refund_amount,omitempty"`
}

// RefundData contains refund response details
type RefundData struct {
	GatewayRefundStatus string `json:"gateway_refund_status"`
	RefundStatus        int    `json:"refund_status"`
	RefundReference     string `json:"refund_reference"`
}

// RefundResponse represents the response from refund endpoint
type RefundResponse struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    *RefundData `json:"data"`
}

// ==================== LEDGER BALANCE ====================

// LedgerBalanceResponse represents the response from ledger balance endpoint
type LedgerBalanceResponse struct {
	Status  int                `json:"status"`
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    *LedgerBalanceData `json:"data"`
}

// LedgerBalanceData contains ledger balance information
type LedgerBalanceData struct {
	Balance    string `json:"balance"`
	CurrencyID string `json:"currency_id"`
	MerchantID string `json:"merchant_id"`
}
