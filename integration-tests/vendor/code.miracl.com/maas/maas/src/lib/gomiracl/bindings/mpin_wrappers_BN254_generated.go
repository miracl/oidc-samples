// Generated by gen/wrappers/main.go from wrap/mpin_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || BN254

package bindings

// #cgo LDFLAGS: -lamcl_curve_BN254 -lamcl_mpin_BN254 -lamcl_pairing_BN254
// #include "amcl/amcl.h"
// #include "amcl/mpin_BN254.h"
// #include "amcl/randapi.h"
// #include "amcl/utils.h"
import "C"

// Curve constants
const (
	PASBN254        = int(C.MPIN_PAS)
	PGSBN254        = int(C.PGS_BN254)
	PFSBN254        = int(C.PFS_BN254)
	G1SBN254        = 2*PFSBN254 + 1
	G2SBN254        = 4 * PFSBN254
	GTSBN254        = 12 * PFSBN254
	AESKeySizeBN254 = int(C.AESKEY_BN254)
)

// Client1BN254 is a go wrapper for C.MPIN_BN254_CLIENT_1.
func Client1BN254(h int, d int, ID *Octet, R *Rand, x *Octet, pin int, T *Octet, S *Octet, U *Octet, UT *Octet, TP *Octet) error {

	var IDCOct *C.octet = nil
	if ID != nil {
		IDCOct = (*C.octet)(&ID.cOctet)
	}

	var xCOct *C.octet = nil
	if x != nil {
		xCOct = (*C.octet)(&x.cOctet)
	}

	var TCOct *C.octet = nil
	if T != nil {
		TCOct = (*C.octet)(&T.cOctet)
	}

	var SCOct *C.octet = nil
	if S != nil {
		SCOct = (*C.octet)(&S.cOctet)
	}

	var UCOct *C.octet = nil
	if U != nil {
		UCOct = (*C.octet)(&U.cOctet)
	}

	var UTCOct *C.octet = nil
	if UT != nil {
		UTCOct = (*C.octet)(&UT.cOctet)
	}

	var TPCOct *C.octet = nil
	if TP != nil {
		TPCOct = (*C.octet)(&TP.cOctet)
	}

	code := C.MPIN_BN254_CLIENT_1(C.int(h), C.int(d), IDCOct, (*C.csprng)(R), xCOct, C.int(pin), TCOct, SCOct, UCOct, UTCOct, TPCOct)

	return newError(code)
}

// Client2BN254 is a go wrapper for C.MPIN_BN254_CLIENT_2.
func Client2BN254(x *Octet, y *Octet, V *Octet) error {
	var xCOct *C.octet = nil
	if x != nil {
		xCOct = (*C.octet)(&x.cOctet)
	}

	var yCOct *C.octet = nil
	if y != nil {
		yCOct = (*C.octet)(&y.cOctet)
	}

	var VCOct *C.octet = nil
	if V != nil {
		VCOct = (*C.octet)(&V.cOctet)
	}

	code := C.MPIN_BN254_CLIENT_2(xCOct, yCOct, VCOct)

	return newError(code)
}

// ClientKeyBN254 is a go wrapper for C.MPIN_BN254_CLIENT_KEY.
func ClientKeyBN254(h int, g1 *Octet, g2 *Octet, pin int, r *Octet, x *Octet, p *Octet, T *Octet, K *Octet) error {
	var g1COct *C.octet = nil
	if g1 != nil {
		g1COct = (*C.octet)(&g1.cOctet)
	}

	var g2COct *C.octet = nil
	if g2 != nil {
		g2COct = (*C.octet)(&g2.cOctet)
	}

	var rCOct *C.octet = nil
	if r != nil {
		rCOct = (*C.octet)(&r.cOctet)
	}

	var xCOct *C.octet = nil
	if x != nil {
		xCOct = (*C.octet)(&x.cOctet)
	}

	var pCOct *C.octet = nil
	if p != nil {
		pCOct = (*C.octet)(&p.cOctet)
	}

	var TCOct *C.octet = nil
	if T != nil {
		TCOct = (*C.octet)(&T.cOctet)
	}

	var KCOct *C.octet = nil
	if K != nil {
		KCOct = (*C.octet)(&K.cOctet)
	}

	code := C.MPIN_BN254_CLIENT_KEY(C.int(h), g1COct, g2COct, C.int(pin), rCOct, xCOct, pCOct, TCOct, KCOct)

	return newError(code)
}

