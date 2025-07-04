// Generated by gen/wrappers/main.go from wcc_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || FP256BN

package wrap

import "code.miracl.com/maas/maas/src/lib/gomiracl/bindings"

// WCCRandomGenerateFP256BN is a wrapper of bindings.WCC_FP256BN_RANDOM_GENERATE.
func WCCRandomGenerateFP256BN(RNG *bindings.Rand) (S []byte, err error) {

	SSize := bindings.WCCPGSFP256BN
	SOct := bindings.MakeOctet(SSize)
	defer SOct.Free()

	err = bindings.WCCRandomGenerateFP256BN(RNG, SOct)

	S = SOct.ToBytes()

	return
}

// WCCHqFP256BN is a wrapper of bindings.WCC_FP256BN_Hq.
func WCCHqFP256BN(h int, A []byte, B []byte, Ci []byte, D []byte) (res []byte) {

	AOct := bindings.NewOctet(A)
	defer AOct.Free()

	BOct := bindings.NewOctet(B)
	defer BOct.Free()

	CiOct := bindings.NewOctet(Ci)
	defer CiOct.Free()

	DOct := bindings.NewOctet(D)
	defer DOct.Free()

	resSize := bindings.WCCPGSFP256BN
	resOct := bindings.MakeOctet(resSize)
	defer resOct.Free()

	bindings.WCCHqFP256BN(h, AOct, BOct, CiOct, DOct, resOct)

	res = resOct.ToBytes()

	return
}

// WCCGetG1MultipleFP256BN is a wrapper of bindings.WCC_FP256BN_GET_G1_MULTIPLE.
func WCCGetG1MultipleFP256BN(S []byte, HID []byte) (VG1 []byte, err error) {

	SOct := bindings.NewOctet(S)
	defer SOct.Free()

	HIDOct := bindings.NewOctet(HID)
	defer HIDOct.Free()

	VG1Size := 2*bindings.WCCPFSFP256BN + 1
	VG1Oct := bindings.MakeOctet(VG1Size)
	defer VG1Oct.Free()

	err = bindings.WCCGetG1MultipleFP256BN(SOct, HIDOct, VG1Oct)

	VG1 = VG1Oct.ToBytes()

	return
}

// WCCRecombineG1FP256BN is a wrapper of bindings.WCC_FP256BN_RECOMBINE_G1.
func WCCRecombineG1FP256BN(R1 []byte, R2 []byte) (R []byte, err error) {

	R1Oct := bindings.NewOctet(R1)
	defer R1Oct.Free()

	R2Oct := bindings.NewOctet(R2)
	defer R2Oct.Free()

	RSize := 2*bindings.WCCPFSFP256BN + 1
	ROct := bindings.MakeOctet(RSize)
	defer ROct.Free()

	err = bindings.WCCRecombineG1FP256BN(R1Oct, R2Oct, ROct)

	R = ROct.ToBytes()

	return
}

// WCCGetG2MultipleFP256BN is a wrapper of bindings.WCC_FP256BN_GET_G2_MULTIPLE.
func WCCGetG2MultipleFP256BN(S []byte, HID []byte) (VG2 []byte, err error) {

	SOct := bindings.NewOctet(S)
	defer SOct.Free()

	HIDOct := bindings.NewOctet(HID)
	defer HIDOct.Free()

	VG2Size := 4 * bindings.WCCPFSFP256BN
	VG2Oct := bindings.MakeOctet(VG2Size)
	defer VG2Oct.Free()

	err = bindings.WCCGetG2MultipleFP256BN(SOct, HIDOct, VG2Oct)

	VG2 = VG2Oct.ToBytes()

	return
}

// WCCRecombineG2FP256BN is a wrapper of bindings.WCC_FP256BN_RECOMBINE_G2.
func WCCRecombineG2FP256BN(W1 []byte, W2 []byte) (W []byte, err error) {

	W1Oct := bindings.NewOctet(W1)
	defer W1Oct.Free()

	W2Oct := bindings.NewOctet(W2)
	defer W2Oct.Free()

	WSize := 4 * bindings.WCCPFSFP256BN
	WOct := bindings.MakeOctet(WSize)
	defer WOct.Free()

	err = bindings.WCCRecombineG2FP256BN(W1Oct, W2Oct, WOct)

	W = WOct.ToBytes()

	return
}

// WCCReceiverKeyFP256BN is a wrapper of bindings.WCC_FP256BN_RECEIVER_KEY.
func WCCReceiverKeyFP256BN(h int, y []byte, w []byte, pia []byte, pib []byte, PaG1 []byte, PgG1 []byte, BKeyG2 []byte, IDA []byte) (AESKey []byte, err error) {

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

	AESKeySize := bindings.AESKeySizeFP256BN
	AESKeyOct := bindings.MakeOctet(AESKeySize)
	defer AESKeyOct.Free()

	err = bindings.WCCReceiverKeyFP256BN(h, yOct, wOct, piaOct, pibOct, PaG1Oct, PgG1Oct, BKeyG2Oct, IDAOct, AESKeyOct)

	AESKey = AESKeyOct.ToBytes()

	return
}

// WCCSenderKeyFP256BN is a wrapper of bindings.WCC_FP256BN_SENDER_KEY.
func WCCSenderKeyFP256BN(h int, x []byte, pia []byte, pib []byte, PbG2 []byte, PgG1 []byte, AKeyG1 []byte, IDB []byte) (AESKey []byte, err error) {

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

	AESKeySize := bindings.AESKeySizeFP256BN
	AESKeyOct := bindings.MakeOctet(AESKeySize)
	defer AESKeyOct.Free()

	err = bindings.WCCSenderKeyFP256BN(h, xOct, piaOct, pibOct, PbG2Oct, PgG1Oct, AKeyG1Oct, IDBOct, AESKeyOct)

	AESKey = AESKeyOct.ToBytes()

	return
}
