package gotp

import (
	"github.com/xlzd/gotp"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) GenerateNow(secret string) string {
	return gotp.NewDefaultTOTP(secret).Now()
}

func (s *Service) GenerateAtTime(secret string, time time.Time) string {
	return gotp.NewDefaultTOTP(secret).At(int(time.Unix()))
}

func (s *Service) Verify(secret, code string, time time.Time) bool {
	return gotp.NewDefaultTOTP(secret).Verify(code, int(time.Unix()))
}