// ClientBN254 is a go wrapper for C.MPIN_BN254_CLIENT.
func ClientBN254(h int, d int, ID *Octet, R *Rand, x *Octet, pin int, T *Octet, V *Octet, U *Octet, UT *Octet, TP *Octet, MESSAGE *Octet, t int, y *Octet) error {

	var IDCOct *C.octet = nil
	if ID != nil {
		IDCOct = (*C.octet)(&ID.cOctet)
	}

	var xCOct *C.octet = nil
	if x != nil {
		xCOct = (*C.octet)(&x.cOctet)
	}

	var TCOct *C.octet = nil
	if T != nil {
		TCOct = (*C.octet)(&T.cOctet)
	}

	var VCOct *C.octet = nil
	if V != nil {
		VCOct = (*C.octet)(&V.cOctet)
	}

	var UCOct *C.octet = nil
	if U != nil {
		UCOct = (*C.octet)(&U.cOctet)
	}

	var UTCOct *C.octet = nil
	if UT != nil {
		UTCOct = (*C.octet)(&UT.cOctet)
	}

	var TPCOct *C.octet = nil
	if TP != nil {
		TPCOct = (*C.octet)(&TP.cOctet)
	}

	var MESSAGECOct *C.octet = nil
	if MESSAGE != nil {
		MESSAGECOct = (*C.octet)(&MESSAGE.cOctet)
	}

	var yCOct *C.octet = nil
	if y != nil {
		yCOct = (*C.octet)(&y.cOctet)
	}

	code := C.MPIN_BN254_CLIENT(C.int(h), C.int(d), IDCOct, (*C.csprng)(R), xCOct, C.int(pin), TCOct, VCOct, UCOct, UTCOct, TPCOct, MESSAGECOct, C.int(t), yCOct)

	return newError(code)
}

// ExtractPINBN254 is a go wrapper for C.MPIN_BN254_EXTRACT_PIN.
func ExtractPINBN254(h int, ID *Octet, pin int, CS *Octet) error {
	var IDCOct *C.octet = nil
	if ID != nil {
		IDCOct = (*C.octet)(&ID.cOctet)
	}

	var CSCOct *C.octet = nil
	if CS != nil {
		CSCOct = (*C.octet)(&CS.cOctet)
	}

	code := C.MPIN_BN254_EXTRACT_PIN(C.int(h), IDCOct, C.int(pin), CSCOct)

	return newError(code)
}

// GetClientPermitBN254 is a go wrapper for C.MPIN_BN254_GET_CLIENT_PERMIT.
func GetClientPermitBN254(h int, d int, S *Octet, ID *Octet, TP *Octet) error {

	var SCOct *C.octet = nil
	if S != nil {
		SCOct = (*C.octet)(&S.cOctet)
	}

	var IDCOct *C.octet = nil
	if ID != nil {
		IDCOct = (*C.octet)(&ID.cOctet)
	}

	var TPCOct *C.octet = nil
	if TP != nil {
		TPCOct = (*C.octet)(&TP.cOctet)
	}

	code := C.MPIN_BN254_GET_CLIENT_PERMIT(C.int(h), C.int(d), SCOct, IDCOct, TPCOct)

	return newError(code)
}

// GetClientSecretBN254 is a go wrapper for C.MPIN_BN254_GET_CLIENT_SECRET.
func GetClientSecretBN254(S *Octet, ID *Octet, CS *Octet) error {
	var SCOct *C.octet = nil
	if S != nil {
		SCOct = (*C.octet)(&S.cOctet)
	}

	var IDCOct *C.octet = nil
	if ID != nil {
		IDCOct = (*C.octet)(&ID.cOctet)
	}

	var CSCOct *C.octet = nil
	if CS != nil {
		CSCOct = (*C.octet)(&CS.cOctet)
	}

	code := C.MPIN_BN254_GET_CLIENT_SECRET(SCOct, IDCOct, CSCOct)

	return newError(code)
}

