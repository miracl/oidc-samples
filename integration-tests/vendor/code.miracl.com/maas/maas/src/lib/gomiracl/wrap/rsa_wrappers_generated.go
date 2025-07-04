// Generated by gen/wrappers/main.go from rsa_wrappers.go.tmpl.

package wrap

import "code.miracl.com/maas/maas/src/lib/gomiracl/bindings"

// PKCS15 is a wrapper of bindings.PKCS15.
func PKCS15(h int, m []byte, wSize int) (w []byte, err error) {

	mOct := bindings.NewOctet(m)
	defer mOct.Free()

	wOct := bindings.MakeOctet(wSize)
	defer wOct.Free()

	err = bindings.PKCS15(h, mOct, wOct)

	w = wOct.ToBytes()

	return
}

// OAEPencode is a wrapper of bindings.OAEP_ENCODE.
func OAEPencode(h int, m []byte, rng *bindings.Rand, p []byte, fSize int) (f []byte, err error) {

	mOct := bindings.NewOctet(m)
	defer mOct.Free()

	pOct := bindings.NewOctet(p)
	defer pOct.Free()

	fOct := bindings.MakeOctet(fSize)
	defer fOct.Free()

	err = bindings.OAEPencode(h, mOct, rng, pOct, fOct)

	f = fOct.ToBytes()

	return
}

// OAEPdecode is a wrapper of bindings.OAEP_DECODE.
func OAEPdecode(h int, p []byte, f []byte) (fResult []byte, err error) {

	pOct := bindings.NewOctet(p)
	defer pOct.Free()

	fOct := bindings.NewOctet(f)
	defer fOct.Free()

	err = bindings.OAEPdecode(h, pOct, fOct)

	fResult = fOct.ToBytes()

	return
}
