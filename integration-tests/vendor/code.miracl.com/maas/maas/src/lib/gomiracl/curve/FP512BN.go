//go:build !ignoredefaultcurves || FP512BN

//nolint:dupl,revive // definitions
package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// FP512BN elliptic curve.
var FP512BN = &mpinCurve{
	Curve: &amclCurve{
		name:           "FP512BN",
		EGS:            bindings.EGSFP512BN,
		getKeyPair:     wrap.ECPKeyPairGenerateFP512BN,
		validatePubKey: wrap.ECPPublicKeyValidateFP512BN,
		sign:           wrap.ECPSpDsaFP512BN,
		verify:         wrap.ECPVpDsaFP512BN,
	},
	PFS:                  bindings.PFSFP512BN,
	PGS:                  bindings.PGSFP512BN,
	getRandomSecret:      wrap.RandomGenerateFP512BN,
	newServerSecretShare: wrap.GetServerSecretFP512BN,
	newClientSecretShare: wrap.GetClientSecretFP512BN,
	newTimePermitShare:   wrap.GetClientPermitFP512BN,
	recombineG1:          wrap.RecombineG1FP512BN,
	recombineG2:          wrap.RecombineG2FP512BN,
	extractPin:           wrap.ExtractPINFP512BN,
	clientOnePass:        wrap.ClientFP512BN,
	serverOnePass:        wrap.ServerFP512BN,
	clientPass1:          wrap.Client1FP512BN,
	clientPass2:          wrap.Client2FP512BN,
	serverPass1:          wrap.Server1FP512BN,
	serverPass2:          wrap.Server2FP512BN,
	precompute:           wrap.PrecomputeFP512BN,
	getG1Multiple:        wrap.GetG1MultipleFP512BN,
	serverKey:            wrap.ServerKeyFP512BN,
	clientKey:            wrap.ClientKeyFP512BN,
	getDVSKeyPair:        wrap.GetDVSKeyPairFP512BN,

	// Wang Chow-Choo
	wccRandomGenerate: wrap.WCCRandomGenerateFP512BN,
	wccHq:             wrap.WCCHqFP512BN,
	wccGetG1Multiple:  wrap.WCCGetG1MultipleFP512BN,
	wccRecombineG1:    wrap.WCCRecombineG1FP512BN,
	wccGetG2Multiple:  wrap.WCCGetG2MultipleFP512BN,
	wccRecombineG2:    wrap.WCCRecombineG2FP512BN,
	wccReceiverKey:    wrap.WCCReceiverKeyFP512BN,
	wccSenderKey:      wrap.WCCSenderKeyFP512BN,
}

//nolint:gochecknoinits // It is how building with limited curve set works.
func init() {
	All["FP512BN"] = FP512BN
}
