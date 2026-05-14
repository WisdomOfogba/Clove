package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/chibx/vendor-pulse/api/routes"
	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	_ "github.com/joho/godotenv/autoload"
)

var isShuttingDown atomic.Bool

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func getApp() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableHeadAutoRegister: true,
		StructValidator:         &structValidator{validate: global.Validator},
	})

	app.Get("/readyz", func(c fiber.Ctx) error {
		if isShuttingDown.Load() {
			return c.SendStatus(fiber.StatusServiceUnavailable)
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Hooks().OnPostShutdown(func(e error) error {
		// global.DB.Close()
		err := global.Redis.Close()
		return err
	})

	routes.AddRoutes(app)

	return app
}

func main() {
	global.InitGlobals()
	app := getApp()
	/*
			mux := http.NewServeMux()
		  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		    fmt.Fprintln(w, "Hello from Go on Vercel")
		  })

		  port := os.Getenv("PORT")
		  if port == "" {
		    port = "3000"
		  }

		  log.Fatal(http.ListenAndServe(":"+port, mux))
	*/

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Printf("Listen error: %v", err)
		}
	}()

	// Wait for signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Begin graceful shutdown
	isShuttingDown.Store(true)
	time.Sleep(1 * time.Second) // Let Loadbalancer drain

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
