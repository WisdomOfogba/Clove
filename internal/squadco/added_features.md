# Squadco API Implementation Features

## Implementation Date
May 11, 2026

## Overview
This document tracks all implemented Squadco API methods and features. The implementation provides comprehensive integration with Squadco's payment, transfer, refund, and balance management APIs.

## Files Structure
- `squadco.go` - Core client initialization
- `types.go` - All request/response structs with JSON tags
- `methods.go` - All API method implementations

---

## Implemented Features

### 1. Payment Initiation Methods
#### Status: ✅ Complete

- **InitiatePayment()**
  - Initiates a payment transaction and returns a checkout URL
  - Parameters: `InitiatePaymentRequest` with amount, email, currency, etc.
  - Returns: `InitiatePaymentResponse` with checkout URL and merchant info
  - Supports: Multiple payment channels (card, transfer, USSD, bank)
  - Features: Card tokenization, recurring payments, metadata, callback URLs

- **ChargeCard()**
  - Charges a previously tokenized card
  - Parameters: `ChargeCardRequest` with amount, token_id, transaction_ref
  - Returns: `ChargeCardResponse` with transaction details
  - Use case: For recurring/subscriptions using saved card tokens

- **SimulateTransferPayment()**
  - Simulates a payment into a virtual account (test environment only)
  - Parameters: `SimulateTransferPaymentRequest` with virtual account number and amount
  - Returns: `SimulateTransferPaymentResponse`
  - Note: Used only in sandbox/test environment

- **CancelRecurringCharge()**
  - Cancels previously tokenized card(s)
  - Parameters: `CancelRecurringRequest` with array of auth_codes
  - Returns: `CancelRecurringResponse` with cancelled auth codes
  - Use case: For stopping recurring payments or deleting saved cards

---

### 2. Transaction Verification Methods
#### Status: ✅ Complete

- **VerifyTransaction()**
  - Verifies the status of a transaction using its reference
  - Parameters: `transactionRef` (string)
  - Returns: `VerifyTransactionResponse` with full transaction details
  - Status values: Success, Failed, Abandoned, Pending
  - Features: Returns payment method info, token IDs for recurring

- **QueryAllTransactions()**
  - Retrieves all transactions within a date range
  - Parameters: startDate, endDate (required), plus optional filters
  - Optional: currency, page, perpage, reference
  - Returns: `QueryTransactionsResponse` with array of transaction details
  - Note: Date range must not exceed one month

---

### 3. Transfer API Methods
#### Status: ✅ Complete

- **LookupBankAccount()**
  - Verifies that a bank account belongs to the recipient
  - Parameters: `LookupBankAccountRequest` with bank_code and account_number
  - Returns: `LookupBankAccountResponse` with account name and number
  - Required: Must be called before initiating a transfer
  - Supports: All Nigerian banks via NIP codes

- **InitiateTransfer()**
  - Initiates a transfer from Squad wallet to a bank account
  - Parameters: `InitiateTransferRequest` with account details, amount, reference
  - Returns: `InitiateTransferResponse` with transfer confirmation and NIP reference
  - Requirements: Account must be looked up first via LookupBankAccount()
  - Amount: In Kobo (1 NGN = 100 Kobo)
  - Features: Unique transaction references (must include merchant ID)

- **RequeryTransfer()**
  - Checks the status of a previously initiated transfer
  - Parameters: `transactionRef` (string)
  - Returns: `RequeryTransferResponse` with transfer status
  - Error codes: 200 (Success), 400 (Bad), 422 (Unprocessed), 424 (Timeout/Failed - requery), 404 (Not found), 412 (Reversed)
  - Use case: For determining if pending transfers succeeded or failed

- **GetAllTransfers()**
  - Retrieves all transfers made from the Squad wallet
  - Parameters: page (int), perPage (int), sortOrder (ASC/DESC)
  - Returns: `GetAllTransfersResponse` with array of transfer records
  - Features: Pagination, sorting, transaction status per transfer

---

### 4. Refund API Methods
#### Status: ✅ Complete

- **RefundTransaction()**
  - Initiates a refund for a successful transaction
  - Parameters: `RefundRequest` with gateway_transaction_ref, transaction_ref, refund type, reason
  - Returns: `RefundResponse` with refund reference and status
  - Refund types: 'Full' or 'Partial'
  - For Partial: refund_amount is required (in Kobo)
  - Returns: Refund reference for tracking

