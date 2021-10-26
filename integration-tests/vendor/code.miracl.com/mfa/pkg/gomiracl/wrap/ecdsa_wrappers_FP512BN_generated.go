// Generated by gen/wrappers/main.go from ecdsa_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || FP512BN

package wrap

import "code.miracl.com/mfa/pkg/gomiracl/bindings"

// ECPKeyPairGenerateFP512BN is a wrapper of bindings.ECP_FP512BN_KEY_PAIR_GENERATE.
func ECPKeyPairGenerateFP512BN(R *bindings.Rand, s []byte) (sResult []byte, W []byte, err error) {

	var sOct *bindings.Octet
	if s != nil {
		sOct = bindings.NewOctet(s)
		defer sOct.Free()
	} else {
		sSize := bindings.EGSFP512BN
		sOct = bindings.MakeOctet(sSize)
		defer sOct.Free()
	}

	WSize := 2*bindings.EFSFP512BN + 1
	WOct := bindings.MakeOctet(WSize)
	defer WOct.Free()

	err = bindings.ECPKeyPairGenerateFP512BN(R, sOct, WOct)

	sResult = sOct.ToBytes()

	W = WOct.ToBytes()

	return
}

// ECPPublicKeyValidateFP512BN is a wrapper of bindings.ECP_FP512BN_PUBLIC_KEY_VALIDATE.
func ECPPublicKeyValidateFP512BN(W []byte) (err error) {

	WOct := bindings.NewOctet(W)
	defer WOct.Free()

	err = bindings.ECPPublicKeyValidateFP512BN(WOct)

	return
}

// ECPSpDsaFP512BN is a wrapper of bindings.ECP_FP512BN_SP_DSA.
func ECPSpDsaFP512BN(h int, R *bindings.Rand, key []byte, s []byte, M []byte) (c []byte, d []byte, err error) {

	keyOct := bindings.NewOctet(key)
	defer keyOct.Free()

	sOct := bindings.NewOctet(s)
	defer sOct.Free()

	MOct := bindings.NewOctet(M)
	defer MOct.Free()

	cSize := bindings.EGSFP512BN
	cOct := bindings.MakeOctet(cSize)
	defer cOct.Free()

	dSize := bindings.EGSFP512BN
	dOct := bindings.MakeOctet(dSize)
	defer dOct.Free()

	err = bindings.ECPSpDsaFP512BN(h, R, keyOct, sOct, MOct, cOct, dOct)

	c = cOct.ToBytes()

	d = dOct.ToBytes()

	return
}

// ECPVpDsaFP512BN is a wrapper of bindings.ECP_FP512BN_VP_DSA.
func ECPVpDsaFP512BN(h int, W []byte, M []byte, c []byte, d []byte) (err error) {

	WOct := bindings.NewOctet(W)
	defer WOct.Free()

	MOct := bindings.NewOctet(M)
	defer MOct.Free()

	cOct := bindings.NewOctet(c)
	defer cOct.Free()

	dOct := bindings.NewOctet(d)
	defer dOct.Free()

	err = bindings.ECPVpDsaFP512BN(h, WOct, MOct, cOct, dOct)

	return
}
