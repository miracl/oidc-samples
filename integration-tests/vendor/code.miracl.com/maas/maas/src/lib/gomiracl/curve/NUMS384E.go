//go:build !ignoredefaultcurves || NUMS384E

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// NUMS384E elliptic curve.
var NUMS384E = &amclCurve{
	name:           "NUMS384E",
	EGS:            bindings.EGSNUMS384E,
	getKeyPair:     wrap.ECPKeyPairGenerateNUMS384E,
	validatePubKey: wrap.ECPPublicKeyValidateNUMS384E,
	sign:           wrap.ECPSpDsaNUMS384E,
	verify:         wrap.ECPVpDsaNUMS384E,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["NUMS384E"] = NUMS384E
}
