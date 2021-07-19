package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-totp-example/services/gotp"
	"go-totp-example/services/pquerna"
	"testing"
	"time"
)

func TestEqualCodes(t *testing.T) {
	s1 := gotp.New()
	s2 := pquerna.New()

	tm := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	secret := "3S4IISACJ2CMGWB5"

	assert.Equal(t, s1.GenerateAtTime(secret, tm), s2.GenerateAtTime(secret, tm))
	assert.Equal(t, s1.GenerateNow(secret), s2.GenerateNow(secret))
}

func TestServices(t *testing.T) {
	srvs := []OTPService{gotp.New(), pquerna.New()}

	type testCase struct {
		secret   string
		unixTime int
		code     string
	}

	cases := []testCase{
		{"WL5KLCWL3FGASBKW", 1262304000, "115296"},
		{"ZHFVDUJZYWKE", 1607749200, "328056"},
		{"WNE4MVPWTGYKLBSO75F4R2SNM2ZRGKVFJWCBZWOIOLM2O53YW6IEGELNXX44WHMM", 1178377810, "222382"},
	}

	for i := range srvs {
		for j, tc := range cases {
			t.Run(fmt.Sprintf("verify codes, srv%d case%d", i, j), func(t *testing.T) {
				t.Parallel()

				valid := srvs[i].Verify(tc.secret, tc.code, time.Unix(int64(tc.unixTime), 0))
				assert.True(t, valid)

				valid = srvs[i].Verify(tc.secret, tc.code, time.Unix(int64(tc.unixTime)+60, 0))
				assert.False(t, valid)
			})
		}
	}
}
