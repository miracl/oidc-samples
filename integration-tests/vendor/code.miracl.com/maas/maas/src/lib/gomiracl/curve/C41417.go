//go:build !ignoredefaultcurves || C41417

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// C41417 elliptic curve.
var C41417 = &amclCurve{
	name:           "C41417",
	EGS:            bindings.EGSC41417,
	getKeyPair:     wrap.ECPKeyPairGenerateC41417,
	validatePubKey: wrap.ECPPublicKeyValidateC41417,
	sign:           wrap.ECPSpDsaC41417,
	verify:         wrap.ECPVpDsaC41417,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["C41417"] = C41417
}
