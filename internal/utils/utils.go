package utils

import "github.com/gofiber/fiber/v3"

func AsPointer[T any](a T) *T {
	return &a
}

func ClearCookies(ctx fiber.Ctx, cookies ...string) {
	for _, v := range cookies {
		ctx.ClearCookie(v)
	}
}
