package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net"
	"time"

	"github.com/chibx/vendor-pulse/internal/constants"
	"github.com/chibx/vendor-pulse/internal/db"
	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/chibx/vendor-pulse/internal/model"

	serverErrors "github.com/chibx/vendor-pulse/internal/server_errors"
	"github.com/gofiber/fiber/v3"
)

// GenerateRefreshToken (as before).
func GenerateRefreshToken() (string, error) {
	// ... (crypto/rand + hex)
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func CompositeRefreshToken() (token string, refreshHash string, err error) {
	refreshToken, err := GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}
	refreshTokenHash, err := GenerateHashFromString(refreshToken, DefaultHashParams)
	if err != nil {
		return "", "", err
	}

	return refreshToken, refreshTokenHash, nil
}

func ValidateBackendUserSess(ctx fiber.Ctx, session *model.UserSession) error {
	created_at := session.CreatedAt
	lastIp := session.LastIP
	// Handle user_agent validation later
	_ = session.UserAgent

	// TODO: Add validation logic for expiry and IP address
	_ = created_at
	_ = lastIp

	current_time := time.Now()
	if current_time.Sub(created_at) > constants.UserRefreshTkDur {
		return serverErrors.NewSessionErr(serverErrors.SessionExpired, "Session has expired")
	}
	// CAUTION: This is a basic IP check.
	ip := net.ParseIP(ctx.IP())
	if ip == nil {
		return serverErrors.NewSessionErr(serverErrors.SessionInvalidIpAddr, "IP address is either missing or invalid")
	}

	if ip.String() != lastIp {
		// TODO: Maybe check if the IP range is valid and secure instead of rejecting it outright
		return serverErrors.NewSessionErr(serverErrors.SessionDiffIpAddr, "IP address does not match")
	}

	return nil
}

func CreateBackendSession(ctx context.Context, session *model.UserSession) error {
	var err error
	logger := global.Logger
	err = db.Users().CreateSession(ctx, session)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create backend session")
	}

	return err
}

func GetUserSession(ctx context.Context, tokenHash string) (*model.UserSession, error) {
	logger := global.Logger

	var user_session *model.UserSession

	user_session, err := db.Users().GetSessionByTokenHash(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, serverErrors.ErrDBRecordNotFound) {
			logger.Error().Err(err).Msg("backend user session not found in db")
			return nil, serverErrors.NewServerErr(fiber.StatusUnauthorized, "User Session not found. Consider logging in again")
		}
		logger.Error().Err(err).Msg("failed to get backend user session from db")
		return nil, serverErrors.NewServerErr(fiber.StatusInternalServerError, "Something went wrong while fetching your session data. Please try again later.")
	}

	return user_session, nil
}
