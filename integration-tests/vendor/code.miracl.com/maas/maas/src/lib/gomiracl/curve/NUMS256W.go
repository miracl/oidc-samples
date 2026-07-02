//go:build !ignoredefaultcurves || NUMS256W

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// NUMS256W elliptic curve.
var NUMS256W = &amclCurve{
	name:           "NUMS256W",
	EGS:            bindings.EGSNUMS256W,
	getKeyPair:     wrap.ECPKeyPairGenerateNUMS256W,
	validatePubKey: wrap.ECPPublicKeyValidateNUMS256W,
	sign:           wrap.ECPSpDsaNUMS256W,
	verify:         wrap.ECPVpDsaNUMS256W,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["NUMS256W"] = NUMS256W
}
