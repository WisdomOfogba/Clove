package dojah

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/gofiber/fiber/v3"
)

const dojahWebhookSignatureHeader = "x-dojah-signature"
const dojahWebhookSignatureHeaderV2 = "x-dojah-signature-v2"

// ConvertToBase64 converts byte slice to base64 string
// It removes the data:image/jpeg;base64, prefix if present
func ConvertToBase64(data []byte) (string, error) {
	if len(data) == 0 {
		return "", errors.New("empty data provided")
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// VerifyWebhookSignature verifies the Dojah webhook signature using payload and secret key
// Supports x-dojah-signature (HMAC-SHA256 of payload)
func VerifyWebhookSignature(payload []byte, signatureHeader, secretKey string) (bool, error) {
	if len(secretKey) == 0 {
		return false, errors.New("secret key is required")
	}

	sig := strings.TrimSpace(signatureHeader)
	if sig == "" {
		return false, errors.New("missing x-dojah-signature header")
	}

	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write(payload)
	if err != nil {
		return false, err
	}

	expected := strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
	received := strings.ToUpper(sig)
	return hmac.Equal([]byte(expected), []byte(received)), nil
}

// VerifyWebhookSignatureV2 verifies the Dojah webhook signature using only secret key
// Supports x-dojah-signature-v2 (HMAC-SHA256 of secret key only)
func VerifyWebhookSignatureV2(signatureHeader, secretKey string) (bool, error) {
	if len(secretKey) == 0 {
		return false, errors.New("secret key is required")
	}

	sig := strings.TrimSpace(signatureHeader)
	if sig == "" {
		return false, errors.New("missing x-dojah-signature-v2 header")
	}

	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(secretKey))
	if err != nil {
		return false, err
	}

	expected := strings.ToLower(hex.EncodeToString(mac.Sum(nil)))
	received := strings.ToLower(sig)
	return hmac.Equal([]byte(expected), []byte(received)), nil
}

// ReadWebhookRequest reads the request body and returns the payload and signature header
func ReadWebhookRequest(c fiber.Ctx) ([]byte, string, error) {
	signature := c.Get(dojahWebhookSignatureHeader)
	if signature == "" {
		return nil, "", errors.New("missing webhook signature header")
	}

	payload, err := io.ReadAll(c.Request().BodyStream())
	if err != nil {
		return nil, "", err
	}

	return payload, signature, nil
}

// ReadWebhookRequestV2 reads the request body and returns the payload and signature header v2
func ReadWebhookRequestV2(c fiber.Ctx) ([]byte, string, error) {
	signature := c.Get(dojahWebhookSignatureHeaderV2)
	if signature == "" {
		return nil, "", errors.New("missing webhook signature header v2")
	}

	payload, err := io.ReadAll(c.Request().BodyStream())
	if err != nil {
		return nil, "", err
	}

	return payload, signature, nil
}

// ValidateWebhookRequest validates the webhook request using signature verification with payload
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

// ValidateWebhookRequestV2 validates the webhook request using signature verification with secret only
func ValidateWebhookRequestV2(r fiber.Ctx, secretKey string) ([]byte, error) {
	payload, signature, err := ReadWebhookRequestV2(r)
	if err != nil {
		return nil, err
	}

	valid, err := VerifyWebhookSignatureV2(signature, secretKey)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, errors.New("webhook signature validation failed (v2)")
	}

	return payload, nil
}

// ParseWebhookPayload parses webhook JSON payload into the provided interface
func ParseWebhookPayload(payload []byte, v any) error {
	if err := json.Unmarshal(payload, v); err != nil {
		return err
	}
	return nil
}
