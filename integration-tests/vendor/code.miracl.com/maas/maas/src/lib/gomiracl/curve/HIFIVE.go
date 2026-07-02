//go:build !ignoredefaultcurves || HIFIVE

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// HIFIVE elliptic curve.
var HIFIVE = &amclCurve{
	name:           "HIFIVE",
	EGS:            bindings.EGSHIFIVE,
	getKeyPair:     wrap.ECPKeyPairGenerateHIFIVE,
	validatePubKey: wrap.ECPPublicKeyValidateHIFIVE,
	sign:           wrap.ECPSpDsaHIFIVE,
	verify:         wrap.ECPVpDsaHIFIVE,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["HIFIVE"] = HIFIVE
}
