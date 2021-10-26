package bindings

// #include "amcl/amcl.h"
// #include "amcl/randapi.h"
import "C"

// Rand is a cryptographically secure random number generator
type Rand C.csprng

// NewRand create new seeded Rand
func NewRand(seed []byte) *Rand {
	var rng C.csprng

	sOct := NewOctet(seed)
	defer sOct.Free()

	C.CREATE_CSPRNG(&rng, sOct)

	return (*Rand)(&rng)
}

// GetByte returns one random byte
func (rng *Rand) GetByte() byte {
	r := C.RAND_byte((*C.csprng)(rng))

	return byte(r)
}

func (rng *Rand) csprng() *C.csprng {
	return (*C.csprng)(rng)
}
