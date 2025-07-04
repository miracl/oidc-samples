// Generated by gen/wrappers/main.go from mpin_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || FP256BN

package wrap

import "code.miracl.com/maas/maas/src/lib/gomiracl/bindings"

// Client1FP256BN is a wrapper of bindings.MPIN_FP256BN_CLIENT_1.
func Client1FP256BN(h int, d int, ID []byte, R *bindings.Rand, x []byte, pin int, T []byte, TP []byte) (xResult []byte, S []byte, U []byte, UT []byte, err error) {

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	var xOct *bindings.Octet
	if x != nil {
		xOct = bindings.NewOctet(x)
		defer xOct.Free()
	} else {
		xSize := bindings.PGSFP256BN
		xOct = bindings.MakeOctet(xSize)
		defer xOct.Free()
	}

	TOct := bindings.NewOctet(T)
	defer TOct.Free()

	SSize := bindings.G1SFP256BN
	SOct := bindings.MakeOctet(SSize)
	defer SOct.Free()

	USize := bindings.G1SFP256BN
	UOct := bindings.MakeOctet(USize)
	defer UOct.Free()

	UTSize := bindings.G1SFP256BN
	UTOct := bindings.MakeOctet(UTSize)
	defer UTOct.Free()

	TPOct := bindings.NewOctet(TP)
	defer TPOct.Free()

	err = bindings.Client1FP256BN(h, d, IDOct, R, xOct, pin, TOct, SOct, UOct, UTOct, TPOct)

	xResult = xOct.ToBytes()

	S = SOct.ToBytes()

	U = UOct.ToBytes()

	UT = UTOct.ToBytes()

	return
}

// Client2FP256BN is a wrapper of bindings.MPIN_FP256BN_CLIENT_2.
func Client2FP256BN(x []byte, y []byte, V []byte) (VResult []byte, err error) {

	xOct := bindings.NewOctet(x)
	defer xOct.Free()

	yOct := bindings.NewOctet(y)
	defer yOct.Free()

	VOct := bindings.NewOctet(V)
	defer VOct.Free()

	err = bindings.Client2FP256BN(xOct, yOct, VOct)

	VResult = VOct.ToBytes()

	return
}

// ClientKeyFP256BN is a wrapper of bindings.MPIN_FP256BN_CLIENT_KEY.
func ClientKeyFP256BN(h int, g1 []byte, g2 []byte, pin int, r []byte, x []byte, p []byte, T []byte) (K []byte, err error) {

	g1Oct := bindings.NewOctet(g1)
	defer g1Oct.Free()

	g2Oct := bindings.NewOctet(g2)
	defer g2Oct.Free()

	rOct := bindings.NewOctet(r)
	defer rOct.Free()

	xOct := bindings.NewOctet(x)
	defer xOct.Free()

	pOct := bindings.NewOctet(p)
	defer pOct.Free()

	TOct := bindings.NewOctet(T)
	defer TOct.Free()

	KSize := bindings.PASFP256BN
	KOct := bindings.MakeOctet(KSize)
	defer KOct.Free()

	err = bindings.ClientKeyFP256BN(h, g1Oct, g2Oct, pin, rOct, xOct, pOct, TOct, KOct)

	K = KOct.ToBytes()

	return
}

// ClientFP256BN is a wrapper of bindings.MPIN_FP256BN_CLIENT.
func ClientFP256BN(h int, d int, ID []byte, R *bindings.Rand, x []byte, pin int, T []byte, TP []byte, MESSAGE []byte, t int) (xResult []byte, V []byte, U []byte, UT []byte, y []byte, err error) {

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	var xOct *bindings.Octet
	if x != nil {
		xOct = bindings.NewOctet(x)
		defer xOct.Free()
	} else {
		xSize := bindings.PGSFP256BN
		xOct = bindings.MakeOctet(xSize)
		defer xOct.Free()
	}

	TOct := bindings.NewOctet(T)
	defer TOct.Free()

	VSize := bindings.G1SFP256BN
	VOct := bindings.MakeOctet(VSize)
	defer VOct.Free()

	USize := bindings.G1SFP256BN
	UOct := bindings.MakeOctet(USize)
	defer UOct.Free()

	UTSize := bindings.G1SFP256BN
	UTOct := bindings.MakeOctet(UTSize)
	defer UTOct.Free()

	TPOct := bindings.NewOctet(TP)
	defer TPOct.Free()

	MESSAGEOct := bindings.NewOctet(MESSAGE)
	defer MESSAGEOct.Free()

	ySize := bindings.PGSFP256BN
	yOct := bindings.MakeOctet(ySize)
	defer yOct.Free()

	err = bindings.ClientFP256BN(h, d, IDOct, R, xOct, pin, TOct, VOct, UOct, UTOct, TPOct, MESSAGEOct, t, yOct)

	xResult = xOct.ToBytes()

	V = VOct.ToBytes()

	U = UOct.ToBytes()

	UT = UTOct.ToBytes()

	y = yOct.ToBytes()

	return
}

