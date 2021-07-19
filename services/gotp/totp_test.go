package gotp

import (
	"github.com/xlzd/gotp"
	"os"
	"testing"
	"time"
)

var totpService *Service

func TestMain(m *testing.M) {
	totpService = New()

	code := m.Run()
	os.Exit(code)
}

func TestService_CreateAtTime(t *testing.T) {

	a := totpService.GenerateAtTime("asf", time.Now())
	println(a)
}

func TestExpireAt(t *testing.T) {
	secret := "J5O7QRQY7BMEVKZE"

	totp := gotp.NewTOTP(secret, 5, 600, nil)

	tn := int(time.Now().Unix())

	code := totp.At(tn)
	println(code)

	code2 := totp.At(tn + 40)
	println(code2)
}
