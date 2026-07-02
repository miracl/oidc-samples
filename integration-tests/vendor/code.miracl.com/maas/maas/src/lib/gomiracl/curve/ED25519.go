//go:build !ignoredefaultcurves || ED25519

package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// ED25519 elliptic curve.
var ED25519 = &amclCurve{
	name:           "ED25519",
	EGS:            bindings.EGSED25519,
	getKeyPair:     wrap.ECPKeyPairGenerateED25519,
	validatePubKey: wrap.ECPPublicKeyValidateED25519,
	sign:           wrap.ECPSpDsaED25519,
	verify:         wrap.ECPVpDsaED25519,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["ED25519"] = ED25519
}