// ExtractPINFP256BN is a wrapper of bindings.MPIN_FP256BN_EXTRACT_PIN.
func ExtractPINFP256BN(h int, ID []byte, pin int, CS []byte) (CSResult []byte, err error) {

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	CSOct := bindings.NewOctet(CS)
	defer CSOct.Free()

	err = bindings.ExtractPINFP256BN(h, IDOct, pin, CSOct)

	CSResult = CSOct.ToBytes()

	return
}

// GetClientPermitFP256BN is a wrapper of bindings.MPIN_FP256BN_GET_CLIENT_PERMIT.
func GetClientPermitFP256BN(h int, d int, S []byte, ID []byte) (TP []byte, err error) {

	SOct := bindings.NewOctet(S)
	defer SOct.Free()

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	TPSize := bindings.G1SFP256BN
	TPOct := bindings.MakeOctet(TPSize)
	defer TPOct.Free()

	err = bindings.GetClientPermitFP256BN(h, d, SOct, IDOct, TPOct)

	TP = TPOct.ToBytes()

	return
}

// GetClientSecretFP256BN is a wrapper of bindings.MPIN_FP256BN_GET_CLIENT_SECRET.
func GetClientSecretFP256BN(S []byte, ID []byte) (CS []byte, err error) {

	SOct := bindings.NewOctet(S)
	defer SOct.Free()

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	CSSize := bindings.G1SFP256BN
	CSOct := bindings.MakeOctet(CSSize)
	defer CSOct.Free()

	err = bindings.GetClientSecretFP256BN(SOct, IDOct, CSOct)

	CS = CSOct.ToBytes()

	return
}

// GetDVSKeyPairFP256BN is a wrapper of bindings.MPIN_FP256BN_GET_DVS_KEYPAIR.
func GetDVSKeyPairFP256BN(R *bindings.Rand, Z []byte) (ZResult []byte, Pa []byte, err error) {

	var ZOct *bindings.Octet
	if Z != nil {
		ZOct = bindings.NewOctet(Z)
		defer ZOct.Free()
	} else {
		ZSize := bindings.PGSFP256BN
		ZOct = bindings.MakeOctet(ZSize)
		defer ZOct.Free()
	}

	PaSize := bindings.G2SFP256BN
	PaOct := bindings.MakeOctet(PaSize)
	defer PaOct.Free()

	err = bindings.GetDVSKeyPairFP256BN(R, ZOct, PaOct)

	ZResult = ZOct.ToBytes()

	Pa = PaOct.ToBytes()

	return
}

// GetG1MultipleFP256BN is a wrapper of bindings.MPIN_FP256BN_GET_G1_MULTIPLE.
func GetG1MultipleFP256BN(R *bindings.Rand, t int, x []byte, G []byte) (xResult []byte, W []byte, err error) {

	var xOct *bindings.Octet
	if x != nil {
		xOct = bindings.NewOctet(x)
		defer xOct.Free()
	} else {
		xSize := bindings.PGSFP256BN
		xOct = bindings.MakeOctet(xSize)
		defer xOct.Free()
	}

	GOct := bindings.NewOctet(G)
	defer GOct.Free()

	WSize := bindings.G1SFP256BN
	WOct := bindings.MakeOctet(WSize)
	defer WOct.Free()

	err = bindings.GetG1MultipleFP256BN(R, t, xOct, GOct, WOct)

	xResult = xOct.ToBytes()

	W = WOct.ToBytes()

	return
}

