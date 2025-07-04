// Generated by gen/wrappers/main.go from wcc_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || BN254

package wrap

import "code.miracl.com/maas/maas/src/lib/gomiracl/bindings"

// WCCRandomGenerateBN254 is a wrapper of bindings.WCC_BN254_RANDOM_GENERATE.
func WCCRandomGenerateBN254(RNG *bindings.Rand) (S []byte, err error) {

	SSize := bindings.WCCPGSBN254
	SOct := bindings.MakeOctet(SSize)
	defer SOct.Free()

	err = bindings.WCCRandomGenerateBN254(RNG, SOct)

	S = SOct.ToBytes()

	return
}

// WCCHqBN254 is a wrapper of bindings.WCC_BN254_Hq.
func WCCHqBN254(h int, A []byte, B []byte, Ci []byte, D []byte) (res []byte) {

	AOct := bindings.NewOctet(A)
	defer AOct.Free()

	BOct := bindings.NewOctet(B)
	defer BOct.Free()

	CiOct := bindings.NewOctet(Ci)
	defer CiOct.Free()

	DOct := bindings.NewOctet(D)
	defer DOct.Free()

	resSize := bindings.WCCPGSBN254
	resOct := bindings.MakeOctet(resSize)
	defer resOct.Free()

	bindings.WCCHqBN254(h, AOct, BOct, CiOct, DOct, resOct)

	res = resOct.ToBytes()

	return
}

// WCCGetG1MultipleBN254 is a wrapper of bindings.WCC_BN254_GET_G1_MULTIPLE.
func WCCGetG1MultipleBN254(S []byte, HID []byte) (VG1 []byte, err error) {

	SOct := bindings.NewOctet(S)
	defer SOct.Free()

	HIDOct := bindings.NewOctet(HID)
	defer HIDOct.Free()

	VG1Size := 2*bindings.WCCPFSBN254 + 1
	VG1Oct := bindings.MakeOctet(VG1Size)
	defer VG1Oct.Free()

	err = bindings.WCCGetG1MultipleBN254(SOct, HIDOct, VG1Oct)

	VG1 = VG1Oct.ToBytes()

	return
}

// WCCRecombineG1BN254 is a wrapper of bindings.WCC_BN254_RECOMBINE_G1.
func WCCRecombineG1BN254(R1 []byte, R2 []byte) (R []byte, err error) {

	R1Oct := bindings.NewOctet(R1)
	defer R1Oct.Free()

	R2Oct := bindings.NewOctet(R2)
	defer R2Oct.Free()

	RSize := 2*bindings.WCCPFSBN254 + 1
	ROct := bindings.MakeOctet(RSize)
	defer ROct.Free()

	err = bindings.WCCRecombineG1BN254(R1Oct, R2Oct, ROct)

	R = ROct.ToBytes()

	return
}

// WCCGetG2MultipleBN254 is a wrapper of bindings.WCC_BN254_GET_G2_MULTIPLE.
func WCCGetG2MultipleBN254(S []byte, HID []byte) (VG2 []byte, err error) {

	SOct := bindings.NewOctet(S)
	defer SOct.Free()

	HIDOct := bindings.NewOctet(HID)
	defer HIDOct.Free()

	VG2Size := 4 * bindings.WCCPFSBN254
	VG2Oct := bindings.MakeOctet(VG2Size)
	defer VG2Oct.Free()

	err = bindings.WCCGetG2MultipleBN254(SOct, HIDOct, VG2Oct)

	VG2 = VG2Oct.ToBytes()

	return
}

// WCCRecombineG2BN254 is a wrapper of bindings.WCC_BN254_RECOMBINE_G2.
func WCCRecombineG2BN254(W1 []byte, W2 []byte) (W []byte, err error) {

	W1Oct := bindings.NewOctet(W1)
	defer W1Oct.Free()

	W2Oct := bindings.NewOctet(W2)
	defer W2Oct.Free()

	WSize := 4 * bindings.WCCPFSBN254
	WOct := bindings.MakeOctet(WSize)
	defer WOct.Free()

	err = bindings.WCCRecombineG2BN254(W1Oct, W2Oct, WOct)

	W = WOct.ToBytes()

	return
}

// WCCReceiverKeyBN254 is a wrapper of bindings.WCC_BN254_RECEIVER_KEY.
func WCCReceiverKeyBN254(h int, y []byte, w []byte, pia []byte, pib []byte, PaG1 []byte, PgG1 []byte, BKeyG2 []byte, IDA []byte) (AESKey []byte, err error) {

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

	AESKeySize := bindings.AESKeySizeBN254
	AESKeyOct := bindings.MakeOctet(AESKeySize)
	defer AESKeyOct.Free()

	err = bindings.WCCReceiverKeyBN254(h, yOct, wOct, piaOct, pibOct, PaG1Oct, PgG1Oct, BKeyG2Oct, IDAOct, AESKeyOct)

	AESKey = AESKeyOct.ToBytes()

	return
}

// WCCSenderKeyBN254 is a wrapper of bindings.WCC_BN254_SENDER_KEY.
func WCCSenderKeyBN254(h int, x []byte, pia []byte, pib []byte, PbG2 []byte, PgG1 []byte, AKeyG1 []byte, IDB []byte) (AESKey []byte, err error) {

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

	AESKeySize := bindings.AESKeySizeBN254
	AESKeyOct := bindings.MakeOctet(AESKeySize)
	defer AESKeyOct.Free()

	err = bindings.WCCSenderKeyBN254(h, xOct, piaOct, pibOct, PbG2Oct, PgG1Oct, AKeyG1Oct, IDBOct, AESKeyOct)

	AESKey = AESKeyOct.ToBytes()

	return
}
