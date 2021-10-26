// Generated by gen/wrappers/main.go from ecdsa_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || BN254

package wrap

import "code.miracl.com/mfa/pkg/gomiracl/bindings"

// ECPKeyPairGenerateBN254 is a wrapper of bindings.ECP_BN254_KEY_PAIR_GENERATE.
func ECPKeyPairGenerateBN254(R *bindings.Rand, s []byte) (sResult []byte, W []byte, err error) {

	var sOct *bindings.Octet
	if s != nil {
		sOct = bindings.NewOctet(s)
		defer sOct.Free()
	} else {
		sSize := bindings.EGSBN254
		sOct = bindings.MakeOctet(sSize)
		defer sOct.Free()
	}

	WSize := 2*bindings.EFSBN254 + 1
	WOct := bindings.MakeOctet(WSize)
	defer WOct.Free()

	err = bindings.ECPKeyPairGenerateBN254(R, sOct, WOct)

	sResult = sOct.ToBytes()

	W = WOct.ToBytes()

	return
}

// ECPPublicKeyValidateBN254 is a wrapper of bindings.ECP_BN254_PUBLIC_KEY_VALIDATE.
func ECPPublicKeyValidateBN254(W []byte) (err error) {

	WOct := bindings.NewOctet(W)
	defer WOct.Free()

	err = bindings.ECPPublicKeyValidateBN254(WOct)

	return
}

// ECPSpDsaBN254 is a wrapper of bindings.ECP_BN254_SP_DSA.
func ECPSpDsaBN254(h int, R *bindings.Rand, key []byte, s []byte, M []byte) (c []byte, d []byte, err error) {

	keyOct := bindings.NewOctet(key)
	defer keyOct.Free()

	sOct := bindings.NewOctet(s)
	defer sOct.Free()

	MOct := bindings.NewOctet(M)
	defer MOct.Free()

	cSize := bindings.EGSBN254
	cOct := bindings.MakeOctet(cSize)
	defer cOct.Free()

	dSize := bindings.EGSBN254
	dOct := bindings.MakeOctet(dSize)
	defer dOct.Free()

	err = bindings.ECPSpDsaBN254(h, R, keyOct, sOct, MOct, cOct, dOct)

	c = cOct.ToBytes()

	d = dOct.ToBytes()

	return
}

// ECPVpDsaBN254 is a wrapper of bindings.ECP_BN254_VP_DSA.
func ECPVpDsaBN254(h int, W []byte, M []byte, c []byte, d []byte) (err error) {

	WOct := bindings.NewOctet(W)
	defer WOct.Free()

	MOct := bindings.NewOctet(M)
	defer MOct.Free()

	cOct := bindings.NewOctet(c)
	defer cOct.Free()

	dOct := bindings.NewOctet(d)
	defer dOct.Free()

	err = bindings.ECPVpDsaBN254(h, WOct, MOct, cOct, dOct)

	return
}
