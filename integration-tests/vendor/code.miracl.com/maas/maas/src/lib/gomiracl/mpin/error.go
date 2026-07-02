package mpin

import "code.miracl.com/maas/maas/src/lib/gomiracl/bindings"

// IsWrongPin returns true if the err is Wrong PIN.
func IsWrongPin(err error) bool {
	amclError, ok := err.(*bindings.Error) //nolint:errorlint // needs refactoring

	return ok && amclError.Code == bindings.BadPin
}
