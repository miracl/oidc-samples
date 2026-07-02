//go:build !ignoredefaultcurves || NUMS384W

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// NUMS384W elliptic curve.
var NUMS384W = &amclCurve{
	name:           "NUMS384W",
	EGS:            bindings.EGSNUMS384W,
	getKeyPair:     wrap.ECPKeyPairGenerateNUMS384W,
	validatePubKey: wrap.ECPPublicKeyValidateNUMS384W,
	sign:           wrap.ECPSpDsaNUMS384W,
	verify:         wrap.ECPVpDsaNUMS384W,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["NUMS384W"] = NUMS384W
}
