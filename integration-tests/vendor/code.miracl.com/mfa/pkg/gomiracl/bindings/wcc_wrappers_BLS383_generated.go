// Generated by gen/wrappers/main.go from wrap/wcc_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || BLS383

package bindings

// #cgo LDFLAGS: -lamcl_wcc_BLS383
// #include "amcl/wcc_BLS383.h"
import "C"

// Curve constants
const (
	WCCPGSBLS383 = int(C.WCC_PGS_BLS383)
	WCCPFSBLS383 = int(C.WCC_PFS_BLS383)
)

// WCCRandomGenerateBLS383 is a go wrapper for C.WCC_BLS383_RANDOM_GENERATE.
func WCCRandomGenerateBLS383(RNG *Rand, S *Octet) error {
	code := C.WCC_BLS383_RANDOM_GENERATE((*C.csprng)(RNG), (*C.octet)(S))

	return newError(code)
}

// WCCHqBLS383 is a go wrapper for C.WCC_BLS383_Hq.
func WCCHqBLS383(h int, A *Octet, B *Octet, C *Octet, D *Octet, res *Octet) {
	C.WCC_BLS383_Hq(C.int(h), (*C.octet)(A), (*C.octet)(B), (*C.octet)(C), (*C.octet)(D), (*C.octet)(res))
}

// WCCGetG1MultipleBLS383 is a go wrapper for C.WCC_BLS383_GET_G1_MULTIPLE.
func WCCGetG1MultipleBLS383(S *Octet, HID *Octet, VG1 *Octet) error {
	code := C.WCC_BLS383_GET_G1_MULTIPLE((*C.octet)(S), (*C.octet)(HID), (*C.octet)(VG1))

	return newError(code)
}

// WCCRecombineG1BLS383 is a go wrapper for C.WCC_BLS383_RECOMBINE_G1.
func WCCRecombineG1BLS383(R1 *Octet, R2 *Octet, R *Octet) error {
	code := C.WCC_BLS383_RECOMBINE_G1((*C.octet)(R1), (*C.octet)(R2), (*C.octet)(R))

	return newError(code)
}

// WCCGetG2MultipleBLS383 is a go wrapper for C.WCC_BLS383_GET_G2_MULTIPLE.
func WCCGetG2MultipleBLS383(S *Octet, HID *Octet, VG2 *Octet) error {
	code := C.WCC_BLS383_GET_G2_MULTIPLE((*C.octet)(S), (*C.octet)(HID), (*C.octet)(VG2))

	return newError(code)
}

// WCCRecombineG2BLS383 is a go wrapper for C.WCC_BLS383_RECOMBINE_G2.
func WCCRecombineG2BLS383(W1 *Octet, W2 *Octet, W *Octet) error {
	code := C.WCC_BLS383_RECOMBINE_G2((*C.octet)(W1), (*C.octet)(W2), (*C.octet)(W))

	return newError(code)
}

// WCCReceiverKeyBLS383 is a go wrapper for C.WCC_BLS383_RECEIVER_KEY.
func WCCReceiverKeyBLS383(h int, y *Octet, w *Octet, pia *Octet, pib *Octet, PaG1 *Octet, PgG1 *Octet, BKeyG2 *Octet, IDA *Octet, AESKey *Octet) error {
	code := C.WCC_BLS383_RECEIVER_KEY(C.int(h), (*C.octet)(y), (*C.octet)(w), (*C.octet)(pia), (*C.octet)(pib), (*C.octet)(PaG1), (*C.octet)(PgG1), (*C.octet)(BKeyG2), (*C.octet)(IDA), (*C.octet)(AESKey))

	return newError(code)
}

// WCCSenderKeyBLS383 is a go wrapper for C.WCC_BLS383_SENDER_KEY.
func WCCSenderKeyBLS383(h int, x *Octet, pia *Octet, pib *Octet, PbG2 *Octet, PgG1 *Octet, AKeyG1 *Octet, IDB *Octet, AESKey *Octet) error {
	code := C.WCC_BLS383_SENDER_KEY(C.int(h), (*C.octet)(x), (*C.octet)(pia), (*C.octet)(pib), (*C.octet)(PbG2), (*C.octet)(PgG1), (*C.octet)(AKeyG1), (*C.octet)(IDB), (*C.octet)(AESKey))

	return newError(code)
}
