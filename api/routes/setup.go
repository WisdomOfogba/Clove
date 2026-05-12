package routes

import "github.com/gofiber/fiber/v3"

func AddRoutes(app *fiber.App) {
	customerAuth := app.Group("/auth/customer")
	customerAuth.Post("/signup", CustomerRegister())
	customerAuth.Post("/login", CustomerLogin())
	customerAuth.Post("/logout", CustomerLogout())
	
	sellerAuth := app.Group("/auth/seller")
	sellerAuth.Post("/signup", CustomerRegister())
	sellerAuth.Post("/login", CustomerLogin())
	sellerAuth.Post("/logout", CustomerLogout())

	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"uri":  ctx.Request().URI().String(),
			"path": ctx.Path(),
		})
	})
}
