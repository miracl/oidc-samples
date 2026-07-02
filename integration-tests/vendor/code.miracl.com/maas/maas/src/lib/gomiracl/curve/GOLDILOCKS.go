//go:build !ignoredefaultcurves || GOLDILOCKS

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// GOLDILOCKS elliptic curve.
var GOLDILOCKS = &amclCurve{
	name:           "GOLDILOCKS",
	EGS:            bindings.EGSGOLDILOCKS,
	getKeyPair:     wrap.ECPKeyPairGenerateGOLDILOCKS,
	validatePubKey: wrap.ECPPublicKeyValidateGOLDILOCKS,
	sign:           wrap.ECPSpDsaGOLDILOCKS,
	verify:         wrap.ECPVpDsaGOLDILOCKS,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["GOLDILOCKS"] = GOLDILOCKS
}