---

### 5. Balance & Account Methods
#### Status: ✅ Complete

- **GetLedgerBalance()**
  - Retrieves the current balance in the Squad wallet
  - Parameters: `currencyID` (optional, defaults to NGN)
  - Returns: `LedgerBalanceResponse` with balance, currency, merchant ID
  - Note: Balance is in Kobo (1 NGN = 100 Kobo)
  - Limitation: Only NGN currency is supported

---

### 6. Helper & Internal Methods
#### Status: ✅ Complete

- **buildURL()** - Constructs full URLs for API endpoints
- **sendRequest()** - Sends HTTP POST/PATCH requests with proper headers and authentication
- **sendQueryRequest()** - Sends HTTP GET requests with query parameters
- **decodeResponse()** - Decodes JSON responses into Go structs
- **Authorization** - All requests include Bearer token authentication

---

## Type Definitions (types.go)

### Request Types
1. `InitiatePaymentRequest` - Payment initiation payload
2. `ChargeCardRequest` - Card charging payload
3. `CancelRecurringRequest` - Cancel recurring charge payload
4. `SimulateTransferPaymentRequest` - Simulate transfer payment
5. `LookupBankAccountRequest` - Bank account lookup
6. `InitiateTransferRequest` - Fund transfer payload
7. `RequeryTransferRequest` - Transfer requery payload
8. `RefundRequest` - Refund initiation payload
9. `QueryTransactionsRequest` - Transaction query filters

### Response Types
1. `InitiatePaymentResponse` - Checkout URL and payment details
2. `ChargeCardResponse` - Card charge confirmation
3. `CancelRecurringResponse` - Cancelled auth codes
4. `SimulateTransferPaymentResponse` - Simulation confirmation
5. `VerifyTransactionResponse` - Transaction verification details
6. `QueryTransactionsResponse` - List of transactions
7. `LookupBankAccountResponse` - Bank account details
8. `InitiateTransferResponse` - Transfer confirmation
9. `RequeryTransferResponse` - Transfer status
10. `GetAllTransfersResponse` - List of transfers
11. `RefundResponse` - Refund confirmation
12. `LedgerBalanceResponse` - Account balance

### Supporting Structs
- `TransactionDetails` - Full transaction information
- `BankAccountInfo` - Bank account details
- `TransferData` - Transfer response data
- `TransferRecord` - Individual transfer in list
- `RefundData` - Refund response details
- `LedgerBalanceData` - Balance details
- `PaymentInformation` - Payment method details (card, etc.)
- `MerchantInfo` - Merchant information
- `RecurringInfo` - Recurring payment details

---

## API Endpoints Implemented

### Payment Endpoints
- `POST /transaction/initiate` - Initiate payment
- `POST /transaction/charge_card` - Charge tokenized card
- `POST /virtual-account/simulate/payment` - Simulate transfer payment
- `PATCH /transaction/cancel/recurring` - Cancel recurring charge

### Dynamic Virtual Account Endpoints
- `POST /virtual-account/create-dynamic-virtual-account` - Create dynamic virtual account
- `POST /virtual-account/initiate-dynamic-virtual-account` - Initiate dynamic virtual account transaction
- `GET /virtual-account/get-dynamic-virtual-account-transactions/{transaction_reference}` - Requery dynamic virtual account transaction status
- `PATCH /virtual-account/update-dynamic-virtual-account-time-and-amount` - Update DVA amount/duration
- `POST /virtual-account/simulate/payment` - Simulate dynamic virtual account payment (sandbox)

### Transaction Endpoints
- `GET /transaction/verify/{transaction_ref}` - Verify transaction
- `GET /transaction` - Query all transactions

### POS Endpoints
- `GET /softpos/transaction/verify/{transaction_ref}` - Verify POS transaction

### Transfer Endpoints
- `POST /payout/account/lookup` - Lookup bank account
- `POST /payout/transfer` - Initiate transfer
- `POST /payout/requery` - Requery transfer status
- `GET /payout/list` - Get all transfers

### Refund Endpoints
- `POST /transaction/refund` - Initiate refund

### Balance Endpoints
- `GET /merchant/balance` - Get ledger balance

