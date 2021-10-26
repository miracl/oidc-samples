// Generated by gen/wrappers/main.go from wrap/mpin_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || BLS461

package bindings

// #cgo LDFLAGS: -lamcl_curve_BLS461 -lamcl_mpin_BLS461 -lamcl_pairing_BLS461
// #include "amcl/amcl.h"
// #include "amcl/mpin_BLS461.h"
// #include "amcl/randapi.h"
// #include "amcl/utils.h"
import "C"

// Curve constants
const (
	PASBLS461        = int(C.MPIN_PAS)
	PGSBLS461        = int(C.PGS_BLS461)
	PFSBLS461        = int(C.PFS_BLS461)
	G1SBLS461        = 2*PFSBLS461 + 1
	G2SBLS461        = 4 * PFSBLS461
	GTSBLS461        = 12 * PFSBLS461
	AESKeySizeBLS461 = int(C.AESKEY_BLS461)
)

// Client1BLS461 is a go wrapper for C.MPIN_BLS461_CLIENT_1.
func Client1BLS461(h int, d int, ID *Octet, R *Rand, x *Octet, pin int, T *Octet, S *Octet, U *Octet, UT *Octet, TP *Octet) error {
	code := C.MPIN_BLS461_CLIENT_1(C.int(h), C.int(d), (*C.octet)(ID), (*C.csprng)(R), (*C.octet)(x), C.int(pin), (*C.octet)(T), (*C.octet)(S), (*C.octet)(U), (*C.octet)(UT), (*C.octet)(TP))

	return newError(code)
}

// Client2BLS461 is a go wrapper for C.MPIN_BLS461_CLIENT_2.
func Client2BLS461(x *Octet, y *Octet, V *Octet) error {
	code := C.MPIN_BLS461_CLIENT_2((*C.octet)(x), (*C.octet)(y), (*C.octet)(V))

	return newError(code)
}

// ClientKeyBLS461 is a go wrapper for C.MPIN_BLS461_CLIENT_KEY.
func ClientKeyBLS461(h int, g1 *Octet, g2 *Octet, pin int, r *Octet, x *Octet, p *Octet, T *Octet, K *Octet) error {
	code := C.MPIN_BLS461_CLIENT_KEY(C.int(h), (*C.octet)(g1), (*C.octet)(g2), C.int(pin), (*C.octet)(r), (*C.octet)(x), (*C.octet)(p), (*C.octet)(T), (*C.octet)(K))

	return newError(code)
}

// ClientBLS461 is a go wrapper for C.MPIN_BLS461_CLIENT.
func ClientBLS461(h int, d int, ID *Octet, R *Rand, x *Octet, pin int, T *Octet, V *Octet, U *Octet, UT *Octet, TP *Octet, MESSAGE *Octet, t int, y *Octet) error {
	code := C.MPIN_BLS461_CLIENT(C.int(h), C.int(d), (*C.octet)(ID), (*C.csprng)(R), (*C.octet)(x), C.int(pin), (*C.octet)(T), (*C.octet)(V), (*C.octet)(U), (*C.octet)(UT), (*C.octet)(TP), (*C.octet)(MESSAGE), C.int(t), (*C.octet)(y))

	return newError(code)
}

// ExtractPINBLS461 is a go wrapper for C.MPIN_BLS461_EXTRACT_PIN.
func ExtractPINBLS461(h int, ID *Octet, pin int, CS *Octet) error {
	code := C.MPIN_BLS461_EXTRACT_PIN(C.int(h), (*C.octet)(ID), C.int(pin), (*C.octet)(CS))

	return newError(code)
}

// GetClientPermitBLS461 is a go wrapper for C.MPIN_BLS461_GET_CLIENT_PERMIT.
func GetClientPermitBLS461(h int, d int, S *Octet, ID *Octet, TP *Octet) error {
	code := C.MPIN_BLS461_GET_CLIENT_PERMIT(C.int(h), C.int(d), (*C.octet)(S), (*C.octet)(ID), (*C.octet)(TP))

	return newError(code)
}

// GetClientSecretBLS461 is a go wrapper for C.MPIN_BLS461_GET_CLIENT_SECRET.
func GetClientSecretBLS461(S *Octet, ID *Octet, CS *Octet) error {
	code := C.MPIN_BLS461_GET_CLIENT_SECRET((*C.octet)(S), (*C.octet)(ID), (*C.octet)(CS))

	return newError(code)
}

// GetDVSKeyPairBLS461 is a go wrapper for C.MPIN_BLS461_GET_DVS_KEYPAIR.
func GetDVSKeyPairBLS461(R *Rand, Z *Octet, Pa *Octet) error {
	code := C.MPIN_BLS461_GET_DVS_KEYPAIR((*C.csprng)(R), (*C.octet)(Z), (*C.octet)(Pa))

	return newError(code)
}

