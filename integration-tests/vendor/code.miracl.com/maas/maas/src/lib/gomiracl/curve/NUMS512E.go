//go:build !ignoredefaultcurves || NUMS512E

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// NUMS512E elliptic curve.
var NUMS512E = &amclCurve{
	name:           "NUMS512E",
	EGS:            bindings.EGSNUMS512E,
	getKeyPair:     wrap.ECPKeyPairGenerateNUMS512E,
	validatePubKey: wrap.ECPPublicKeyValidateNUMS512E,
	sign:           wrap.ECPSpDsaNUMS512E,
	verify:         wrap.ECPVpDsaNUMS512E,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["NUMS512E"] = NUMS512E
}
