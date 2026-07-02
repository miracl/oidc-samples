//go:build !ignoredefaultcurves || BLS383

//nolint:dupl // definitions
package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// BLS383 elliptic curve.
var BLS383 = &mpinCurve{
	Curve: &amclCurve{
		name:           "BLS383",
		EGS:            bindings.EGSBLS383,
		getKeyPair:     wrap.ECPKeyPairGenerateBLS383,
		validatePubKey: wrap.ECPPublicKeyValidateBLS383,
		sign:           wrap.ECPSpDsaBLS383,
		verify:         wrap.ECPVpDsaBLS383,
	},
	PFS:                  bindings.PFSBLS383,
	PGS:                  bindings.PGSBLS383,
	getRandomSecret:      wrap.RandomGenerateBLS383,
	newServerSecretShare: wrap.GetServerSecretBLS383,
	newClientSecretShare: wrap.GetClientSecretBLS383,
	newTimePermitShare:   wrap.GetClientPermitBLS383,
	recombineG1:          wrap.RecombineG1BLS383,
	recombineG2:          wrap.RecombineG2BLS383,
	extractPin:           wrap.ExtractPINBLS383,
	clientOnePass:        wrap.ClientBLS383,
	serverOnePass:        wrap.ServerBLS383,
	clientPass1:          wrap.Client1BLS383,
	clientPass2:          wrap.Client2BLS383,
	serverPass1:          wrap.Server1BLS383,
	serverPass2:          wrap.Server2BLS383,
	precompute:           wrap.PrecomputeBLS383,
	getG1Multiple:        wrap.GetG1MultipleBLS383,
	serverKey:            wrap.ServerKeyBLS383,
	clientKey:            wrap.ClientKeyBLS383,
	getDVSKeyPair:        wrap.GetDVSKeyPairBLS383,

	// Wang Chow-Choo
	wccRandomGenerate: wrap.WCCRandomGenerateBLS383,
	wccHq:             wrap.WCCHqBLS383,
	wccGetG1Multiple:  wrap.WCCGetG1MultipleBLS383,
	wccRecombineG1:    wrap.WCCRecombineG1BLS383,
	wccGetG2Multiple:  wrap.WCCGetG2MultipleBLS383,
	wccRecombineG2:    wrap.WCCRecombineG2BLS383,
	wccReceiverKey:    wrap.WCCReceiverKeyBLS383,
	wccSenderKey:      wrap.WCCSenderKeyBLS383,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["BLS383"] = BLS383
}
