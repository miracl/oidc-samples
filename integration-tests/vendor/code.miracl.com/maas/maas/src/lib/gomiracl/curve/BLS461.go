//go:build !ignoredefaultcurves || BLS461

//nolint:dupl,revive // definitions
package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// BLS461 elliptic curve.
var BLS461 = &mpinCurve{
	Curve: &amclCurve{
		name:           "BLS461",
		EGS:            bindings.EGSBLS461,
		getKeyPair:     wrap.ECPKeyPairGenerateBLS461,
		validatePubKey: wrap.ECPPublicKeyValidateBLS461,
		sign:           wrap.ECPSpDsaBLS461,
		verify:         wrap.ECPVpDsaBLS461,
	},
	PFS:                  bindings.PFSBLS461,
	PGS:                  bindings.PGSBLS461,
	getRandomSecret:      wrap.RandomGenerateBLS461,
	newServerSecretShare: wrap.GetServerSecretBLS461,
	newClientSecretShare: wrap.GetClientSecretBLS461,
	newTimePermitShare:   wrap.GetClientPermitBLS461,
	recombineG1:          wrap.RecombineG1BLS461,
	recombineG2:          wrap.RecombineG2BLS461,
	extractPin:           wrap.ExtractPINBLS461,
	clientOnePass:        wrap.ClientBLS461,
	serverOnePass:        wrap.ServerBLS461,
	clientPass1:          wrap.Client1BLS461,
	clientPass2:          wrap.Client2BLS461,
	serverPass1:          wrap.Server1BLS461,
	serverPass2:          wrap.Server2BLS461,
	precompute:           wrap.PrecomputeBLS461,
	getG1Multiple:        wrap.GetG1MultipleBLS461,
	serverKey:            wrap.ServerKeyBLS461,
	clientKey:            wrap.ClientKeyBLS461,
	getDVSKeyPair:        wrap.GetDVSKeyPairBLS461,

	// Wang Chow-Choo
	wccRandomGenerate: wrap.WCCRandomGenerateBLS461,
	wccHq:             wrap.WCCHqBLS461,
	wccGetG1Multiple:  wrap.WCCGetG1MultipleBLS461,
	wccRecombineG1:    wrap.WCCRecombineG1BLS461,
	wccGetG2Multiple:  wrap.WCCGetG2MultipleBLS461,
	wccRecombineG2:    wrap.WCCRecombineG2BLS461,
	wccReceiverKey:    wrap.WCCReceiverKeyBLS461,
	wccSenderKey:      wrap.WCCSenderKeyBLS461,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["BLS461"] = BLS461
}
