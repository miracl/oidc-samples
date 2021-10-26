// Generated by gen/wrappers/main.go from mpin_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || BLS461

package wrap

import "code.miracl.com/mfa/pkg/gomiracl/bindings"

// Client1BLS461 is a wrapper of bindings.MPIN_BLS461_CLIENT_1.
func Client1BLS461(h int, d int, ID []byte, R *bindings.Rand, x []byte, pin int, T []byte, TP []byte) (xResult []byte, S []byte, U []byte, UT []byte, err error) {

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	var xOct *bindings.Octet
	if x != nil {
		xOct = bindings.NewOctet(x)
		defer xOct.Free()
	} else {
		xSize := bindings.PGSBLS461
		xOct = bindings.MakeOctet(xSize)
		defer xOct.Free()
	}

	TOct := bindings.NewOctet(T)
	defer TOct.Free()

	SSize := bindings.G1SBLS461
	SOct := bindings.MakeOctet(SSize)
	defer SOct.Free()

	USize := bindings.G1SBLS461
	UOct := bindings.MakeOctet(USize)
	defer UOct.Free()

	UTSize := bindings.G1SBLS461
	UTOct := bindings.MakeOctet(UTSize)
	defer UTOct.Free()

	TPOct := bindings.NewOctet(TP)
	defer TPOct.Free()

	err = bindings.Client1BLS461(h, d, IDOct, R, xOct, pin, TOct, SOct, UOct, UTOct, TPOct)

	xResult = xOct.ToBytes()

	S = SOct.ToBytes()

	U = UOct.ToBytes()

	UT = UTOct.ToBytes()

	return
}

// Client2BLS461 is a wrapper of bindings.MPIN_BLS461_CLIENT_2.
func Client2BLS461(x []byte, y []byte, V []byte) (VResult []byte, err error) {

	xOct := bindings.NewOctet(x)
	defer xOct.Free()

	yOct := bindings.NewOctet(y)
	defer yOct.Free()

	VOct := bindings.NewOctet(V)
	defer VOct.Free()

	err = bindings.Client2BLS461(xOct, yOct, VOct)

	VResult = VOct.ToBytes()

	return
}

// ClientKeyBLS461 is a wrapper of bindings.MPIN_BLS461_CLIENT_KEY.
func ClientKeyBLS461(h int, g1 []byte, g2 []byte, pin int, r []byte, x []byte, p []byte, T []byte) (K []byte, err error) {

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

	KSize := bindings.PASBLS461
	KOct := bindings.MakeOctet(KSize)
	defer KOct.Free()

	err = bindings.ClientKeyBLS461(h, g1Oct, g2Oct, pin, rOct, xOct, pOct, TOct, KOct)

	K = KOct.ToBytes()

	return
}

// ClientBLS461 is a wrapper of bindings.MPIN_BLS461_CLIENT.
func ClientBLS461(h int, d int, ID []byte, R *bindings.Rand, x []byte, pin int, T []byte, TP []byte, MESSAGE []byte, t int) (xResult []byte, V []byte, U []byte, UT []byte, y []byte, err error) {

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	var xOct *bindings.Octet
	if x != nil {
		xOct = bindings.NewOctet(x)
		defer xOct.Free()
	} else {
		xSize := bindings.PGSBLS461
		xOct = bindings.MakeOctet(xSize)
		defer xOct.Free()
	}

	TOct := bindings.NewOctet(T)
	defer TOct.Free()

	VSize := bindings.G1SBLS461
	VOct := bindings.MakeOctet(VSize)
	defer VOct.Free()

	USize := bindings.G1SBLS461
	UOct := bindings.MakeOctet(USize)
	defer UOct.Free()

	UTSize := bindings.G1SBLS461
	UTOct := bindings.MakeOctet(UTSize)
	defer UTOct.Free()

	TPOct := bindings.NewOctet(TP)
	defer TPOct.Free()

	MESSAGEOct := bindings.NewOctet(MESSAGE)
	defer MESSAGEOct.Free()

	ySize := bindings.PGSBLS461
	yOct := bindings.MakeOctet(ySize)
	defer yOct.Free()

	err = bindings.ClientBLS461(h, d, IDOct, R, xOct, pin, TOct, VOct, UOct, UTOct, TPOct, MESSAGEOct, t, yOct)

	xResult = xOct.ToBytes()

	V = VOct.ToBytes()

	U = UOct.ToBytes()

	UT = UTOct.ToBytes()

	y = yOct.ToBytes()

	return
}

