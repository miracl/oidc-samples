// Generated by gen/wrappers/main.go from wrap/pbc_wrappers.go.tmpl.

package bindings

// #include "amcl/ecdh_support.h"
// #include "amcl/pbc_support.h"
// #include "amcl/utils.h"
import "C"

// PBKDF2 is a go wrapper for C.PBKDF2.
func PBKDF2(h int, P *Octet, S *Octet, rep int, len int, K *Octet) {
	C.PBKDF2(C.int(h), (*C.octet)(P), (*C.octet)(S), C.int(rep), C.int(len), (*C.octet)(K))
}

// HashID is a go wrapper for C.HASH_ID.
func HashID(h int, ID *Octet, HID *Octet) {
	C.HASH_ID(C.int(h), (*C.octet)(ID), (*C.octet)(HID))
}

// HashAll is a go wrapper for C.HASH_ALL.
func HashAll(h int, I *Octet, U *Octet, CU *Octet, Y *Octet, V *Octet, R *Octet, W *Octet, H *Octet) {
	C.HASH_ALL(C.int(h), (*C.octet)(I), (*C.octet)(U), (*C.octet)(CU), (*C.octet)(Y), (*C.octet)(V), (*C.octet)(R), (*C.octet)(W), (*C.octet)(H))
}

// GenerateRandom is a go wrapper for C.generateRandom.
func GenerateRandom(RNG *Rand, randomValue *Octet) {
	C.generateRandom((*C.csprng)(RNG), (*C.octet)(randomValue))
}

// AESGCMEncrypt is a go wrapper for C.AES_GCM_ENCRYPT.
func AESGCMEncrypt(K *Octet, IV *Octet, H *Octet, P *Octet, C *Octet, T *Octet) {
	C.AES_GCM_ENCRYPT((*C.octet)(K), (*C.octet)(IV), (*C.octet)(H), (*C.octet)(P), (*C.octet)(C), (*C.octet)(T))
}

// AESGCMDecrypt is a go wrapper for C.AES_GCM_DECRYPT.
func AESGCMDecrypt(K *Octet, IV *Octet, H *Octet, C *Octet, P *Octet, T *Octet) {
	C.AES_GCM_DECRYPT((*C.octet)(K), (*C.octet)(IV), (*C.octet)(H), (*C.octet)(C), (*C.octet)(P), (*C.octet)(T))
}

// GenerateOTP is a go wrapper for C.generateOTP.
func GenerateOTP(RNG *Rand) int {
	return int(C.generateOTP((*C.csprng)(RNG)))
}
