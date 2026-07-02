//go:build !ignoredefaultcurves || FP256BN

//nolint:dupl,revive // definitions
package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// FP256BN elliptic curve.
var FP256BN = &mpinCurve{
	Curve: &amclCurve{
		name:           "FP256BN",
		EGS:            bindings.EGSFP256BN,
		getKeyPair:     wrap.ECPKeyPairGenerateFP256BN,
		validatePubKey: wrap.ECPPublicKeyValidateFP256BN,
		sign:           wrap.ECPSpDsaFP256BN,
		verify:         wrap.ECPVpDsaFP256BN,
	},
	PFS:                  bindings.PFSFP256BN,
	PGS:                  bindings.PGSFP256BN,
	getRandomSecret:      wrap.RandomGenerateFP256BN,
	newServerSecretShare: wrap.GetServerSecretFP256BN,
	newClientSecretShare: wrap.GetClientSecretFP256BN,
	newTimePermitShare:   wrap.GetClientPermitFP256BN,
	recombineG1:          wrap.RecombineG1FP256BN,
	recombineG2:          wrap.RecombineG2FP256BN,
	extractPin:           wrap.ExtractPINFP256BN,
	clientOnePass:        wrap.ClientFP256BN,
	serverOnePass:        wrap.ServerFP256BN,
	clientPass1:          wrap.Client1FP256BN,
	clientPass2:          wrap.Client2FP256BN,
	serverPass1:          wrap.Server1FP256BN,
	serverPass2:          wrap.Server2FP256BN,
	precompute:           wrap.PrecomputeFP256BN,
	getG1Multiple:        wrap.GetG1MultipleFP256BN,
	serverKey:            wrap.ServerKeyFP256BN,
	clientKey:            wrap.ClientKeyFP256BN,
	getDVSKeyPair:        wrap.GetDVSKeyPairFP256BN,

	// Wang Chow-Choo
	wccRandomGenerate: wrap.WCCRandomGenerateFP256BN,
	wccHq:             wrap.WCCHqFP256BN,
	wccGetG1Multiple:  wrap.WCCGetG1MultipleFP256BN,
	wccRecombineG1:    wrap.WCCRecombineG1FP256BN,
	wccGetG2Multiple:  wrap.WCCGetG2MultipleFP256BN,
	wccRecombineG2:    wrap.WCCRecombineG2FP256BN,
	wccReceiverKey:    wrap.WCCReceiverKeyFP256BN,
	wccSenderKey:      wrap.WCCSenderKeyFP256BN,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["FP256BN"] = FP256BN
}
