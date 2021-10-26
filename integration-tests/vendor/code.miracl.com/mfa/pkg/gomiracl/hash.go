package gomiracl

import "code.miracl.com/mfa/pkg/gomiracl/bindings"

// Hash is Milagro hash function.
type Hash int

// Supported hash functions.
const (
	SHA256 Hash = bindings.SHA256 // sha256
	SHA384 Hash = bindings.SHA384 // sha384
	SHA512 Hash = bindings.SHA512 // sha512
)
