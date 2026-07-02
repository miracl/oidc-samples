//go:build !ignoredefaultcurves || NUMS256E

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// NUMS256E elliptic curve.
var NUMS256E = &amclCurve{
	name:           "NUMS256E",
	EGS:            bindings.EGSNUMS256E,
	getKeyPair:     wrap.ECPKeyPairGenerateNUMS256E,
	validatePubKey: wrap.ECPPublicKeyValidateNUMS256E,
	sign:           wrap.ECPSpDsaNUMS256E,
	verify:         wrap.ECPVpDsaNUMS256E,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["NUMS256E"] = NUMS256E
}
