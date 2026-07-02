//go:build !ignoredefaultcurves || NIST521

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// NIST521 elliptic curve.
var NIST521 = &amclCurve{
	name:           "NIST521",
	EGS:            bindings.EGSNIST521,
	getKeyPair:     wrap.ECPKeyPairGenerateNIST521,
	validatePubKey: wrap.ECPPublicKeyValidateNIST521,
	sign:           wrap.ECPSpDsaNIST521,
	verify:         wrap.ECPVpDsaNIST521,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["NIST521"] = NIST521
}