// GetG1MultipleBLS461 is a go wrapper for C.MPIN_BLS461_GET_G1_MULTIPLE.
func GetG1MultipleBLS461(R *Rand, t int, x *Octet, G *Octet, W *Octet) error {
	code := C.MPIN_BLS461_GET_G1_MULTIPLE((*C.csprng)(R), C.int(t), (*C.octet)(x), (*C.octet)(G), (*C.octet)(W))

	return newError(code)
}

// GetServerSecretBLS461 is a go wrapper for C.MPIN_BLS461_GET_SERVER_SECRET.
func GetServerSecretBLS461(S *Octet, SS *Octet) error {
	code := C.MPIN_BLS461_GET_SERVER_SECRET((*C.octet)(S), (*C.octet)(SS))

	return newError(code)
}

// KangarooBLS461 is a go wrapper for C.MPIN_BLS461_KANGAROO.
func KangarooBLS461(E *Octet, F *Octet) error {
	code := C.MPIN_BLS461_KANGAROO((*C.octet)(E), (*C.octet)(F))

	return newError(code)
}

// PrecomputeBLS461 is a go wrapper for C.MPIN_BLS461_PRECOMPUTE.
func PrecomputeBLS461(T *Octet, ID *Octet, CP *Octet, g1 *Octet, g2 *Octet) error {
	code := C.MPIN_BLS461_PRECOMPUTE((*C.octet)(T), (*C.octet)(ID), (*C.octet)(CP), (*C.octet)(g1), (*C.octet)(g2))

	return newError(code)
}

// RandomGenerateBLS461 is a go wrapper for C.MPIN_BLS461_RANDOM_GENERATE.
func RandomGenerateBLS461(R *Rand, S *Octet) error {
	code := C.MPIN_BLS461_RANDOM_GENERATE((*C.csprng)(R), (*C.octet)(S))

	return newError(code)
}

// RecombineG1BLS461 is a go wrapper for C.MPIN_BLS461_RECOMBINE_G1.
func RecombineG1BLS461(Q1 *Octet, Q2 *Octet, Q *Octet) error {
	code := C.MPIN_BLS461_RECOMBINE_G1((*C.octet)(Q1), (*C.octet)(Q2), (*C.octet)(Q))

	return newError(code)
}

// RecombineG2BLS461 is a go wrapper for C.MPIN_BLS461_RECOMBINE_G2.
func RecombineG2BLS461(P1 *Octet, P2 *Octet, P *Octet) error {
	code := C.MPIN_BLS461_RECOMBINE_G2((*C.octet)(P1), (*C.octet)(P2), (*C.octet)(P))

	return newError(code)
}

// Server2BLS461 is a go wrapper for C.MPIN_BLS461_SERVER_2.
func Server2BLS461(d int, HID *Octet, HTID *Octet, y *Octet, SS *Octet, U *Octet, UT *Octet, V *Octet, E *Octet, F *Octet, Pa *Octet) error {
	code := C.MPIN_BLS461_SERVER_2(C.int(d), (*C.octet)(HID), (*C.octet)(HTID), (*C.octet)(y), (*C.octet)(SS), (*C.octet)(U), (*C.octet)(UT), (*C.octet)(V), (*C.octet)(E), (*C.octet)(F), (*C.octet)(Pa))

	return newError(code)
}

// ServerKeyBLS461 is a go wrapper for C.MPIN_BLS461_SERVER_KEY.
func ServerKeyBLS461(h int, Z *Octet, SS *Octet, w *Octet, p *Octet, I *Octet, U *Octet, UT *Octet, K *Octet) error {
	code := C.MPIN_BLS461_SERVER_KEY(C.int(h), (*C.octet)(Z), (*C.octet)(SS), (*C.octet)(w), (*C.octet)(p), (*C.octet)(I), (*C.octet)(U), (*C.octet)(UT), (*C.octet)(K))

	return newError(code)
}

// ServerBLS461 is a go wrapper for C.MPIN_BLS461_SERVER.
func ServerBLS461(h int, d int, HID *Octet, HTID *Octet, y *Octet, SS *Octet, U *Octet, UT *Octet, V *Octet, E *Octet, F *Octet, ID *Octet, MESSAGE *Octet, t int, Pa *Octet) error {
	code := C.MPIN_BLS461_SERVER(C.int(h), C.int(d), (*C.octet)(HID), (*C.octet)(HTID), (*C.octet)(y), (*C.octet)(SS), (*C.octet)(U), (*C.octet)(UT), (*C.octet)(V), (*C.octet)(E), (*C.octet)(F), (*C.octet)(ID), (*C.octet)(MESSAGE), C.int(t), (*C.octet)(Pa))

	return newError(code)
}

// Server1BLS461 is a go wrapper for C.MPIN_BLS461_SERVER_1.
func Server1BLS461(h int, d int, ID *Octet, HID *Octet, HTID *Octet) {
	C.MPIN_BLS461_SERVER_1(C.int(h), C.int(d), (*C.octet)(ID), (*C.octet)(HID), (*C.octet)(HTID))
}
