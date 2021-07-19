package http

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"go-totp-example/services"
	"go-totp-example/utils"
)

type TOTPController struct {
	otpService services.OTPService
	validate   *validator.Validate
}

func NewTOTPController(optService services.OTPService) *TOTPController {
	return &TOTPController{
		otpService: optService,
		validate:   validator.New(),
	}
}

func (tc *TOTPController) Validate(ctx *fiber.Ctx) error {
	req := new(Request)

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := tc.validate.Struct(*req); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !tc.otpService.Verify(utils.UserSecret, req.Code, req.Time) {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "invalid totp code",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
