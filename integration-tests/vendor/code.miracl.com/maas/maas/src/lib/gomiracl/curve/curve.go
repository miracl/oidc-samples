// Package curve define elliptic curves.
package curve

import (
	"code.miracl.com/maas/maas/src/lib/gomiracl"
	"code.miracl.com/maas/maas/src/lib/gomiracl/bindings"
	"code.miracl.com/maas/maas/src/lib/gomiracl/rand"
)

// All are all the curves the package is currently built with.
var All = map[string]Curve{}

// AllNames are all the curves the package supports.
var AllNames = []string{
	"ANSSI",
	"BLS383",
	"BLS461",
	"BN254",
	"BN254CX",
	"BRAINPOOL",
	"C41417",
	"ED25519",
	"FP256BN",
	"FP512BN",
	"GOLDILOCKS",
	"HIFIVE",
	"NIST256",
	"NIST384",
	"NIST521",
	"NUMS256E",
	"NUMS256W",
	"NUMS384E",
	"NUMS384W",
	"NUMS512E",
	"NUMS512W",
}

// Curve is elliptic curve.
type Curve interface {
	Name() string
	EGroupSize() int
	GetKeyPair(rng *rand.Rand, s []byte) (priv, pub []byte, err error)
	ValidatePubKey(pub []byte) (err error)
	Sign(hash gomiracl.Hash, rng *rand.Rand, k []byte, priv []byte, msg []byte) (c, d []byte, err error)
	Verify(hash gomiracl.Hash, pub, msg, c, d []byte) (err error)
}

// Curve is elliptic curve.
type amclCurve struct {
	name string
	EGS  int

	getKeyPair     func(rng *bindings.Rand, s []byte) (priv, pub []byte, err error)
	validatePubKey func(pub []byte) (err error)
	sign           func(hashType int, rng *bindings.Rand, k []byte, priv []byte, msg []byte) (c, d []byte, err error)
	verify         func(hashType int, pub, msg, c, d []byte) (err error)
}

// Name is the name of the curve.
func (c *amclCurve) Name() string {
	return c.name
}

// EGroupSize is the curve EGS.
func (c *amclCurve) EGroupSize() int {
	return c.EGS
}

// GetKeyPair generates ECC public/private key pair.
func (c *amclCurve) GetKeyPair(rng *rand.Rand, s []byte) (priv, pub []byte, err error) {
	return c.getKeyPair((*bindings.Rand)(rng), s)
}

// ValidatePubKey validates ECC public key.
func (c *amclCurve) ValidatePubKey(pub []byte) (err error) {
	return c.validatePubKey(pub)
}

// Sign generates ECDSA Signature.
func (c *amclCurve) Sign(hash gomiracl.Hash, rng *rand.Rand, k, priv, msg []byte) (s, d []byte, err error) {
	return c.sign(int(hash), (*bindings.Rand)(rng), k, priv, msg)
}

// Verify verifies ECDSA Signature.
func (c *amclCurve) Verify(hash gomiracl.Hash, pub, msg, s, d []byte) (err error) {
	return c.verify(int(hash), pub, msg, s, d)
}