// ExtractPINBLS461 is a wrapper of bindings.MPIN_BLS461_EXTRACT_PIN.
func ExtractPINBLS461(h int, ID []byte, pin int, CS []byte) (CSResult []byte, err error) {

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	CSOct := bindings.NewOctet(CS)
	defer CSOct.Free()

	err = bindings.ExtractPINBLS461(h, IDOct, pin, CSOct)

	CSResult = CSOct.ToBytes()

	return
}

// GetClientPermitBLS461 is a wrapper of bindings.MPIN_BLS461_GET_CLIENT_PERMIT.
func GetClientPermitBLS461(h int, d int, S []byte, ID []byte) (TP []byte, err error) {

	SOct := bindings.NewOctet(S)
	defer SOct.Free()

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	TPSize := bindings.G1SBLS461
	TPOct := bindings.MakeOctet(TPSize)
	defer TPOct.Free()

	err = bindings.GetClientPermitBLS461(h, d, SOct, IDOct, TPOct)

	TP = TPOct.ToBytes()

	return
}

// GetClientSecretBLS461 is a wrapper of bindings.MPIN_BLS461_GET_CLIENT_SECRET.
func GetClientSecretBLS461(S []byte, ID []byte) (CS []byte, err error) {

	SOct := bindings.NewOctet(S)
	defer SOct.Free()

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	CSSize := bindings.G1SBLS461
	CSOct := bindings.MakeOctet(CSSize)
	defer CSOct.Free()

	err = bindings.GetClientSecretBLS461(SOct, IDOct, CSOct)

	CS = CSOct.ToBytes()

	return
}

// GetDVSKeyPairBLS461 is a wrapper of bindings.MPIN_BLS461_GET_DVS_KEYPAIR.
func GetDVSKeyPairBLS461(R *bindings.Rand, Z []byte) (ZResult []byte, Pa []byte, err error) {

	var ZOct *bindings.Octet
	if Z != nil {
		ZOct = bindings.NewOctet(Z)
		defer ZOct.Free()
	} else {
		ZSize := bindings.PGSBLS461
		ZOct = bindings.MakeOctet(ZSize)
		defer ZOct.Free()
	}

	PaSize := bindings.G2SBLS461
	PaOct := bindings.MakeOctet(PaSize)
	defer PaOct.Free()

	err = bindings.GetDVSKeyPairBLS461(R, ZOct, PaOct)

	ZResult = ZOct.ToBytes()

	Pa = PaOct.ToBytes()

	return
}

// GetG1MultipleBLS461 is a wrapper of bindings.MPIN_BLS461_GET_G1_MULTIPLE.
func GetG1MultipleBLS461(R *bindings.Rand, t int, x []byte, G []byte) (xResult []byte, W []byte, err error) {

	var xOct *bindings.Octet
	if x != nil {
		xOct = bindings.NewOctet(x)
		defer xOct.Free()
	} else {
		xSize := bindings.PGSBLS461
		xOct = bindings.MakeOctet(xSize)
		defer xOct.Free()
	}

	GOct := bindings.NewOctet(G)
	defer GOct.Free()

	WSize := bindings.G1SBLS461
	WOct := bindings.MakeOctet(WSize)
	defer WOct.Free()

	err = bindings.GetG1MultipleBLS461(R, t, xOct, GOct, WOct)

	xResult = xOct.ToBytes()

	W = WOct.ToBytes()

	return
}

// GetServerSecretBLS461 is a wrapper of bindings.MPIN_BLS461_GET_SERVER_SECRET.
func GetServerSecretBLS461(S []byte) (SS []byte, err error) {

	SOct := bindings.NewOctet(S)
	defer SOct.Free()

	SSSize := bindings.G2SBLS461
	SSOct := bindings.MakeOctet(SSSize)
	defer SSOct.Free()

	err = bindings.GetServerSecretBLS461(SOct, SSOct)

	SS = SSOct.ToBytes()

	return
}

// KangarooBLS461 is a wrapper of bindings.MPIN_BLS461_KANGAROO.
func KangarooBLS461(E []byte, F []byte) (err error) {

	EOct := bindings.NewOctet(E)
	defer EOct.Free()

	FOct := bindings.NewOctet(F)
	defer FOct.Free()

	err = bindings.KangarooBLS461(EOct, FOct)

	return
}

// PrecomputeBLS461 is a wrapper of bindings.MPIN_BLS461_PRECOMPUTE.
func PrecomputeBLS461(T []byte, ID []byte, CP []byte) (g1 []byte, g2 []byte, err error) {

	TOct := bindings.NewOctet(T)
	defer TOct.Free()

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	CPOct := bindings.NewOctet(CP)
	defer CPOct.Free()

	g1Size := bindings.GTSBLS461
	g1Oct := bindings.MakeOctet(g1Size)
	defer g1Oct.Free()

	g2Size := bindings.GTSBLS461
	g2Oct := bindings.MakeOctet(g2Size)
	defer g2Oct.Free()

	err = bindings.PrecomputeBLS461(TOct, IDOct, CPOct, g1Oct, g2Oct)

	g1 = g1Oct.ToBytes()

	g2 = g2Oct.ToBytes()

	return
}

