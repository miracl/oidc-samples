//go:build !ignoredefaultcurves || BN254CX

//nolint:dupl // definitions
package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// BN254CX elliptic curve.
var BN254CX = &mpinCurve{
	Curve: &amclCurve{
		name:           "BN254CX",
		EGS:            bindings.EGSBN254CX,
		getKeyPair:     wrap.ECPKeyPairGenerateBN254CX,
		validatePubKey: wrap.ECPPublicKeyValidateBN254CX,
		sign:           wrap.ECPSpDsaBN254CX,
		verify:         wrap.ECPVpDsaBN254CX,
	},
	PFS:                  bindings.PFSBN254CX,
	PGS:                  bindings.PGSBN254CX,
	getRandomSecret:      wrap.RandomGenerateBN254CX,
	newServerSecretShare: wrap.GetServerSecretBN254CX,
	newClientSecretShare: wrap.GetClientSecretBN254CX,
	newTimePermitShare:   wrap.GetClientPermitBN254CX,
	recombineG1:          wrap.RecombineG1BN254CX,
	recombineG2:          wrap.RecombineG2BN254CX,
	extractPin:           wrap.ExtractPINBN254CX,
	clientOnePass:        wrap.ClientBN254CX,
	serverOnePass:        wrap.ServerBN254CX,
	clientPass1:          wrap.Client1BN254CX,
	clientPass2:          wrap.Client2BN254CX,
	serverPass1:          wrap.Server1BN254CX,
	serverPass2:          wrap.Server2BN254CX,
	precompute:           wrap.PrecomputeBN254CX,
	getG1Multiple:        wrap.GetG1MultipleBN254CX,
	serverKey:            wrap.ServerKeyBN254CX,
	clientKey:            wrap.ClientKeyBN254CX,
	getDVSKeyPair:        wrap.GetDVSKeyPairBN254CX,

	// Wang Chow-Choo
	wccRandomGenerate: wrap.WCCRandomGenerateBN254CX,
	wccHq:             wrap.WCCHqBN254CX,
	wccGetG1Multiple:  wrap.WCCGetG1MultipleBN254CX,
	wccRecombineG1:    wrap.WCCRecombineG1BN254CX,
	wccGetG2Multiple:  wrap.WCCGetG2MultipleBN254CX,
	wccRecombineG2:    wrap.WCCRecombineG2BN254CX,
	wccReceiverKey:    wrap.WCCReceiverKeyBN254CX,
	wccSenderKey:      wrap.WCCSenderKeyBN254CX,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["BN254CX"] = BN254CX
}