// GetDVSKeyPairBN254 is a go wrapper for C.MPIN_BN254_GET_DVS_KEYPAIR.
func GetDVSKeyPairBN254(R *Rand, Z *Octet, Pa *Octet) error {
	var ZCOct *C.octet = nil
	if Z != nil {
		ZCOct = (*C.octet)(&Z.cOctet)
	}

	var PaCOct *C.octet = nil
	if Pa != nil {
		PaCOct = (*C.octet)(&Pa.cOctet)
	}

	code := C.MPIN_BN254_GET_DVS_KEYPAIR((*C.csprng)(R), ZCOct, PaCOct)

	return newError(code)
}

// GetG1MultipleBN254 is a go wrapper for C.MPIN_BN254_GET_G1_MULTIPLE.
func GetG1MultipleBN254(R *Rand, t int, x *Octet, G *Octet, W *Octet) error {

	var xCOct *C.octet = nil
	if x != nil {
		xCOct = (*C.octet)(&x.cOctet)
	}

	var GCOct *C.octet = nil
	if G != nil {
		GCOct = (*C.octet)(&G.cOctet)
	}

	var WCOct *C.octet = nil
	if W != nil {
		WCOct = (*C.octet)(&W.cOctet)
	}

	code := C.MPIN_BN254_GET_G1_MULTIPLE((*C.csprng)(R), C.int(t), xCOct, GCOct, WCOct)

	return newError(code)
}

// GetServerSecretBN254 is a go wrapper for C.MPIN_BN254_GET_SERVER_SECRET.
func GetServerSecretBN254(S *Octet, SS *Octet) error {
	var SCOct *C.octet = nil
	if S != nil {
		SCOct = (*C.octet)(&S.cOctet)
	}

	var SSCOct *C.octet = nil
	if SS != nil {
		SSCOct = (*C.octet)(&SS.cOctet)
	}

	code := C.MPIN_BN254_GET_SERVER_SECRET(SCOct, SSCOct)

	return newError(code)
}

// KangarooBN254 is a go wrapper for C.MPIN_BN254_KANGAROO.
func KangarooBN254(E *Octet, F *Octet) error {
	var ECOct *C.octet = nil
	if E != nil {
		ECOct = (*C.octet)(&E.cOctet)
	}

	var FCOct *C.octet = nil
	if F != nil {
		FCOct = (*C.octet)(&F.cOctet)
	}

	code := C.MPIN_BN254_KANGAROO(ECOct, FCOct)

	return newError(code)
}

// PrecomputeBN254 is a go wrapper for C.MPIN_BN254_PRECOMPUTE.
func PrecomputeBN254(T *Octet, ID *Octet, CP *Octet, g1 *Octet, g2 *Octet) error {
	var TCOct *C.octet = nil
	if T != nil {
		TCOct = (*C.octet)(&T.cOctet)
	}

	var IDCOct *C.octet = nil
	if ID != nil {
		IDCOct = (*C.octet)(&ID.cOctet)
	}

	var CPCOct *C.octet = nil
	if CP != nil {
		CPCOct = (*C.octet)(&CP.cOctet)
	}

	var g1COct *C.octet = nil
	if g1 != nil {
		g1COct = (*C.octet)(&g1.cOctet)
	}

	var g2COct *C.octet = nil
	if g2 != nil {
		g2COct = (*C.octet)(&g2.cOctet)
	}

	code := C.MPIN_BN254_PRECOMPUTE(TCOct, IDCOct, CPCOct, g1COct, g2COct)

	return newError(code)
}

