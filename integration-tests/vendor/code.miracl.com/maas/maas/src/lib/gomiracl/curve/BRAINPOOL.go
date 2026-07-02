//go:build !ignoredefaultcurves || BRAINPOOL

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// BRAINPOOL elliptic curve.
var BRAINPOOL = &amclCurve{
	name:           "BRAINPOOL",
	EGS:            bindings.EGSBRAINPOOL,
	getKeyPair:     wrap.ECPKeyPairGenerateBRAINPOOL,
	validatePubKey: wrap.ECPPublicKeyValidateBRAINPOOL,
	sign:           wrap.ECPSpDsaBRAINPOOL,
	verify:         wrap.ECPVpDsaBRAINPOOL,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["BRAINPOOL"] = BRAINPOOL
}
