// Generated by gen/wrappers/main.go from wrap/ecdsa_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || NUMS384W

package bindings

// #cgo LDFLAGS: -lamcl_curve_NUMS384W
// #include <stdio.h>
// #include <stdlib.h>
// #include "amcl/amcl.h"
// #include "amcl/ecdh_NUMS384W.h"
import "C"

// Curve constants
const (
	EFSNUMS384W = int(C.EFS_NUMS384W) // EFS is the ECC Field Size in bytes
	EGSNUMS384W = int(C.EGS_NUMS384W) // EGS is the ECC Group Size in bytes
)

// ECPKeyPairGenerateNUMS384W is a go wrapper for C.ECP_NUMS384W_KEY_PAIR_GENERATE.
func ECPKeyPairGenerateNUMS384W(R *Rand, s *Octet, W *Octet) error {
	var sCOct *C.octet = nil
	if s != nil {
		sCOct = (*C.octet)(&s.cOctet)
	}

	var WCOct *C.octet = nil
	if W != nil {
		WCOct = (*C.octet)(&W.cOctet)
	}

	code := C.ECP_NUMS384W_KEY_PAIR_GENERATE((*C.csprng)(R), sCOct, WCOct)

	return newError(code)
}

// ECPPublicKeyValidateNUMS384W is a go wrapper for C.ECP_NUMS384W_PUBLIC_KEY_VALIDATE.
func ECPPublicKeyValidateNUMS384W(W *Octet) error {
	var WCOct *C.octet = nil
	if W != nil {
		WCOct = (*C.octet)(&W.cOctet)
	}

	code := C.ECP_NUMS384W_PUBLIC_KEY_VALIDATE(WCOct)

	return newError(code)
}

// ECPSpDsaNUMS384W is a go wrapper for C.ECP_NUMS384W_SP_DSA.
func ECPSpDsaNUMS384W(h int, R *Rand, key *Octet, s *Octet, M *Octet, c *Octet, d *Octet) error {

	var keyCOct *C.octet = nil
	if key != nil {
		keyCOct = (*C.octet)(&key.cOctet)
	}

	var sCOct *C.octet = nil
	if s != nil {
		sCOct = (*C.octet)(&s.cOctet)
	}

	var MCOct *C.octet = nil
	if M != nil {
		MCOct = (*C.octet)(&M.cOctet)
	}

	var cCOct *C.octet = nil
	if c != nil {
		cCOct = (*C.octet)(&c.cOctet)
	}

	var dCOct *C.octet = nil
	if d != nil {
		dCOct = (*C.octet)(&d.cOctet)
	}

	code := C.ECP_NUMS384W_SP_DSA(C.int(h), (*C.csprng)(R), keyCOct, sCOct, MCOct, cCOct, dCOct)

	return newError(code)
}

// ECPVpDsaNUMS384W is a go wrapper for C.ECP_NUMS384W_VP_DSA.
func ECPVpDsaNUMS384W(h int, W *Octet, M *Octet, c *Octet, d *Octet) error {
	var WCOct *C.octet = nil
	if W != nil {
		WCOct = (*C.octet)(&W.cOctet)
	}

	var MCOct *C.octet = nil
	if M != nil {
		MCOct = (*C.octet)(&M.cOctet)
	}

	var cCOct *C.octet = nil
	if c != nil {
		cCOct = (*C.octet)(&c.cOctet)
	}

	var dCOct *C.octet = nil
	if d != nil {
		dCOct = (*C.octet)(&d.cOctet)
	}

	code := C.ECP_NUMS384W_VP_DSA(C.int(h), WCOct, MCOct, cCOct, dCOct)

	return newError(code)
}
