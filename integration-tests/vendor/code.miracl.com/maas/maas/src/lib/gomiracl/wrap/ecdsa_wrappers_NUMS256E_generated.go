// Generated by gen/wrappers/main.go from ecdsa_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || NUMS256E

package wrap

import "code.miracl.com/maas/maas/src/lib/gomiracl/bindings"

// ECPKeyPairGenerateNUMS256E is a wrapper of bindings.ECP_NUMS256E_KEY_PAIR_GENERATE.
func ECPKeyPairGenerateNUMS256E(R *bindings.Rand, s []byte) (sResult []byte, W []byte, err error) {

	var sOct *bindings.Octet
	if s != nil {
		sOct = bindings.NewOctet(s)
		defer sOct.Free()
	} else {
		sSize := bindings.EGSNUMS256E
		sOct = bindings.MakeOctet(sSize)
		defer sOct.Free()
	}

	WSize := 2*bindings.EFSNUMS256E + 1
	WOct := bindings.MakeOctet(WSize)
	defer WOct.Free()

	err = bindings.ECPKeyPairGenerateNUMS256E(R, sOct, WOct)

	sResult = sOct.ToBytes()

	W = WOct.ToBytes()

	return
}

// ECPPublicKeyValidateNUMS256E is a wrapper of bindings.ECP_NUMS256E_PUBLIC_KEY_VALIDATE.
func ECPPublicKeyValidateNUMS256E(W []byte) (err error) {

	WOct := bindings.NewOctet(W)
	defer WOct.Free()

	err = bindings.ECPPublicKeyValidateNUMS256E(WOct)

	return
}

// ECPSpDsaNUMS256E is a wrapper of bindings.ECP_NUMS256E_SP_DSA.
func ECPSpDsaNUMS256E(h int, R *bindings.Rand, key []byte, s []byte, M []byte) (c []byte, d []byte, err error) {

	keyOct := bindings.NewOctet(key)
	defer keyOct.Free()

	sOct := bindings.NewOctet(s)
	defer sOct.Free()

	MOct := bindings.NewOctet(M)
	defer MOct.Free()

	cSize := bindings.EGSNUMS256E
	cOct := bindings.MakeOctet(cSize)
	defer cOct.Free()

	dSize := bindings.EGSNUMS256E
	dOct := bindings.MakeOctet(dSize)
	defer dOct.Free()

	err = bindings.ECPSpDsaNUMS256E(h, R, keyOct, sOct, MOct, cOct, dOct)

	c = cOct.ToBytes()

	d = dOct.ToBytes()

	return
}

// ECPVpDsaNUMS256E is a wrapper of bindings.ECP_NUMS256E_VP_DSA.
func ECPVpDsaNUMS256E(h int, W []byte, M []byte, c []byte, d []byte) (err error) {

	WOct := bindings.NewOctet(W)
	defer WOct.Free()

	MOct := bindings.NewOctet(M)
	defer MOct.Free()

	cOct := bindings.NewOctet(c)
	defer cOct.Free()

	dOct := bindings.NewOctet(d)
	defer dOct.Free()

	err = bindings.ECPVpDsaNUMS256E(h, WOct, MOct, cOct, dOct)

	return
}