// RandomGenerateBN254 is a go wrapper for C.MPIN_BN254_RANDOM_GENERATE.
func RandomGenerateBN254(R *Rand, S *Octet) error {
	var SCOct *C.octet = nil
	if S != nil {
		SCOct = (*C.octet)(&S.cOctet)
	}

	code := C.MPIN_BN254_RANDOM_GENERATE((*C.csprng)(R), SCOct)

	return newError(code)
}

// RecombineG1BN254 is a go wrapper for C.MPIN_BN254_RECOMBINE_G1.
func RecombineG1BN254(Q1 *Octet, Q2 *Octet, Q *Octet) error {
	var Q1COct *C.octet = nil
	if Q1 != nil {
		Q1COct = (*C.octet)(&Q1.cOctet)
	}

	var Q2COct *C.octet = nil
	if Q2 != nil {
		Q2COct = (*C.octet)(&Q2.cOctet)
	}

	var QCOct *C.octet = nil
	if Q != nil {
		QCOct = (*C.octet)(&Q.cOctet)
	}

	code := C.MPIN_BN254_RECOMBINE_G1(Q1COct, Q2COct, QCOct)

	return newError(code)
}

// RecombineG2BN254 is a go wrapper for C.MPIN_BN254_RECOMBINE_G2.
func RecombineG2BN254(P1 *Octet, P2 *Octet, P *Octet) error {
	var P1COct *C.octet = nil
	if P1 != nil {
		P1COct = (*C.octet)(&P1.cOctet)
	}

	var P2COct *C.octet = nil
	if P2 != nil {
		P2COct = (*C.octet)(&P2.cOctet)
	}

	var PCOct *C.octet = nil
	if P != nil {
		PCOct = (*C.octet)(&P.cOctet)
	}

	code := C.MPIN_BN254_RECOMBINE_G2(P1COct, P2COct, PCOct)

	return newError(code)
}

// Server2BN254 is a go wrapper for C.MPIN_BN254_SERVER_2.
func Server2BN254(d int, HID *Octet, HTID *Octet, y *Octet, SS *Octet, U *Octet, UT *Octet, V *Octet, E *Octet, F *Octet, Pa *Octet) error {
	var HIDCOct *C.octet = nil
	if HID != nil {
		HIDCOct = (*C.octet)(&HID.cOctet)
	}

	var HTIDCOct *C.octet = nil
	if HTID != nil {
		HTIDCOct = (*C.octet)(&HTID.cOctet)
	}

	var yCOct *C.octet = nil
	if y != nil {
		yCOct = (*C.octet)(&y.cOctet)
	}

	var SSCOct *C.octet = nil
	if SS != nil {
		SSCOct = (*C.octet)(&SS.cOctet)
	}

	var UCOct *C.octet = nil
	if U != nil {
		UCOct = (*C.octet)(&U.cOctet)
	}

	var UTCOct *C.octet = nil
	if UT != nil {
		UTCOct = (*C.octet)(&UT.cOctet)
	}

	var VCOct *C.octet = nil
	if V != nil {
		VCOct = (*C.octet)(&V.cOctet)
	}

	var ECOct *C.octet = nil
	if E != nil {
		ECOct = (*C.octet)(&E.cOctet)
	}

	var FCOct *C.octet = nil
	if F != nil {
		FCOct = (*C.octet)(&F.cOctet)
	}

	var PaCOct *C.octet = nil
	if Pa != nil {
		PaCOct = (*C.octet)(&Pa.cOctet)
	}

	code := C.MPIN_BN254_SERVER_2(C.int(d), HIDCOct, HTIDCOct, yCOct, SSCOct, UCOct, UTCOct, VCOct, ECOct, FCOct, PaCOct)

	return newError(code)
}

