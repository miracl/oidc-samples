// Generated by gen/wrappers/main.go from ecdsa_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || NIST521

package wrap

import "code.miracl.com/maas/maas/src/lib/gomiracl/bindings"

// ECPKeyPairGenerateNIST521 is a wrapper of bindings.ECP_NIST521_KEY_PAIR_GENERATE.
func ECPKeyPairGenerateNIST521(R *bindings.Rand, s []byte) (sResult []byte, W []byte, err error) {

	var sOct *bindings.Octet
	if s != nil {
		sOct = bindings.NewOctet(s)
		defer sOct.Free()
	} else {
		sSize := bindings.EGSNIST521
		sOct = bindings.MakeOctet(sSize)
		defer sOct.Free()
	}

	WSize := 2*bindings.EFSNIST521 + 1
	WOct := bindings.MakeOctet(WSize)
	defer WOct.Free()

	err = bindings.ECPKeyPairGenerateNIST521(R, sOct, WOct)

	sResult = sOct.ToBytes()

	W = WOct.ToBytes()

	return
}

// ECPPublicKeyValidateNIST521 is a wrapper of bindings.ECP_NIST521_PUBLIC_KEY_VALIDATE.
func ECPPublicKeyValidateNIST521(W []byte) (err error) {

	WOct := bindings.NewOctet(W)
	defer WOct.Free()

	err = bindings.ECPPublicKeyValidateNIST521(WOct)

	return
}

// ECPSpDsaNIST521 is a wrapper of bindings.ECP_NIST521_SP_DSA.
func ECPSpDsaNIST521(h int, R *bindings.Rand, key []byte, s []byte, M []byte) (c []byte, d []byte, err error) {

	keyOct := bindings.NewOctet(key)
	defer keyOct.Free()

	sOct := bindings.NewOctet(s)
	defer sOct.Free()

	MOct := bindings.NewOctet(M)
	defer MOct.Free()

	cSize := bindings.EGSNIST521
	cOct := bindings.MakeOctet(cSize)
	defer cOct.Free()

	dSize := bindings.EGSNIST521
	dOct := bindings.MakeOctet(dSize)
	defer dOct.Free()

	err = bindings.ECPSpDsaNIST521(h, R, keyOct, sOct, MOct, cOct, dOct)

	c = cOct.ToBytes()

	d = dOct.ToBytes()

	return
}

// ECPVpDsaNIST521 is a wrapper of bindings.ECP_NIST521_VP_DSA.
func ECPVpDsaNIST521(h int, W []byte, M []byte, c []byte, d []byte) (err error) {

	WOct := bindings.NewOctet(W)
	defer WOct.Free()

	MOct := bindings.NewOctet(M)
	defer MOct.Free()

	cOct := bindings.NewOctet(c)
	defer cOct.Free()

	dOct := bindings.NewOctet(d)
	defer dOct.Free()

	err = bindings.ECPVpDsaNIST521(h, WOct, MOct, cOct, dOct)

	return
}
