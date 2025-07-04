// Generated by gen/wrappers/main.go from ecdsa_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || ANSSI

package wrap

import "code.miracl.com/maas/maas/src/lib/gomiracl/bindings"

// ECPKeyPairGenerateANSSI is a wrapper of bindings.ECP_ANSSI_KEY_PAIR_GENERATE.
func ECPKeyPairGenerateANSSI(R *bindings.Rand, s []byte) (sResult []byte, W []byte, err error) {

	var sOct *bindings.Octet
	if s != nil {
		sOct = bindings.NewOctet(s)
		defer sOct.Free()
	} else {
		sSize := bindings.EGSANSSI
		sOct = bindings.MakeOctet(sSize)
		defer sOct.Free()
	}

	WSize := 2*bindings.EFSANSSI + 1
	WOct := bindings.MakeOctet(WSize)
	defer WOct.Free()

	err = bindings.ECPKeyPairGenerateANSSI(R, sOct, WOct)

	sResult = sOct.ToBytes()

	W = WOct.ToBytes()

	return
}

// ECPPublicKeyValidateANSSI is a wrapper of bindings.ECP_ANSSI_PUBLIC_KEY_VALIDATE.
func ECPPublicKeyValidateANSSI(W []byte) (err error) {

	WOct := bindings.NewOctet(W)
	defer WOct.Free()

	err = bindings.ECPPublicKeyValidateANSSI(WOct)

	return
}

// ECPSpDsaANSSI is a wrapper of bindings.ECP_ANSSI_SP_DSA.
func ECPSpDsaANSSI(h int, R *bindings.Rand, key []byte, s []byte, M []byte) (c []byte, d []byte, err error) {

	keyOct := bindings.NewOctet(key)
	defer keyOct.Free()

	sOct := bindings.NewOctet(s)
	defer sOct.Free()

	MOct := bindings.NewOctet(M)
	defer MOct.Free()

	cSize := bindings.EGSANSSI
	cOct := bindings.MakeOctet(cSize)
	defer cOct.Free()

	dSize := bindings.EGSANSSI
	dOct := bindings.MakeOctet(dSize)
	defer dOct.Free()

	err = bindings.ECPSpDsaANSSI(h, R, keyOct, sOct, MOct, cOct, dOct)

	c = cOct.ToBytes()

	d = dOct.ToBytes()

	return
}

// ECPVpDsaANSSI is a wrapper of bindings.ECP_ANSSI_VP_DSA.
func ECPVpDsaANSSI(h int, W []byte, M []byte, c []byte, d []byte) (err error) {

	WOct := bindings.NewOctet(W)
	defer WOct.Free()

	MOct := bindings.NewOctet(M)
	defer MOct.Free()

	cOct := bindings.NewOctet(c)
	defer cOct.Free()

	dOct := bindings.NewOctet(d)
	defer dOct.Free()

	err = bindings.ECPVpDsaANSSI(h, WOct, MOct, cOct, dOct)

	return
}