// mpinCurve is elliptic curve that supports MPin.
type mpinCurve struct {
	Curve

	PFS int
	PGS int

	getRandomSecret      func(rng *bindings.Rand) (secret []byte, err error)
	newServerSecretShare func(masterSecret []byte) (share []byte, err error)
	newClientSecretShare func(masterSecret, id []byte) (share []byte, err error)
	newTimePermitShare   func(hashType, date int, masterSecret, id []byte) (timePermit []byte, err error)
	recombineG1          func(share1, share2 []byte) (secret []byte, err error)
	recombineG2          func(share1, share2 []byte) (secret []byte, err error)
	extractPin           func(hashType int, id []byte, pin int, clientSecret []byte) (token []byte, err error)
	clientOnePass        func(hashType, date int, id []byte, rng *bindings.Rand, x []byte, pin int, token, timePermit, msg []byte, time int) (xOut, v, u, ut, y []byte, err error)
	serverOnePass        func(hashType, date int, serverSecret, u, ut, v, id, msg []byte, time int, pa []byte) (hid, htid, y []byte, err error)
	clientPass1          func(hashType, date int, id []byte, rng *bindings.Rand, x []byte, pin int, token, timePermit []byte) (xOut, sec, u, ut []byte, err error)
	clientPass2          func(x, y, sec []byte) (vOut []byte, err error)
	serverPass1          func(hashType, date int, id []byte) (hid, htid []byte)
	serverPass2          func(date int, hid, htid, y, serverSecret, u, ut, vOut, pa []byte) (err error)
	precompute           func(token, id, pub []byte) (g1, g2 []byte, err error)
	getG1Multiple        func(rng *bindings.Rand, opType int, x, g []byte) (xOut, w []byte, err error)
	serverKey            func(hashType int, z, serverSecret, w, p, i, u, ut []byte) (k []byte, err error)
	clientKey            func(hashType int, g1, g2 []byte, pin int, r, x, p, t []byte) (k []byte, err error)
	getDVSKeyPair        func(rng *bindings.Rand, z []byte) (zOut, pub []byte, err error)

	// Wang Chow-Choo
	wccRandomGenerate func(rng *bindings.Rand) (secret []byte, err error)
	wccHq             func(hashType int, a, b, c, d []byte) []byte
	wccGetG1Multiple  func(s, hid []byte) (vG1 []byte, err error)
	wccRecombineG1    func(share1, share2 []byte) (whole []byte, err error)
	wccGetG2Multiple  func(s, hid []byte) (vG2 []byte, err error)
	wccRecombineG2    func(share1, share2 []byte) (whole []byte, err error)
	wccReceiverKey    func(hashType int, y, w, pia, pib, paG1, pgG1, bG2Key, aID []byte) (key []byte, err error)
	wccSenderKey      func(hashType int, x, pia, pib, pbG2, pgG1, aG1Key, bID []byte) (key []byte, err error)
}

// FieldSize is the curves field size.
func (c *mpinCurve) FieldSize() int {
	return c.PFS
}

// PGroupSize is the curve PGS.
func (c *mpinCurve) PGroupSize() int {
	return c.PGS
}

// GetRandomSecret generates random secret.
func (c *mpinCurve) GetRandomSecret(rng *rand.Rand) (secret []byte, err error) {
	return c.getRandomSecret((*bindings.Rand)(rng))
}

// NewServerSecretShare returns new server secret share.
func (c *mpinCurve) NewServerSecretShare(masterSecret []byte) (share []byte, err error) {
	return c.newServerSecretShare(masterSecret)
}

// NewClientSecretShare returns new client secret share.
func (c *mpinCurve) NewClientSecretShare(masterSecret, id []byte) (share []byte, err error) {
	return c.newClientSecretShare(masterSecret, id)
}

// NewTimePermitShare returns new time permit share.
func (c *mpinCurve) NewTimePermitShare(hash gomiracl.Hash, date int, masterSecret, id []byte) (timePermit []byte, err error) {
	return c.newTimePermitShare(int(hash), date, masterSecret, id)
}

// RecombineG1 returns the sum of members from the group G1.
func (c *mpinCurve) RecombineG1(share1, share2 []byte) (secret []byte, err error) {
	return c.recombineG1(share1, share2)
}

// RecombineG2 returns the sum of members from the group G2.
func (c *mpinCurve) RecombineG2(share1, share2 []byte) (secret []byte, err error) {
	return c.recombineG2(share1, share2)
}

// ExtractPin extracts a PIN number from a client secret.
func (c *mpinCurve) ExtractPin(hash gomiracl.Hash, id []byte, pin int, clientSecret []byte) (token []byte, err error) {
	return c.extractPin(int(hash), id, pin, clientSecret)
}

// ClientOnePass performs client side of the one-pass version of the M-Pin
// protocol.
//
//nolint:gocritic // needs refactoring
func (c *mpinCurve) ClientOnePass(hash gomiracl.Hash, date int, id []byte, rng *rand.Rand, x []byte, pin int, token, timePermit, msg []byte, time int) (xOut, v, u, ut, y []byte, err error) {
	return c.clientOnePass(int(hash), date, id, (*bindings.Rand)(rng), x, pin, token, timePermit, msg, time)
}

// ServerOnePass performs server side of the one-pass version of the M-Pin protocol.
func (c *mpinCurve) ServerOnePass(hash gomiracl.Hash, date int, serverSecret, u, ut, v, id, msg []byte, time int, pa []byte) (hid, htid, y []byte, err error) {
	return c.serverOnePass(int(hash), date, serverSecret, u, ut, v, id, msg, time, pa)
}

// ClientPass1 performs first pass of the client side of the 3-pass version of
// the M-Pin protocol.
func (c *mpinCurve) ClientPass1(hash gomiracl.Hash, date int, id []byte, rng *rand.Rand, x []byte, pin int, token, timePermit []byte) (xOut, sec, u, ut []byte, err error) {
	return c.clientPass1(int(hash), date, id, (*bindings.Rand)(rng), x, pin, token, timePermit)
}

