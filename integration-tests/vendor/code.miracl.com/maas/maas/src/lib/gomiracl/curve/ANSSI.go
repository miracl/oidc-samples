//go:build !ignoredefaultcurves || ANSSI

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// ANSSI elliptic curve.
var ANSSI = &amclCurve{
	name:           "ANSSI",
	EGS:            bindings.EGSANSSI,
	getKeyPair:     wrap.ECPKeyPairGenerateANSSI,
	validatePubKey: wrap.ECPPublicKeyValidateANSSI,
	sign:           wrap.ECPSpDsaANSSI,
	verify:         wrap.ECPVpDsaANSSI,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["ANSSI"] = ANSSI
}
