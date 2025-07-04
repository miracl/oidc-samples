// Generated by gen/wrappers/main.go from ecdsa_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || NUMS384W

package wrap

import "code.miracl.com/maas/maas/src/lib/gomiracl/bindings"

// ECPKeyPairGenerateNUMS384W is a wrapper of bindings.ECP_NUMS384W_KEY_PAIR_GENERATE.
func ECPKeyPairGenerateNUMS384W(R *bindings.Rand, s []byte) (sResult []byte, W []byte, err error) {

	var sOct *bindings.Octet
	if s != nil {
		sOct = bindings.NewOctet(s)
		defer sOct.Free()
	} else {
		sSize := bindings.EGSNUMS384W
		sOct = bindings.MakeOctet(sSize)
		defer sOct.Free()
	}

	WSize := 2*bindings.EFSNUMS384W + 1
	WOct := bindings.MakeOctet(WSize)
	defer WOct.Free()

	err = bindings.ECPKeyPairGenerateNUMS384W(R, sOct, WOct)

	sResult = sOct.ToBytes()

	W = WOct.ToBytes()

	return
}

// ECPPublicKeyValidateNUMS384W is a wrapper of bindings.ECP_NUMS384W_PUBLIC_KEY_VALIDATE.
func ECPPublicKeyValidateNUMS384W(W []byte) (err error) {

	WOct := bindings.NewOctet(W)
	defer WOct.Free()

	err = bindings.ECPPublicKeyValidateNUMS384W(WOct)

	return
}

// ECPSpDsaNUMS384W is a wrapper of bindings.ECP_NUMS384W_SP_DSA.
func ECPSpDsaNUMS384W(h int, R *bindings.Rand, key []byte, s []byte, M []byte) (c []byte, d []byte, err error) {

	keyOct := bindings.NewOctet(key)
	defer keyOct.Free()

	sOct := bindings.NewOctet(s)
	defer sOct.Free()

	MOct := bindings.NewOctet(M)
	defer MOct.Free()

	cSize := bindings.EGSNUMS384W
	cOct := bindings.MakeOctet(cSize)
	defer cOct.Free()

	dSize := bindings.EGSNUMS384W
	dOct := bindings.MakeOctet(dSize)
	defer dOct.Free()

	err = bindings.ECPSpDsaNUMS384W(h, R, keyOct, sOct, MOct, cOct, dOct)

	c = cOct.ToBytes()

	d = dOct.ToBytes()

	return
}

// ECPVpDsaNUMS384W is a wrapper of bindings.ECP_NUMS384W_VP_DSA.
func ECPVpDsaNUMS384W(h int, W []byte, M []byte, c []byte, d []byte) (err error) {

	WOct := bindings.NewOctet(W)
	defer WOct.Free()

	MOct := bindings.NewOctet(M)
	defer MOct.Free()

	cOct := bindings.NewOctet(c)
	defer cOct.Free()

	dOct := bindings.NewOctet(d)
	defer dOct.Free()

	err = bindings.ECPVpDsaNUMS384W(h, WOct, MOct, cOct, dOct)

	return
}
