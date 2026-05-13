package middlewares

import (
	"strings"

	"github.com/chibx/vendor-pulse/internal/auth"
	"github.com/chibx/vendor-pulse/internal/constants"
	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/chibx/vendor-pulse/internal/types/request"
	"github.com/chibx/vendor-pulse/internal/types/response"

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware(c fiber.Ctx) error {
	// var tokenErr error
	// authHeader := strings.TrimSpace(c.Get("Authorization"))
	// var tokenStr string
	// if authHeader != "" {
	// 	tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
	// 	if tokenStr == authHeader {
	// 		return response.FromFiberError(c, fiber.ErrBadRequest)
	// 	}
	// }
	var authUser request.UserCtx

	userAccessToken := strings.TrimSpace(c.Cookies(constants.UserAccessTkKey))

	if userAccessToken != "" {
		validJWT, err := auth.ValidateBackendAccessToken(userAccessToken, global.SecretKey)
		if err == nil {
			authUser = request.UserCtx{ID: validJWT.UserID}
			c.Locals(constants.UserCtxKey, authUser)
		} else {
			logger.Error().Err(err).Msg("Error during authentication")
		}
	}

	return c.Next()
}

func HardenBackendEndpoint(c fiber.Ctx) error {
	backendUser, ok := c.Locals(constants.UserCtxKey).(*request.UserCtx)

	if !ok || backendUser == nil {
		return response.FromFiberError(c, fiber.ErrUnauthorized, "You need to login.")
	}

	return c.Next()
}
