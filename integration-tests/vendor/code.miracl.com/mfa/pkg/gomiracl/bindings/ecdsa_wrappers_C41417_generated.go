// Generated by gen/wrappers/main.go from wrap/ecdsa_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || C41417

package bindings

// #cgo LDFLAGS: -lamcl_curve_C41417
// #include <stdio.h>
// #include <stdlib.h>
// #include "amcl/amcl.h"
// #include "amcl/ecdh_C41417.h"
import "C"

// Curve constants
const (
	EFSC41417 = int(C.EFS_C41417) // EFS is the ECC Field Size in bytes
	EGSC41417 = int(C.EGS_C41417) // EGS is the ECC Group Size in bytes
)

// ECPKeyPairGenerateC41417 is a go wrapper for C.ECP_C41417_KEY_PAIR_GENERATE.
func ECPKeyPairGenerateC41417(R *Rand, s *Octet, W *Octet) error {
	code := C.ECP_C41417_KEY_PAIR_GENERATE((*C.csprng)(R), (*C.octet)(s), (*C.octet)(W))

	return newError(code)
}

// ECPPublicKeyValidateC41417 is a go wrapper for C.ECP_C41417_PUBLIC_KEY_VALIDATE.
func ECPPublicKeyValidateC41417(W *Octet) error {
	code := C.ECP_C41417_PUBLIC_KEY_VALIDATE((*C.octet)(W))

	return newError(code)
}

// ECPSpDsaC41417 is a go wrapper for C.ECP_C41417_SP_DSA.
func ECPSpDsaC41417(h int, R *Rand, key *Octet, s *Octet, M *Octet, c *Octet, d *Octet) error {
	code := C.ECP_C41417_SP_DSA(C.int(h), (*C.csprng)(R), (*C.octet)(key), (*C.octet)(s), (*C.octet)(M), (*C.octet)(c), (*C.octet)(d))

	return newError(code)
}

// ECPVpDsaC41417 is a go wrapper for C.ECP_C41417_VP_DSA.
func ECPVpDsaC41417(h int, W *Octet, M *Octet, c *Octet, d *Octet) error {
	code := C.ECP_C41417_VP_DSA(C.int(h), (*C.octet)(W), (*C.octet)(M), (*C.octet)(c), (*C.octet)(d))

	return newError(code)
}
