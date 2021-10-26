// Generated by gen/wrappers/main.go from wcc_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || BLS383

package wrap

import "code.miracl.com/mfa/pkg/gomiracl/bindings"

// WCCRandomGenerateBLS383 is a wrapper of bindings.WCC_BLS383_RANDOM_GENERATE.
func WCCRandomGenerateBLS383(RNG *bindings.Rand) (S []byte, err error) {

	SSize := bindings.WCCPGSBLS383
	SOct := bindings.MakeOctet(SSize)
	defer SOct.Free()

	err = bindings.WCCRandomGenerateBLS383(RNG, SOct)

	S = SOct.ToBytes()

	return
}

// WCCHqBLS383 is a wrapper of bindings.WCC_BLS383_Hq.
func WCCHqBLS383(h int, A []byte, B []byte, C []byte, D []byte) (res []byte) {

	AOct := bindings.NewOctet(A)
	defer AOct.Free()

	BOct := bindings.NewOctet(B)
	defer BOct.Free()

	COct := bindings.NewOctet(C)
	defer COct.Free()

	DOct := bindings.NewOctet(D)
	defer DOct.Free()

	resSize := bindings.WCCPGSBLS383
	resOct := bindings.MakeOctet(resSize)
	defer resOct.Free()

	bindings.WCCHqBLS383(h, AOct, BOct, COct, DOct, resOct)

	res = resOct.ToBytes()

	return
}

// WCCGetG1MultipleBLS383 is a wrapper of bindings.WCC_BLS383_GET_G1_MULTIPLE.
func WCCGetG1MultipleBLS383(S []byte, HID []byte) (VG1 []byte, err error) {

	SOct := bindings.NewOctet(S)
	defer SOct.Free()

	HIDOct := bindings.NewOctet(HID)
	defer HIDOct.Free()

	VG1Size := 2*bindings.WCCPFSBLS383 + 1
	VG1Oct := bindings.MakeOctet(VG1Size)
	defer VG1Oct.Free()

	err = bindings.WCCGetG1MultipleBLS383(SOct, HIDOct, VG1Oct)

	VG1 = VG1Oct.ToBytes()

	return
}

// WCCRecombineG1BLS383 is a wrapper of bindings.WCC_BLS383_RECOMBINE_G1.
func WCCRecombineG1BLS383(R1 []byte, R2 []byte) (R []byte, err error) {

	R1Oct := bindings.NewOctet(R1)
	defer R1Oct.Free()

	R2Oct := bindings.NewOctet(R2)
	defer R2Oct.Free()

	RSize := 2*bindings.WCCPFSBLS383 + 1
	ROct := bindings.MakeOctet(RSize)
	defer ROct.Free()

	err = bindings.WCCRecombineG1BLS383(R1Oct, R2Oct, ROct)

	R = ROct.ToBytes()

	return
}

// WCCGetG2MultipleBLS383 is a wrapper of bindings.WCC_BLS383_GET_G2_MULTIPLE.
func WCCGetG2MultipleBLS383(S []byte, HID []byte) (VG2 []byte, err error) {

	SOct := bindings.NewOctet(S)
	defer SOct.Free()

	HIDOct := bindings.NewOctet(HID)
	defer HIDOct.Free()

	VG2Size := 4 * bindings.WCCPFSBLS383
	VG2Oct := bindings.MakeOctet(VG2Size)
	defer VG2Oct.Free()

	err = bindings.WCCGetG2MultipleBLS383(SOct, HIDOct, VG2Oct)

	VG2 = VG2Oct.ToBytes()

	return
}

// WCCRecombineG2BLS383 is a wrapper of bindings.WCC_BLS383_RECOMBINE_G2.
func WCCRecombineG2BLS383(W1 []byte, W2 []byte) (W []byte, err error) {

	W1Oct := bindings.NewOctet(W1)
	defer W1Oct.Free()

	W2Oct := bindings.NewOctet(W2)
	defer W2Oct.Free()

	WSize := 4 * bindings.WCCPFSBLS383
	WOct := bindings.MakeOctet(WSize)
	defer WOct.Free()

	err = bindings.WCCRecombineG2BLS383(W1Oct, W2Oct, WOct)

	W = WOct.ToBytes()

	return
}

// WCCReceiverKeyBLS383 is a wrapper of bindings.WCC_BLS383_RECEIVER_KEY.
func WCCReceiverKeyBLS383(h int, y []byte, w []byte, pia []byte, pib []byte, PaG1 []byte, PgG1 []byte, BKeyG2 []byte, IDA []byte) (AESKey []byte, err error) {

	yOct := bindings.NewOctet(y)
	defer yOct.Free()

	wOct := bindings.NewOctet(w)
	defer wOct.Free()

	piaOct := bindings.NewOctet(pia)
	defer piaOct.Free()

	pibOct := bindings.NewOctet(pib)
	defer pibOct.Free()

	PaG1Oct := bindings.NewOctet(PaG1)
	defer PaG1Oct.Free()

	PgG1Oct := bindings.NewOctet(PgG1)
	defer PgG1Oct.Free()

	BKeyG2Oct := bindings.NewOctet(BKeyG2)
	defer BKeyG2Oct.Free()

	IDAOct := bindings.NewOctet(IDA)
	defer IDAOct.Free()

	AESKeySize := bindings.AESKeySizeBLS383
	AESKeyOct := bindings.MakeOctet(AESKeySize)
	defer AESKeyOct.Free()

	err = bindings.WCCReceiverKeyBLS383(h, yOct, wOct, piaOct, pibOct, PaG1Oct, PgG1Oct, BKeyG2Oct, IDAOct, AESKeyOct)

	AESKey = AESKeyOct.ToBytes()

	return
}

// WCCSenderKeyBLS383 is a wrapper of bindings.WCC_BLS383_SENDER_KEY.
func WCCSenderKeyBLS383(h int, x []byte, pia []byte, pib []byte, PbG2 []byte, PgG1 []byte, AKeyG1 []byte, IDB []byte) (AESKey []byte, err error) {

	xOct := bindings.NewOctet(x)
	defer xOct.Free()

	piaOct := bindings.NewOctet(pia)
	defer piaOct.Free()

	pibOct := bindings.NewOctet(pib)
	defer pibOct.Free()

	PbG2Oct := bindings.NewOctet(PbG2)
	defer PbG2Oct.Free()

	PgG1Oct := bindings.NewOctet(PgG1)
	defer PgG1Oct.Free()

	AKeyG1Oct := bindings.NewOctet(AKeyG1)
	defer AKeyG1Oct.Free()

	IDBOct := bindings.NewOctet(IDB)
	defer IDBOct.Free()

	AESKeySize := bindings.AESKeySizeBLS383
	AESKeyOct := bindings.MakeOctet(AESKeySize)
	defer AESKeyOct.Free()

	err = bindings.WCCSenderKeyBLS383(h, xOct, piaOct, pibOct, PbG2Oct, PgG1Oct, AKeyG1Oct, IDBOct, AESKeyOct)

	AESKey = AESKeyOct.ToBytes()

	return
}