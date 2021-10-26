package bindings

// #include "amcl/amcl.h"
import "C"

// Hash types
const (
	SHA256 = C.SHA256
	SHA384 = C.SHA384
	SHA512 = C.SHA512
)