// ServerKeyBN254 is a go wrapper for C.MPIN_BN254_SERVER_KEY.
func ServerKeyBN254(h int, Z *Octet, SS *Octet, w *Octet, p *Octet, I *Octet, U *Octet, UT *Octet, K *Octet) error {
	var ZCOct *C.octet = nil
	if Z != nil {
		ZCOct = (*C.octet)(&Z.cOctet)
	}

	var SSCOct *C.octet = nil
	if SS != nil {
		SSCOct = (*C.octet)(&SS.cOctet)
	}

	var wCOct *C.octet = nil
	if w != nil {
		wCOct = (*C.octet)(&w.cOctet)
	}

	var pCOct *C.octet = nil
	if p != nil {
		pCOct = (*C.octet)(&p.cOctet)
	}

	var ICOct *C.octet = nil
	if I != nil {
		ICOct = (*C.octet)(&I.cOctet)
	}

	var UCOct *C.octet = nil
	if U != nil {
		UCOct = (*C.octet)(&U.cOctet)
	}

	var UTCOct *C.octet = nil
	if UT != nil {
		UTCOct = (*C.octet)(&UT.cOctet)
	}

	var KCOct *C.octet = nil
	if K != nil {
		KCOct = (*C.octet)(&K.cOctet)
	}

	code := C.MPIN_BN254_SERVER_KEY(C.int(h), ZCOct, SSCOct, wCOct, pCOct, ICOct, UCOct, UTCOct, KCOct)

	return newError(code)
}

// ServerBN254 is a go wrapper for C.MPIN_BN254_SERVER.
func ServerBN254(h int, d int, HID *Octet, HTID *Octet, y *Octet, SS *Octet, U *Octet, UT *Octet, V *Octet, E *Octet, F *Octet, ID *Octet, MESSAGE *Octet, t int, Pa *Octet) error {

	var HIDCOct *C.octet = nil
	if HID != nil {
		HIDCOct = (*C.octet)(&HID.cOctet)
	}

	var HTIDCOct *C.octet = nil
	if HTID != nil {
		HTIDCOct = (*C.octet)(&HTID.cOctet)
	}

	var yCOct *C.octet = nil
	if y != nil {
		yCOct = (*C.octet)(&y.cOctet)
	}

	var SSCOct *C.octet = nil
	if SS != nil {
		SSCOct = (*C.octet)(&SS.cOctet)
	}

	var UCOct *C.octet = nil
	if U != nil {
		UCOct = (*C.octet)(&U.cOctet)
	}

	var UTCOct *C.octet = nil
	if UT != nil {
		UTCOct = (*C.octet)(&UT.cOctet)
	}

	var VCOct *C.octet = nil
	if V != nil {
		VCOct = (*C.octet)(&V.cOctet)
	}

	var ECOct *C.octet = nil
	if E != nil {
		ECOct = (*C.octet)(&E.cOctet)
	}

	var FCOct *C.octet = nil
	if F != nil {
		FCOct = (*C.octet)(&F.cOctet)
	}

	var IDCOct *C.octet = nil
	if ID != nil {
		IDCOct = (*C.octet)(&ID.cOctet)
	}

	var MESSAGECOct *C.octet = nil
	if MESSAGE != nil {
		MESSAGECOct = (*C.octet)(&MESSAGE.cOctet)
	}

	var PaCOct *C.octet = nil
	if Pa != nil {
		PaCOct = (*C.octet)(&Pa.cOctet)
	}

	code := C.MPIN_BN254_SERVER(C.int(h), C.int(d), HIDCOct, HTIDCOct, yCOct, SSCOct, UCOct, UTCOct, VCOct, ECOct, FCOct, IDCOct, MESSAGECOct, C.int(t), PaCOct)

	return newError(code)
}

// Server1BN254 is a go wrapper for C.MPIN_BN254_SERVER_1.
func Server1BN254(h int, d int, ID *Octet, HID *Octet, HTID *Octet) {

	var IDCOct *C.octet = nil
	if ID != nil {
		IDCOct = (*C.octet)(&ID.cOctet)
	}

	var HIDCOct *C.octet = nil
	if HID != nil {
		HIDCOct = (*C.octet)(&HID.cOctet)
	}

	var HTIDCOct *C.octet = nil
	if HTID != nil {
		HTIDCOct = (*C.octet)(&HTID.cOctet)
	}

	C.MPIN_BN254_SERVER_1(C.int(h), C.int(d), IDCOct, HIDCOct, HTIDCOct)
}
