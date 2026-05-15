package routes

import (
	"encoding/json"

	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/chibx/vendor-pulse/internal/squadco"
	"github.com/gofiber/fiber/v3"
)

func HandleWebHook() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		payload, err := squadco.ValidateWebhookRequest(ctx, global.SquadClient.Secret())
		if err == nil {
			payloadMap := make(map[string]any)
			err := json.Unmarshal(payload, &payloadMap)
			if err != nil {
				logger.Err(err).Msg("Failed to unmarshal squadco webhook body")
			}
			logger.Info().Str("payload", string(payload))
		}

		return nil
	}
}
