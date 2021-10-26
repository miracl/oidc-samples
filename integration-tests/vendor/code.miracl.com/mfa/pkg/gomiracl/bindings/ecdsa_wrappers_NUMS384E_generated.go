// Generated by gen/wrappers/main.go from wrap/ecdsa_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || NUMS384E

package bindings

// #cgo LDFLAGS: -lamcl_curve_NUMS384E
// #include <stdio.h>
// #include <stdlib.h>
// #include "amcl/amcl.h"
// #include "amcl/ecdh_NUMS384E.h"
import "C"

// Curve constants
const (
	EFSNUMS384E = int(C.EFS_NUMS384E) // EFS is the ECC Field Size in bytes
	EGSNUMS384E = int(C.EGS_NUMS384E) // EGS is the ECC Group Size in bytes
)

// ECPKeyPairGenerateNUMS384E is a go wrapper for C.ECP_NUMS384E_KEY_PAIR_GENERATE.
func ECPKeyPairGenerateNUMS384E(R *Rand, s *Octet, W *Octet) error {
	code := C.ECP_NUMS384E_KEY_PAIR_GENERATE((*C.csprng)(R), (*C.octet)(s), (*C.octet)(W))

	return newError(code)
}

// ECPPublicKeyValidateNUMS384E is a go wrapper for C.ECP_NUMS384E_PUBLIC_KEY_VALIDATE.
func ECPPublicKeyValidateNUMS384E(W *Octet) error {
	code := C.ECP_NUMS384E_PUBLIC_KEY_VALIDATE((*C.octet)(W))

	return newError(code)
}

// ECPSpDsaNUMS384E is a go wrapper for C.ECP_NUMS384E_SP_DSA.
func ECPSpDsaNUMS384E(h int, R *Rand, key *Octet, s *Octet, M *Octet, c *Octet, d *Octet) error {
	code := C.ECP_NUMS384E_SP_DSA(C.int(h), (*C.csprng)(R), (*C.octet)(key), (*C.octet)(s), (*C.octet)(M), (*C.octet)(c), (*C.octet)(d))

	return newError(code)
}

// ECPVpDsaNUMS384E is a go wrapper for C.ECP_NUMS384E_VP_DSA.
func ECPVpDsaNUMS384E(h int, W *Octet, M *Octet, c *Octet, d *Octet) error {
	code := C.ECP_NUMS384E_VP_DSA(C.int(h), (*C.octet)(W), (*C.octet)(M), (*C.octet)(c), (*C.octet)(d))

	return newError(code)
}