package bindings

import "C"
import (
	"fmt"
)

var (
	// ECDHInvalidSignature is binding to C.ECDH_INVALID
	ECDHInvalidSignature = -4
	// ECDHInvalidPubKey is binding to C.ECDH_INVALID_PUBLIC_KEY
	ECDHInvalidPubKey = -3

	// InvalidPoint is binding to C.MPIN_INVALID_POINT
	InvalidPoint = -14
	// BadPin is binding to C.MPIN_BAD_PIN
	BadPin = -19
)

var errorStrings = map[int]string{
	InvalidPoint: "Invalid point",

	ECDHInvalidSignature: "Invalid ECDH signature",
	ECDHInvalidPubKey:    "Invalid ECDH public key",

	BadPin: "Bad PIN number entered",
}

// Error is for errors returned from AMCL wrappers
type Error struct {
	Code int
}

func (err *Error) Error() string {
	errStr := fmt.Sprintf("amcl: return code %v", err.Code)
	if str, ok := errorStrings[err.Code]; ok {
		errStr = fmt.Sprintf("%v %v", errStr, str)
	}

	return errStr
}

func newError(code C.int) error {
	if code == 0 {
		return nil
	}

	return &Error{int(code)}
}
