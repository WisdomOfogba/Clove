package routes

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/chibx/vendor-pulse/api/middlewares"
	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/gofiber/fiber/v3"
)

var logger = global.Logger

func validateFileType(file *multipart.FileHeader, allowed []string) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	// Read the first 512 bytes for MIME detection
	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil {
		return err
	}

	mimeType := http.DetectContentType(buf[:n])

	if slices.Contains(allowed, mimeType) {
		return nil
	}

	return fmt.Errorf("file type %s is not allowed", mimeType)
}

func AddRoutes(app *fiber.App) {
	api := app.Group("/api", middlewares.EnsureDeviceIDIsSet, middlewares.AuthMiddleware, middlewares.HardenBackendEndpoint)
	auth := api.Group("/auth")
	auth.Get("/get-identity",GetUserIdentity())
	auth.Post("/signup", UserRegister())
	auth.Post("/login", UserLogin())
	auth.Post("/logout", UserLogout())
	api.Post("/promote-account", UpgradeToBusinessAccount())

	// -------------------MEAL------------------------------
	api.Post("/meal", CreateMeal())
	api.Post("/meal/delete-media", DeleteMealMedia())
	api.Put("/meal/:id", UpdateMeal())
	api.Delete("/meal/:id", DeleteMeal())

	// -------------------ORDER-------------------------------
	api.Post("/place-order", PlaceOrder())
	api.Post("/cancel-order", CancelOrder())

	// --------------------REVIEWS----------------------------
	api.Post("/meal/:id/review", AddReview())
	api.Post("/review/edit", EditReview())

	app.Get("*", func(ctx fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"uri":  ctx.Request().URI().String(),
			"path": ctx.Path(),
		})
	})
}
