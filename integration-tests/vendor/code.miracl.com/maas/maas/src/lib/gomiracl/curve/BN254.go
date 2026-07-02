//go:build !ignoredefaultcurves || BN254

//nolint:dupl,revive // definitions
package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// BN254 elliptic curve.
var BN254 = &mpinCurve{
	Curve: &amclCurve{
		name:           "BN254",
		EGS:            bindings.EGSBN254,
		getKeyPair:     wrap.ECPKeyPairGenerateBN254,
		validatePubKey: wrap.ECPPublicKeyValidateBN254,
		sign:           wrap.ECPSpDsaBN254,
		verify:         wrap.ECPVpDsaBN254,
	},
	PFS:                  bindings.PFSBN254,
	PGS:                  bindings.PGSBN254,
	getRandomSecret:      wrap.RandomGenerateBN254,
	newServerSecretShare: wrap.GetServerSecretBN254,
	newClientSecretShare: wrap.GetClientSecretBN254,
	newTimePermitShare:   wrap.GetClientPermitBN254,
	recombineG1:          wrap.RecombineG1BN254,
	recombineG2:          wrap.RecombineG2BN254,
	extractPin:           wrap.ExtractPINBN254,
	clientOnePass:        wrap.ClientBN254,
	serverOnePass:        wrap.ServerBN254,
	clientPass1:          wrap.Client1BN254,
	clientPass2:          wrap.Client2BN254,
	serverPass1:          wrap.Server1BN254,
	serverPass2:          wrap.Server2BN254,
	precompute:           wrap.PrecomputeBN254,
	getG1Multiple:        wrap.GetG1MultipleBN254,
	serverKey:            wrap.ServerKeyBN254,
	clientKey:            wrap.ClientKeyBN254,
	getDVSKeyPair:        wrap.GetDVSKeyPairBN254,

	// Wang Chow-Choo
	wccRandomGenerate: wrap.WCCRandomGenerateBN254,
	wccHq:             wrap.WCCHqBN254,
	wccGetG1Multiple:  wrap.WCCGetG1MultipleBN254,
	wccRecombineG1:    wrap.WCCRecombineG1BN254,
	wccGetG2Multiple:  wrap.WCCGetG2MultipleBN254,
	wccRecombineG2:    wrap.WCCRecombineG2BN254,
	wccReceiverKey:    wrap.WCCReceiverKeyBN254,
	wccSenderKey:      wrap.WCCSenderKeyBN254,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["BN254"] = BN254
}
