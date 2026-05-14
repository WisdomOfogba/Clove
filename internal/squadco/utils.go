package squadco

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/gofiber/fiber/v3"
)

const squadWebhookSignatureHeader = "x-squad-encrypted-body"

// WebhookNotification is the generic format for Squad webhook payloads.
// The Body field is kept generic because Squad webhook payloads vary by transaction type.
type WebhookNotification struct {
	Event          string          `json:"Event"`
	TransactionRef string          `json:"TransactionRef"`
	Body           json.RawMessage `json:"Body"`
}

// VerifyWebhookSignature checks the Squad webhook signature header against the request body.
// It computes HMAC-SHA512 of the payload using the provided secret key.
func VerifyWebhookSignature(payload []byte, signatureHeader, secretKey string) (bool, error) {
	if len(secretKey) == 0 {
		return false, errors.New("secret key is required")
	}

	sig := strings.TrimSpace(signatureHeader)
	if sig == "" {
		return false, errors.New("missing x-squad-encrypted-body header")
	}

	mac := hmac.New(sha512.New, []byte(secretKey))
	_, err := mac.Write(payload)
	if err != nil {
		return false, err
	}

	expected := strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
	received := strings.ToUpper(sig)
	return hmac.Equal([]byte(expected), []byte(received)), nil
}

// ReadWebhookRequest reads the request body and validates the Squad webhook signature.
// It returns the raw payload, the parsed signature header, and any validation error.
func ReadWebhookRequest(c fiber.Ctx) ([]byte, string, error) {
	signature := c.Get(squadWebhookSignatureHeader)
	if signature == "" {
		return nil, "", errors.New("missing webhook signature header")
	}

	payload, err := io.ReadAll(c.Request().BodyStream())
	if err != nil {
		return nil, "", err
	}

	return payload, signature, nil
}

// ParseWebhookNotification parses a Squad webhook JSON payload into WebhookNotification.
func ParseWebhookNotification(payload []byte) (*WebhookNotification, error) {
	var notification WebhookNotification
	if err := json.Unmarshal(payload, &notification); err != nil {
		return nil, err
	}
	return &notification, nil
}

// ValidateWebhookRequest reads the webhook request, verifies the signature, and returns the payload.
func ValidateWebhookRequest(r fiber.Ctx, secretKey string) ([]byte, error) {
	payload, signature, err := ReadWebhookRequest(r)
	if err != nil {
		return nil, err
	}

	valid, err := VerifyWebhookSignature(payload, signature, secretKey)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, errors.New("webhook signature validation failed")
	}

	return payload, nil
}