// RandomGenerateBLS461 is a wrapper of bindings.MPIN_BLS461_RANDOM_GENERATE.
func RandomGenerateBLS461(R *bindings.Rand) (S []byte, err error) {

	SSize := bindings.PGSBLS461
	SOct := bindings.MakeOctet(SSize)
	defer SOct.Free()

	err = bindings.RandomGenerateBLS461(R, SOct)

	S = SOct.ToBytes()

	return
}

// RecombineG1BLS461 is a wrapper of bindings.MPIN_BLS461_RECOMBINE_G1.
func RecombineG1BLS461(Q1 []byte, Q2 []byte) (Q []byte, err error) {

	Q1Oct := bindings.NewOctet(Q1)
	defer Q1Oct.Free()

	Q2Oct := bindings.NewOctet(Q2)
	defer Q2Oct.Free()

	QSize := bindings.G1SBLS461
	QOct := bindings.MakeOctet(QSize)
	defer QOct.Free()

	err = bindings.RecombineG1BLS461(Q1Oct, Q2Oct, QOct)

	Q = QOct.ToBytes()

	return
}

// RecombineG2BLS461 is a wrapper of bindings.MPIN_BLS461_RECOMBINE_G2.
func RecombineG2BLS461(P1 []byte, P2 []byte) (P []byte, err error) {

	P1Oct := bindings.NewOctet(P1)
	defer P1Oct.Free()

	P2Oct := bindings.NewOctet(P2)
	defer P2Oct.Free()

	PSize := bindings.G2SBLS461
	POct := bindings.MakeOctet(PSize)
	defer POct.Free()

	err = bindings.RecombineG2BLS461(P1Oct, P2Oct, POct)

	P = POct.ToBytes()

	return
}

// Server2BLS461 is a wrapper of bindings.MPIN_BLS461_SERVER_2.
func Server2BLS461(d int, HID []byte, HTID []byte, y []byte, SS []byte, U []byte, UT []byte, V []byte, Pa []byte) (err error) {

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

	err = bindings.Server2BLS461(d, HIDOct, HTIDOct, yOct, SSOct, UOct, UTOct, VOct, nil, nil, PaOct)

	return
}

// ServerKeyBLS461 is a wrapper of bindings.MPIN_BLS461_SERVER_KEY.
func ServerKeyBLS461(h int, Z []byte, SS []byte, w []byte, p []byte, I []byte, U []byte, UT []byte) (K []byte, err error) {

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

	KSize := bindings.PASBLS461
	KOct := bindings.MakeOctet(KSize)
	defer KOct.Free()

	err = bindings.ServerKeyBLS461(h, ZOct, SSOct, wOct, pOct, IOct, UOct, UTOct, KOct)

	K = KOct.ToBytes()

	return
}

// ServerBLS461 is a wrapper of bindings.MPIN_BLS461_SERVER.
func ServerBLS461(h int, d int, SS []byte, U []byte, UT []byte, V []byte, ID []byte, MESSAGE []byte, t int, Pa []byte) (HID []byte, HTID []byte, y []byte, err error) {

	HIDSize := bindings.G1SBLS461
	HIDOct := bindings.MakeOctet(HIDSize)
	defer HIDOct.Free()

	HTIDSize := bindings.G1SBLS461
	HTIDOct := bindings.MakeOctet(HTIDSize)
	defer HTIDOct.Free()

	ySize := bindings.PGSBLS461
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

	err = bindings.ServerBLS461(h, d, HIDOct, HTIDOct, yOct, SSOct, UOct, UTOct, VOct, nil, nil, IDOct, MESSAGEOct, t, PaOct)

	HID = HIDOct.ToBytes()

	HTID = HTIDOct.ToBytes()

	y = yOct.ToBytes()

	return
}

// Server1BLS461 is a wrapper of bindings.MPIN_BLS461_SERVER_1.
func Server1BLS461(h int, d int, ID []byte) (HID []byte, HTID []byte) {

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	HIDSize := bindings.G1SBLS461
	HIDOct := bindings.MakeOctet(HIDSize)
	defer HIDOct.Free()

	HTIDSize := bindings.G1SBLS461
	HTIDOct := bindings.MakeOctet(HTIDSize)
	defer HTIDOct.Free()

	bindings.Server1BLS461(h, d, IDOct, HIDOct, HTIDOct)

	HID = HIDOct.ToBytes()

	HTID = HTIDOct.ToBytes()

	return
}
