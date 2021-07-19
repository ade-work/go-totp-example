package http

import (
	"fmt"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"go-totp-example/services/gotp"
)

func Test_Validate_NoTime(t *testing.T) {
	t.Parallel()
	s := gotp.New()

	body := `{"code": "124111"}`

	c := testCtx(body)

	assert.Nil(t, NewTOTPController(s).Validate(c))
	code, msg := getResponseResult(c)

	assert.Equal(t, code, 422)
	assert.Contains(t, msg, "failed on the 'required' tag")
}

func Test_Validate_InvalidCodeLength(t *testing.T) {
	t.Parallel()
	s := gotp.New()

	body := fmt.Sprintf(`{"code": "123", "time": "%s"}`, time.Now().Format(time.RFC3339))

	c := testCtx(body)

	assert.Nil(t, NewTOTPController(s).Validate(c))
	code, msg := getResponseResult(c)

	assert.Equal(t, code, 422)
	assert.Contains(t, msg, "failed on the 'len' tag")
}

func Test_Validate_InvalidCode(t *testing.T) {
	t.Parallel()
	s := gotp.New()

	tm := time.Date(2007, 1, 1, 0, 0, 0, 0, time.UTC)

	body := fmt.Sprintf(`{"code": "111111", "time": "%s"}`, tm.Format(time.RFC3339))

	c := testCtx(body)

	assert.Nil(t, NewTOTPController(s).Validate(c))
	code, msg := getResponseResult(c)

	assert.Equal(t, code, 422)
	assert.Contains(t, msg, "failed on the 'len' tag")
}

func testCtx(body string) *fiber.Ctx {
	app := fiber.New()

	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	c.Request().Header.SetContentType(fiber.MIMEApplicationJSON)

	if body == "" {
		return c
	}

	c.Request().SetBody([]byte(body))

	return c
}

func getResponseResult(ctx *fiber.Ctx) (code int, msg string) {
	code = ctx.Response().StatusCode()
	msg = string(ctx.Response().Body())

	return
}
