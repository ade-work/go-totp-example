package router

import "github.com/gofiber/fiber/v2"

func New() *fiber.App {
	r := fiber.New()

	r.Post("/totp/validate", validateHandler)

	return r
}

func validateHandler(ctx *fiber.Ctx) error {
	return nil
}