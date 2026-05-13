package routes

import (
	"github.com/chibx/vendor-pulse/api/middlewares"
	"github.com/gofiber/fiber/v3"
)

func AddRoutes(app *fiber.App) {
	api := app.Group("/api", middlewares.EnsureDeviceIDIsSet)
	customerAuth := api.Group("/auth/customer")
	customerAuth.Post("/signup", UserRegister())
	customerAuth.Post("/login", UserLogin())
	customerAuth.Post("/logout", UserLogout())
	customerAuth.Post("/promote-account", UpgradeToBusinessAccount())
	

	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"uri":  ctx.Request().URI().String(),
			"path": ctx.Path(),
		})
	})
}
