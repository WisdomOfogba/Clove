package middlewares

import (
	"time"

	"github.com/chibx/vendor-pulse/internal/constants"
	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/chibx/vendor-pulse/internal/types/response"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

var logger = global.Logger

func EnsureDeviceIDIsSet(ctx fiber.Ctx) error {
	deviceId := ctx.Cookies(constants.DeviceIDKey)
	deviceUUID, err := uuid.Parse(deviceId)
	if err != nil {
		deviceUUID, err = uuid.NewRandom()
		if err != nil {
			global.Logger.Error().Err(err).Msg("Error generating deviceUUID")
			return response.FromFiberError(ctx, fiber.ErrInternalServerError)
		}

		deviceId = deviceUUID.String()

		// Set the deviceId anyways
		ctx.Cookie(&fiber.Cookie{
			Name:     constants.DeviceIDKey,
			Value:    deviceId,
			Expires:  time.Now().Add(constants.DeviceIDDur),
			SameSite: "Strict",
			HTTPOnly: true,
			Secure:   true,
		})
	}
	ctx.Locals(constants.DeviceIDKey, deviceUUID)
	return ctx.Next()
}
