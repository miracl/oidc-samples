// Generated by gen/wrappers/main.go from mpin_wrappers.go.tmpl.

//go:build !ignoredefaultcurves || BLS383

package wrap

import "code.miracl.com/mfa/pkg/gomiracl/bindings"

// Server2BLS383Kangaroo is a wrapper of bindings.MPIN_BLS383_SERVER_2.
func Server2BLS383Kangaroo(d int, HID []byte, HTID []byte, y []byte, SS []byte, U []byte, UT []byte, V []byte, Pa []byte) (E []byte, F []byte, err error) {

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

	ESize := bindings.GTSBLS383
	EOct := bindings.MakeOctet(ESize)
	defer EOct.Free()

	FSize := bindings.GTSBLS383
	FOct := bindings.MakeOctet(FSize)
	defer FOct.Free()

	PaOct := bindings.NewOctet(Pa)
	defer PaOct.Free()

	err = bindings.Server2BLS383(d, HIDOct, HTIDOct, yOct, SSOct, UOct, UTOct, VOct, EOct, FOct, PaOct)

	E = EOct.ToBytes()

	F = FOct.ToBytes()

	return
}

// ServerBLS383Kangaroo is a wrapper of bindings.MPIN_BLS383_SERVER.
func ServerBLS383Kangaroo(h int, d int, SS []byte, U []byte, UT []byte, V []byte, ID []byte, MESSAGE []byte, t int, Pa []byte) (HID []byte, HTID []byte, y []byte, E []byte, F []byte, err error) {

	HIDSize := bindings.G1SBLS383
	HIDOct := bindings.MakeOctet(HIDSize)
	defer HIDOct.Free()

	HTIDSize := bindings.G1SBLS383
	HTIDOct := bindings.MakeOctet(HTIDSize)
	defer HTIDOct.Free()

	ySize := bindings.PGSBLS383
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

	ESize := bindings.GTSBLS383
	EOct := bindings.MakeOctet(ESize)
	defer EOct.Free()

	FSize := bindings.GTSBLS383
	FOct := bindings.MakeOctet(FSize)
	defer FOct.Free()

	IDOct := bindings.NewOctet(ID)
	defer IDOct.Free()

	MESSAGEOct := bindings.NewOctet(MESSAGE)
	defer MESSAGEOct.Free()

	PaOct := bindings.NewOctet(Pa)
	defer PaOct.Free()

	err = bindings.ServerBLS383(h, d, HIDOct, HTIDOct, yOct, SSOct, UOct, UTOct, VOct, EOct, FOct, IDOct, MESSAGEOct, t, PaOct)

	HID = HIDOct.ToBytes()

	HTID = HTIDOct.ToBytes()

	y = yOct.ToBytes()

	E = EOct.ToBytes()

	F = FOct.ToBytes()

	return
}
