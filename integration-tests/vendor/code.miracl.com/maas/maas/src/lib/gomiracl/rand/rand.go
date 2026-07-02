// Package rand defines a cryptographically secure random number generator.
package rand

import (
	"io"

	"code.miracl.com/maas/maas/src/lib/gomiracl"
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

// Rand is a cryptographically secure random number generator.
// Implements io.Reader.
type Rand bindings.Rand

// New is Rand.
func New(r io.Reader, seedSize int) (*Rand, error) {
	seed := make([]byte, seedSize)
	defer gomiracl.CleanMemory(seed)

	_, err := io.ReadFull(r, seed)
	if err != nil {
		return nil, err
	}

	return (*Rand)(bindings.NewRand(seed)), nil
}

// GetByte returns one random byte.
func (rng *Rand) GetByte() byte {
	return byte(wrap.GetRandByte((*bindings.Rand)(rng)))
}

// Read generates len(p) random bytes and writes them into p. It
// always returns len(p) and a nil error.
// Read should not be called concurrently with any other Rand method.
func (rng *Rand) Read(p []byte) (n int, err error) {
	for x := 0; x < len(p); x++ {
		p[x] = rng.GetByte()
	}

	return len(p), nil
}
