package routes

import "github.com/gofiber/fiber/v3"

func AddRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Get("/signup", Register)

	auth.Get("/login", Login)

	auth.Get("/logout", Logout)

	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"uri":  ctx.Request().URI().String(),
			"path": ctx.Path(),
		})
	})
}