---

## Validation & Error Handling

All methods include:
- ✅ Nil request validation
- ✅ Required parameter validation
- ✅ Type validation (e.g., refund_type must be Full or Partial)
- ✅ Amount validation (must be > 0)
- ✅ Email validation requirement
- ✅ HTTP status code checking
- ✅ JSON parsing error handling
- ✅ Descriptive error messages

---

## Usage Examples

### Initiate a Payment
```go
client, _ := squadco.NewSquadClient(squadco.SquadOption{
    ApiKey: "your_api_key",
})

paymentReq := &squadco.InitiatePaymentRequest{
    Amount:        100000, // Amount in kobo (1000 NGN)
    Email:         "customer@example.com",
    Currency:      "NGN",
    InitiateType:  "inline",
    TransactionRef: "unique_ref_12345",
    CallbackURL:   "https://yoursite.com/callback",
}

response, err := client.InitiatePayment(paymentReq)
if err != nil {
    log.Fatal(err)
}

fmt.Println(response.Data.CheckoutURL) // Redirect customer here
```

### Initiate a Transfer
```go
// First, lookup the account
lookupReq := &squadco.LookupBankAccountRequest{
    BankCode:      "000013", // GTBank
    AccountNumber: "0123456789",
}

lookup, err := client.LookupBankAccount(lookupReq)
if err != nil {
    log.Fatal(err)
}

// Then transfer
transferReq := &squadco.InitiateTransferRequest{
    TransactionReference: "MERCHANT_ID_unique_ref",
    Amount:              "50000", // 500 NGN in kobo
    BankCode:            "000013",
    AccountNumber:       "0123456789",
    AccountName:         lookup.Data.AccountName,
    CurrencyID:          "NGN",
    Remark:              "Payment for order #123",
}

transfer, err := client.InitiateTransfer(transferReq)
```

### Verify a Transaction
```go
verify, err := client.VerifyTransaction("transaction_ref_12345")
if err != nil {
    log.Fatal(err)
}

if verify.Data.TransactionStatus == "Success" {
    fmt.Println("Payment successful!")
}
```

---

## Next Steps / Future Enhancements


### Implemented Next-Step Features

- [x] Dynamic Virtual Account API support
  - [x] Create dynamic virtual account
  - [x] Initiate dynamic virtual account transaction
  - [x] Re-query dynamic virtual account transaction status
  - [x] Update dynamic virtual account amount/duration
  - [x] Simulate dynamic virtual account payment in sandbox
- [x] POS transaction verification via `VerifyPOSTransaction()`
- [x] Production/sandbox environment switching via `SetEnvironment()`
- [ ] Webhook signature validation helper

### Webhook Support
- [x] Added Squad webhook helpers in `utils.go`
  - [x] `VerifyWebhookSignature(payload, signature, secretKey)`
  - [x] `ReadWebhookRequest(r)`
  - [x] `ValidateWebhookRequest(r, secretKey)`
  - [x] `ParseWebhookNotification(payload)`
 
### Remaining Enhancements

1. **Subscriptions/Recurring** - Full subscription management
2. **Disputes & Chargebacks** - Handle payment disputes
3. **Webhooks** - Webhook verification and handling
4. **Rate Limiting** - Built-in rate limit handling
5. **Retry Logic** - Automatic retry for transient failures
6. **Logging** - Structured logging for all requests
7. **Monitoring / Telemetry** - Track request latency and failures

---

## Notes

- All amounts in Kobo (NGN) or Cent (USD): 1 NGN = 100 Kobo, 1 USD = 100 Cent
- Transaction references must be unique per transaction
- For transfers, transaction reference must include merchant ID (format: MERCHANT_ID_unique_ref)
- Date queries must not exceed one month gap
- Bearer token authentication required for all endpoints
- Test card for recurring payments: 5200000000000007
- Virtual accounts created in transfer flow are test-only in sandbox

---

## Status Codes Handled

- 200 - Success
- 400 - Bad Request
- 401 - Unauthorized (Invalid/Missing API Key)
- 403 - Forbidden (Authentication failed)
- 404 - Not Found (Transaction/Account not found)
- 422 - Unprocessed (Transfer pending)
- 424 - Timeout/Failed (Transfer - should requery)
- 412 - Reversed (Transfer reversed)
