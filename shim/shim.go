package shim

import (
	"crypto/ed25519"
	"crypto/rand"

	sdk "github.com/r2pq-suite/r2pq-sdk/crypto/pq"
)

// signer/verifier are lightweight stubs behind SDK interfaces.
type signer struct {
	pri ed25519.PrivateKey
	pub ed25519.PublicKey
}

type verifier struct {
	pub ed25519.PublicKey
}

func init() {
	// Register “shim” algorithms in the SDK registry.
	sdk.Register(sdk.AlgoShimSig, newSigner, newVerifier)
}

func newSigner() (sdk.Signer, error) {
	pub, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &signer{pri: pri, pub: pub}, nil
}

func newVerifier(pubRaw []byte) (sdk.Verifier, error) {
	return &verifier{pub: ed25519.PublicKey(pubRaw)}, nil
}

// --- sdk.Signer / sdk.Verifier minimal methods ---

func (s *signer) PublicKey() []byte { return append([]byte(nil), s.pub...) }

func (s *signer) Sign(msg []byte) ([]byte, error) {
	sig := ed25519.Sign(s.pri, msg)
	return sig, nil
}

func (v *verifier) Verify(msg, sig []byte) bool {
	return ed25519.Verify(v.pub, msg, sig)
}
