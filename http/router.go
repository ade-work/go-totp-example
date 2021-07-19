package http

import (
	"github.com/gofiber/fiber/v2"
	"go-totp-example/services/gotp"
)

func NewRouter() *fiber.App {
	r := fiber.New()

	tc := NewTOTPController(gotp.New())

	r.Post("/totp/validate", tc.Validate)

	return r
}
