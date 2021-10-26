package bindings

// #include "amcl/amcl.h"
// #include "stdlib.h"
// #include "string.h"
import "C"
import "unsafe"

// Octet is Go alias for the C octet type.
type Octet = C.octet

// NewOctet creates new Octet with given value.
func NewOctet(val []byte) *Octet {
	if val == nil {
		return nil
	}

	return &Octet{
		len: C.int(len(val)),
		max: C.int(len(val)),
		val: C.CString(string(val)),
	}
}

// MakeOctet create empty Octet.
func MakeOctet(max int) *Octet {
	return &Octet{
		len: C.int(max),
		max: C.int(max),
		val: (*C.char)(C.calloc(1, C.size_t(max))),
	}
}

// ToBytes returns the bytes representation of the Octet.
func (o *Octet) ToBytes() []byte {
	return C.GoBytes(unsafe.Pointer(o.val), o.len) //nolint:nlreturn // mistery
}

// Free frees the allocated memory.
func (o *Octet) Free() {
	if o == nil {
		return
	}

	C.free(unsafe.Pointer(o.val))
}

// Clean overwrites the memory before freeing it.
func (o *Octet) Clean() {
	if o == nil {
		return
	}

	C.OCT_clear(o)
}
