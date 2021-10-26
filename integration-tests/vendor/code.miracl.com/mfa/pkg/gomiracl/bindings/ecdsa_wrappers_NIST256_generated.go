// Generated by gen/wrappers/main.go from wrap/ecdsa_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || NIST256

package bindings

// #cgo LDFLAGS: -lamcl_curve_NIST256
// #include <stdio.h>
// #include <stdlib.h>
// #include "amcl/amcl.h"
// #include "amcl/ecdh_NIST256.h"
import "C"

// Curve constants
const (
	EFSNIST256 = int(C.EFS_NIST256) // EFS is the ECC Field Size in bytes
	EGSNIST256 = int(C.EGS_NIST256) // EGS is the ECC Group Size in bytes
)

// ECPKeyPairGenerateNIST256 is a go wrapper for C.ECP_NIST256_KEY_PAIR_GENERATE.
func ECPKeyPairGenerateNIST256(R *Rand, s *Octet, W *Octet) error {
	code := C.ECP_NIST256_KEY_PAIR_GENERATE((*C.csprng)(R), (*C.octet)(s), (*C.octet)(W))

	return newError(code)
}

// ECPPublicKeyValidateNIST256 is a go wrapper for C.ECP_NIST256_PUBLIC_KEY_VALIDATE.
func ECPPublicKeyValidateNIST256(W *Octet) error {
	code := C.ECP_NIST256_PUBLIC_KEY_VALIDATE((*C.octet)(W))

	return newError(code)
}

// ECPSpDsaNIST256 is a go wrapper for C.ECP_NIST256_SP_DSA.
func ECPSpDsaNIST256(h int, R *Rand, key *Octet, s *Octet, M *Octet, c *Octet, d *Octet) error {
	code := C.ECP_NIST256_SP_DSA(C.int(h), (*C.csprng)(R), (*C.octet)(key), (*C.octet)(s), (*C.octet)(M), (*C.octet)(c), (*C.octet)(d))

	return newError(code)
}

// ECPVpDsaNIST256 is a go wrapper for C.ECP_NIST256_VP_DSA.
func ECPVpDsaNIST256(h int, W *Octet, M *Octet, c *Octet, d *Octet) error {
	code := C.ECP_NIST256_VP_DSA(C.int(h), (*C.octet)(W), (*C.octet)(M), (*C.octet)(c), (*C.octet)(d))

	return newError(code)
}