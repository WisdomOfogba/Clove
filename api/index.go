package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/chibx/vendor-pulse/api/routes"
	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	_ "github.com/joho/godotenv/autoload"
)

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

var isShuttingDown atomic.Bool

// Handler is the main entry point of the application. Think of it like the main() method
func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `fiber.Ctx`
	r.RequestURI = r.URL.String()

	app, httpHandler := handler()
	go func() {
		httpHandler.ServeHTTP(w, r)
	}()

	// Wait for signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Begin graceful shutdown
	isShuttingDown.Store(true)
	time.Sleep(3 * time.Second) // Let Loadbalancer drain

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	done := make(chan error, 1)
	go func() { done <- app.Shutdown() }()

	select {
	case <-done:
		log.Println("Shutdown complete")
	case <-ctx.Done():
		log.Println("Shutdown timed out")
	}
}

// building the fiber application
func handler() (*fiber.App, http.HandlerFunc) {
	global.InitGlobals()
	app := getApp()

	return app, adaptor.FiberApp(app)
}

func getApp() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableHeadAutoRegister: true,
		StructValidator:         &structValidator{validate: global.Validator},
	})

	app.Hooks().OnPostShutdown(func(e error) error {
		global.DB.Close()
		err := global.Redis.Close()
		return err
	})

	app.Get("/readyz", func(c fiber.Ctx) error {
		if isShuttingDown.Load() {
			return c.SendStatus(fiber.StatusServiceUnavailable)
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Use(helmet.New(helmet.Config{
		HSTSMaxAge:         63072000, // 2 years in seconds
		HSTSPreloadEnabled: true,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{""},
		AllowMethods: []string{""},
		AllowHeaders: []string{""},
		MaxAge:       3600,
	}))

	routes.AddRoutes(app)

	return app
}
