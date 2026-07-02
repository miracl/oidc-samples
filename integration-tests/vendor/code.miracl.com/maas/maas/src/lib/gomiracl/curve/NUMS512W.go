//go:build !ignoredefaultcurves || NUMS512W

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// NUMS512W elliptic curve.
var NUMS512W = &amclCurve{
	name:           "NUMS512W",
	EGS:            bindings.EGSNUMS512W,
	getKeyPair:     wrap.ECPKeyPairGenerateNUMS512W,
	validatePubKey: wrap.ECPPublicKeyValidateNUMS512W,
	sign:           wrap.ECPSpDsaNUMS512W,
	verify:         wrap.ECPVpDsaNUMS512W,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["NUMS512W"] = NUMS512W
}
