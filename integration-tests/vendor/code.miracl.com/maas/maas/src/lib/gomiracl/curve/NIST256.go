//go:build !ignoredefaultcurves || NIST256

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// NIST256 elliptic curve.
var NIST256 = &amclCurve{
	name:           "NIST256",
	EGS:            bindings.EGSNIST256,
	getKeyPair:     wrap.ECPKeyPairGenerateNIST256,
	validatePubKey: wrap.ECPPublicKeyValidateNIST256,
	sign:           wrap.ECPSpDsaNIST256,
	verify:         wrap.ECPVpDsaNIST256,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["NIST256"] = NIST256
}