// GetServerSecretFP256BN is a wrapper of bindings.MPIN_FP256BN_GET_SERVER_SECRET.
func GetServerSecretFP256BN(S []byte) (SS []byte, err error) {

	SOct := bindings.NewOctet(S)
	defer SOct.Free()

	SSSize := bindings.G2SFP256BN
	SSOct := bindings.MakeOctet(SSSize)
	defer SSOct.Free()

	err = bindings.GetServerSecretFP256BN(SOct, SSOct)

	SS = SSOct.ToBytes()

	return
}

// KangarooFP256BN is a wrapper of bindings.MPIN_FP256BN_KANGAROO.
func KangarooFP256BN(E []byte, F []byte) (err error) {

	EOct := bindings.NewOctet(E)
	defer EOct.Free()

	FOct := bindings.NewOctet(F)
	defer FOct.Free()

	err = bindings.KangarooFP256BN(EOct, FOct)

	return
}

// PrecomputeFP256BN is a wrapper of bindings.MPIN_FP256BN_PRECOMPUTE.
func PrecomputeFP256BN(T []byte, ID []byte, CP []byte) (g1 []byte, g2 []byte, err error) {

	TOct := bindings.NewOctet(T)
	defer TOct.Free()

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	CPOct := bindings.NewOctet(CP)
	defer CPOct.Free()

	g1Size := bindings.GTSFP256BN
	g1Oct := bindings.MakeOctet(g1Size)
	defer g1Oct.Free()

	g2Size := bindings.GTSFP256BN
	g2Oct := bindings.MakeOctet(g2Size)
	defer g2Oct.Free()

	err = bindings.PrecomputeFP256BN(TOct, IDOct, CPOct, g1Oct, g2Oct)

	g1 = g1Oct.ToBytes()

	g2 = g2Oct.ToBytes()

	return
}

// RandomGenerateFP256BN is a wrapper of bindings.MPIN_FP256BN_RANDOM_GENERATE.
func RandomGenerateFP256BN(R *bindings.Rand) (S []byte, err error) {

	SSize := bindings.PGSFP256BN
	SOct := bindings.MakeOctet(SSize)
	defer SOct.Free()

	err = bindings.RandomGenerateFP256BN(R, SOct)

	S = SOct.ToBytes()

	return
}

// RecombineG1FP256BN is a wrapper of bindings.MPIN_FP256BN_RECOMBINE_G1.
func RecombineG1FP256BN(Q1 []byte, Q2 []byte) (Q []byte, err error) {

	Q1Oct := bindings.NewOctet(Q1)
	defer Q1Oct.Free()

	Q2Oct := bindings.NewOctet(Q2)
	defer Q2Oct.Free()

	QSize := bindings.G1SFP256BN
	QOct := bindings.MakeOctet(QSize)
	defer QOct.Free()

	err = bindings.RecombineG1FP256BN(Q1Oct, Q2Oct, QOct)

	Q = QOct.ToBytes()

	return
}

// RecombineG2FP256BN is a wrapper of bindings.MPIN_FP256BN_RECOMBINE_G2.
func RecombineG2FP256BN(P1 []byte, P2 []byte) (P []byte, err error) {

	P1Oct := bindings.NewOctet(P1)
	defer P1Oct.Free()

	P2Oct := bindings.NewOctet(P2)
	defer P2Oct.Free()

	PSize := bindings.G2SFP256BN
	POct := bindings.MakeOctet(PSize)
	defer POct.Free()

	err = bindings.RecombineG2FP256BN(P1Oct, P2Oct, POct)

	P = POct.ToBytes()

	return
}

