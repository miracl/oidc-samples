// Generated by gen/wrappers/main.go from wrap/ecdsa_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || ED25519

package bindings

// #cgo LDFLAGS: -lamcl_curve_ED25519
// #include <stdio.h>
// #include <stdlib.h>
// #include "amcl/amcl.h"
// #include "amcl/ecdh_ED25519.h"
import "C"

// Curve constants
const (
	EFSED25519 = int(C.EFS_ED25519) // EFS is the ECC Field Size in bytes
	EGSED25519 = int(C.EGS_ED25519) // EGS is the ECC Group Size in bytes
)

// ECPKeyPairGenerateED25519 is a go wrapper for C.ECP_ED25519_KEY_PAIR_GENERATE.
func ECPKeyPairGenerateED25519(R *Rand, s *Octet, W *Octet) error {
	code := C.ECP_ED25519_KEY_PAIR_GENERATE((*C.csprng)(R), (*C.octet)(s), (*C.octet)(W))

	return newError(code)
}

// ECPPublicKeyValidateED25519 is a go wrapper for C.ECP_ED25519_PUBLIC_KEY_VALIDATE.
func ECPPublicKeyValidateED25519(W *Octet) error {
	code := C.ECP_ED25519_PUBLIC_KEY_VALIDATE((*C.octet)(W))

	return newError(code)
}

// ECPSpDsaED25519 is a go wrapper for C.ECP_ED25519_SP_DSA.
func ECPSpDsaED25519(h int, R *Rand, key *Octet, s *Octet, M *Octet, c *Octet, d *Octet) error {
	code := C.ECP_ED25519_SP_DSA(C.int(h), (*C.csprng)(R), (*C.octet)(key), (*C.octet)(s), (*C.octet)(M), (*C.octet)(c), (*C.octet)(d))

	return newError(code)
}

// ECPVpDsaED25519 is a go wrapper for C.ECP_ED25519_VP_DSA.
func ECPVpDsaED25519(h int, W *Octet, M *Octet, c *Octet, d *Octet) error {
	code := C.ECP_ED25519_VP_DSA(C.int(h), (*C.octet)(W), (*C.octet)(M), (*C.octet)(c), (*C.octet)(d))

	return newError(code)
}