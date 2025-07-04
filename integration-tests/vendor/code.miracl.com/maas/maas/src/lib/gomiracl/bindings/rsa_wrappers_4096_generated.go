// Generated by gen/wrappers/main.go from wrap/rsa_wrappers.go.tmpl.

package bindings

// #cgo LDFLAGS: -lamcl_rsa_4096
// #include "amcl/rsa_4096.h"
// #include "amcl/rsa_support.h"
import "C"

// RSA constants
const (
	RFS4096             = int(C.RFS_4096)           // RFS4096 is the RSA Public Key Size in bytes
	FFLEN4096           = int(C.FFLEN_4096)         // FFLEN4096 consists in 2^n multiplier of BIGBITS to specify supported Finite Field size, e.g 2048=256*2^3 where BIGBITS=256
	HashTypeRSA4096 int = int(C.HASH_TYPE_RSA_4096) // HashTypeRSA4096 is the chosen Hash algorithm
)

// NewRSAPrivateKey4096 creates new RSA 4096 Private key
func NewRSAPrivateKey4096() RSAPrivateKey {
	return &C.rsa_private_key_4096{}
}

// NewRSAPublicKey4096 creates new RSA 4096 public key
func NewRSAPublicKey4096() RSAPublicKey {
	return &C.rsa_public_key_4096{}
}

// RSADecrypt4096 is a go wrapper for C.RSA_4096_DECRYPT.
func RSADecrypt4096(priv RSAPrivateKey, G *Octet, F *Octet) {
	var GCOct *C.octet = nil
	if G != nil {
		GCOct = (*C.octet)(&G.cOctet)
	}

	var FCOct *C.octet = nil
	if F != nil {
		FCOct = (*C.octet)(&F.cOctet)
	}

	C.RSA_4096_DECRYPT(priv.(*C.rsa_private_key_4096), GCOct, FCOct)
}

// RSAEncrypt4096 is a go wrapper for C.RSA_4096_ENCRYPT.
func RSAEncrypt4096(pub RSAPublicKey, F *Octet, G *Octet) {
	var FCOct *C.octet = nil
	if F != nil {
		FCOct = (*C.octet)(&F.cOctet)
	}

	var GCOct *C.octet = nil
	if G != nil {
		GCOct = (*C.octet)(&G.cOctet)
	}

	C.RSA_4096_ENCRYPT(pub.(*C.rsa_public_key_4096), FCOct, GCOct)
}

// RSAKeyPair4096 is a go wrapper for C.RSA_4096_KEY_PAIR.
func RSAKeyPair4096(rng *Rand, e int32, priv RSAPrivateKey, pub RSAPublicKey, p *Octet, q *Octet) {

	var pCOct *C.octet = nil
	if p != nil {
		pCOct = (*C.octet)(&p.cOctet)
	}

	var qCOct *C.octet = nil
	if q != nil {
		qCOct = (*C.octet)(&q.cOctet)
	}

	C.RSA_4096_KEY_PAIR((*C.csprng)(rng), C.sign32(e), priv.(*C.rsa_private_key_4096), pub.(*C.rsa_public_key_4096), pCOct, qCOct)
}

// RSAPrivateKeyKill4096 is a go wrapper for C.RSA_4096_PRIVATE_KEY_KILL.
func RSAPrivateKeyKill4096(PRIV RSAPrivateKey) {

	C.RSA_4096_PRIVATE_KEY_KILL(PRIV.(*C.rsa_private_key_4096))
}
