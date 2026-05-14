package routes

import (
	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/chibx/vendor-pulse/internal/squadco"
	"github.com/gofiber/fiber/v3"
)

func HandleWebHook() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		_, err := squadco.ValidateWebhookRequest(ctx, global.SquadClient.Secret())
		if err != nil {
			return fiber.ErrBadRequest
		}

		return nil
	}
}