// Server2FP256BN is a wrapper of bindings.MPIN_FP256BN_SERVER_2.
func Server2FP256BN(d int, HID []byte, HTID []byte, y []byte, SS []byte, U []byte, UT []byte, V []byte, Pa []byte) (err error) {

	HIDOct := bindings.NewOctet(HID)
	defer HIDOct.Free()

	HTIDOct := bindings.NewOctet(HTID)
	defer HTIDOct.Free()

	yOct := bindings.NewOctet(y)
	defer yOct.Free()

	SSOct := bindings.NewOctet(SS)
	defer SSOct.Free()

	UOct := bindings.NewOctet(U)
	defer UOct.Free()

	UTOct := bindings.NewOctet(UT)
	defer UTOct.Free()

	VOct := bindings.NewOctet(V)
	defer VOct.Free()

	PaOct := bindings.NewOctet(Pa)
	defer PaOct.Free()

	err = bindings.Server2FP256BN(d, HIDOct, HTIDOct, yOct, SSOct, UOct, UTOct, VOct, nil, nil, PaOct)

	return
}

// ServerKeyFP256BN is a wrapper of bindings.MPIN_FP256BN_SERVER_KEY.
func ServerKeyFP256BN(h int, Z []byte, SS []byte, w []byte, p []byte, I []byte, U []byte, UT []byte) (K []byte, err error) {

	ZOct := bindings.NewOctet(Z)
	defer ZOct.Free()

	SSOct := bindings.NewOctet(SS)
	defer SSOct.Free()

	wOct := bindings.NewOctet(w)
	defer wOct.Free()

	pOct := bindings.NewOctet(p)
	defer pOct.Free()

	IOct := bindings.NewOctet(I)
	defer IOct.Free()

	UOct := bindings.NewOctet(U)
	defer UOct.Free()

	UTOct := bindings.NewOctet(UT)
	defer UTOct.Free()

	KSize := bindings.PASFP256BN
	KOct := bindings.MakeOctet(KSize)
	defer KOct.Free()

	err = bindings.ServerKeyFP256BN(h, ZOct, SSOct, wOct, pOct, IOct, UOct, UTOct, KOct)

	K = KOct.ToBytes()

	return
}

// ServerFP256BN is a wrapper of bindings.MPIN_FP256BN_SERVER.
func ServerFP256BN(h int, d int, SS []byte, U []byte, UT []byte, V []byte, ID []byte, MESSAGE []byte, t int, Pa []byte) (HID []byte, HTID []byte, y []byte, err error) {

	HIDSize := bindings.G1SFP256BN
	HIDOct := bindings.MakeOctet(HIDSize)
	defer HIDOct.Free()

	HTIDSize := bindings.G1SFP256BN
	HTIDOct := bindings.MakeOctet(HTIDSize)
	defer HTIDOct.Free()

	ySize := bindings.PGSFP256BN
	yOct := bindings.MakeOctet(ySize)
	defer yOct.Free()

	SSOct := bindings.NewOctet(SS)
	defer SSOct.Free()

	UOct := bindings.NewOctet(U)
	defer UOct.Free()

	UTOct := bindings.NewOctet(UT)
	defer UTOct.Free()

	VOct := bindings.NewOctet(V)
	defer VOct.Free()

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	MESSAGEOct := bindings.NewOctet(MESSAGE)
	defer MESSAGEOct.Free()

	PaOct := bindings.NewOctet(Pa)
	defer PaOct.Free()

	err = bindings.ServerFP256BN(h, d, HIDOct, HTIDOct, yOct, SSOct, UOct, UTOct, VOct, nil, nil, IDOct, MESSAGEOct, t, PaOct)

	HID = HIDOct.ToBytes()

	HTID = HTIDOct.ToBytes()

	y = yOct.ToBytes()

	return
}

// Server1FP256BN is a wrapper of bindings.MPIN_FP256BN_SERVER_1.
func Server1FP256BN(h int, d int, ID []byte) (HID []byte, HTID []byte) {

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	HIDSize := bindings.G1SFP256BN
	HIDOct := bindings.MakeOctet(HIDSize)
	defer HIDOct.Free()

	HTIDSize := bindings.G1SFP256BN
	HTIDOct := bindings.MakeOctet(HTIDSize)
	defer HTIDOct.Free()

	bindings.Server1FP256BN(h, d, IDOct, HIDOct, HTIDOct)

	HID = HIDOct.ToBytes()

	HTID = HTIDOct.ToBytes()

	return
}
