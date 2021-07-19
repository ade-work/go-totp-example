package services

import "time"

type OTPService interface {
	GenerateAtTime(secret string, time time.Time) string
	GenerateNow(secret string) string
	Verify(secret, code string, time time.Time) bool
}