// ClientPass2 performs server side of the one-pass version of the M-Pin
// protocol.
func (c *mpinCurve) ClientPass2(x, y, sec []byte) (vOut []byte, err error) {
	return c.clientPass2(x, y, sec)
}

// ServerPass1 performs first pass of the server side of the 3-pass version of
// the M-Pin protocol.
func (c *mpinCurve) ServerPass1(hash gomiracl.Hash, date int, id []byte) (hid, htid []byte) {
	return c.serverPass1(int(hash), date, id)
}

// ServerPass2 performs third pass on the server side of the 3-pass version of
// the M-Pin protocol.
func (c *mpinCurve) ServerPass2(date int, hid, htid, y, serverSecret, u, ut, vOut, pa []byte) (err error) {
	return c.serverPass2(date, hid, htid, y, serverSecret, u, ut, vOut, pa)
}

// Precompute precomputes values for use by the client side of M-Pin Full.
func (c *mpinCurve) Precompute(token, id, pub []byte) (g1, g2 []byte, err error) {
	return c.precompute(token, id, pub)
}

// GetG1Multiple finds a random multiple of a point in G1.
func (c *mpinCurve) GetG1Multiple(rng *rand.Rand, opType int, x, g []byte) (xOut, w []byte, err error) {
	return c.getG1Multiple((*bindings.Rand)(rng), opType, x, g)
}

// ServerKey calculates key on Server side for M-Pin Full.
func (c *mpinCurve) ServerKey(hash gomiracl.Hash, z, serverSecret, w, p, i, u, ut []byte) (k []byte, err error) {
	return c.serverKey(int(hash), z, serverSecret, w, p, i, u, ut)
}

// ClientKey calculates key on Client side for M-Pin Full.
func (c *mpinCurve) ClientKey(hash gomiracl.Hash, g1, g2 []byte, pin int, r, x, p, t []byte) (k []byte, err error) {
	return c.clientKey(int(hash), g1, g2, pin, r, x, p, t)
}

// GetDVSKeyPair generates a random public key for the client.
func (c *mpinCurve) GetDVSKeyPair(rng *rand.Rand, z []byte) (zOut, pub []byte, err error) {
	return c.getDVSKeyPair((*bindings.Rand)(rng), z)
}

// WCCRandomGenerate generate random secret for WCC.
func (c *mpinCurve) WCCGetRandomSecret(rng *rand.Rand) (secret []byte, err error) {
	return c.wccRandomGenerate((*bindings.Rand)(rng))
}

// WCCHq hashes ec points and id to an integer.
func (c *mpinCurve) WCCHash(hash gomiracl.Hash, a, b, ci, d []byte) []byte {
	return c.wccHq(int(hash), a, b, ci, d)
}

// WCCGetG1Multiple returns value in G1 multiplied by an integer.
func (c *mpinCurve) WCCGetG1Multiple(s, hid []byte) (vG1 []byte, err error) {
	return c.wccGetG1Multiple(s, hid)
}

// WCCRecombineG1 returns the addition of two members from the group G1.
func (c *mpinCurve) WCCRecombineG1(share1, share2 []byte) (whole []byte, err error) {
	return c.wccRecombineG1(share1, share2)
}

// WCCGetG2Multiple returns value in G2 multiplied by an integer.
func (c *mpinCurve) WCCGetG2Multiple(s, hid []byte) (vG2 []byte, err error) {
	return c.wccGetG2Multiple(s, hid)
}

// WCCRecombineG2 returns the addition of two members from the group G2.
func (c *mpinCurve) WCCRecombineG2(share1, share2 []byte) (whole []byte, err error) {
	return c.wccRecombineG2(share1, share2)
}

// WCCReceiverKey returns the shared key on the receiver side.
func (c *mpinCurve) WCCReceiverKey(hash gomiracl.Hash, y, w, pia, pib, paG1, pgG1, bG2Key, aID []byte) (key []byte, err error) {
	return c.wccReceiverKey(int(hash), y, w, pia, pib, paG1, pgG1, bG2Key, aID)
}

// WCCSenderKey returns the shared key on the sender side.
func (c *mpinCurve) WCCSenderKey(hash gomiracl.Hash, x, pia, pib, pbG2, pgG1, aG1Key, bID []byte) (key []byte, err error) {
	return c.wccSenderKey(int(hash), x, pia, pib, pbG2, pgG1, aG1Key, bID)
}
