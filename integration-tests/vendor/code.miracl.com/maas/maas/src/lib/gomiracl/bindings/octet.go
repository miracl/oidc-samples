package bindings

// #include "amcl/amcl.h"
// #include "stdlib.h"
// #include "string.h"
import "C" //nolint:gocritic // not now
import (
	"unsafe" //nolint:gocritic // not now
)

// Octet is Go alias for the C octet type.
type Octet struct {
	cOctet C.octet
}

// NewOctet creates new Octet with given value.
func NewOctet(val []byte) *Octet {
	if val == nil {
		return nil
	}

	return &Octet{
		C.octet{
			len: C.int(len(val)),
			max: C.int(len(val)),
			val: C.CString(string(val)),
		},
	}
}

// MakeOctet create empty Octet.
func MakeOctet(m int) *Octet {
	return &Octet{
		C.octet{
			len: C.int(m),
			max: C.int(m),
			val: (*C.char)(C.calloc(1, C.size_t(m))),
		},
	}
}

// ToBytes returns the bytes representation of the Octet.
func (o *Octet) ToBytes() []byte {
	return C.GoBytes(unsafe.Pointer(o.cOctet.val), o.cOctet.len) //nolint:nlreturn // mystery
}

// Free frees the allocated memory.
func (o *Octet) Free() {
	if o == nil {
		return
	}

	C.free(unsafe.Pointer(o.cOctet.val))
}

// Clean overwrites the memory before freeing it.
func (o *Octet) Clean() {
	if o == nil {
		return
	}

	C.OCT_clear(&o.cOctet) //nolint:gocritic // false positive
}
