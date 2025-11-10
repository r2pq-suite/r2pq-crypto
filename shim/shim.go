package shim

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"

	sdk "github.com/r2pq-suite/r2pq-sdk"
)

// signer implements the sdk.Signer expected by r2pq-sdk.
type signer struct {
	pri ed25519.PrivateKey
	pub ed25519.PublicKey
}

// verifier implements the sdk.Verifier expected by r2pq-sdk.
type verifier struct{}

func init() {
	// Make this backend discoverable by the SDK as a placeholder "PQ" algo.
	// (The constant name lives in r2pq-sdk; this keeps the coupling minimal.)
	sdk.Register(sdk.AlgoShimSig, newSigner, newVerifier)
}

func newSigner() (sdk.Signer, error) {
	pub, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &signer{pri: pri, pub: pub}, nil
}

func newVerifier() (sdk.Verifier, error) {
	return &verifier{}, nil
}

// --- sdk.Signer ---

func (s *signer) PublicKeyHex() string {
	return hex.EncodeToString(s.pub)
}

func (s *signer) Sign(msg []byte) ([]byte, error) {
	return ed25519.Sign(s.pri, msg), nil
}

// --- sdk.Verifier ---

func (v *verifier) Verify(pubHex string, msg, sig []byte) bool {
	pubBytes, err := hex.DecodeString(pubHex)
	if err != nil {
		return false
	}
	return ed25519.Verify(ed25519.PublicKey(pubBytes), msg, sig)
}
