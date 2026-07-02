package mpin

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/rand"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// GenerateOTP generated new 6 digit OTP.
func GenerateOTP(rng *rand.Rand) int {
	return wrap.GenerateOTP((*bindings.Rand)(rng))
}
