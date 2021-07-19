package pquerna

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s Service) GenerateAtTime(secret string, time time.Time) string {
	code, _ := totp.GenerateCode(secret, time)
	return code
}

func (s Service) GenerateNow(secret string) string {
	code, _ := totp.GenerateCode(secret, time.Now())
	return code
}

func (s Service) Verify(secret, code string, time time.Time) bool {
	valid, err := totp.ValidateCustom(code, secret, time, totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})

	if err != nil {
		return false
	}

	return valid
}
