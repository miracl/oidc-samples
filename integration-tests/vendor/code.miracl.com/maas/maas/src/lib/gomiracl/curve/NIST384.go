//go:build !ignoredefaultcurves || NIST384

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// NIST384 elliptic curve.
var NIST384 = &amclCurve{
	name:           "NIST384",
	EGS:            bindings.EGSNIST384,
	getKeyPair:     wrap.ECPKeyPairGenerateNIST384,
	validatePubKey: wrap.ECPPublicKeyValidateNIST384,
	sign:           wrap.ECPSpDsaNIST384,
	verify:         wrap.ECPVpDsaNIST384,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["NIST384"] = NIST384
}
