package shim

import (
	"crypto/ed25519"
	"crypto/rand"

	sdk "github.com/r2pq-suite/r2pq-sdk"
)

// signer implements sdk.Signer
type signer struct {
	pri ed25519.PrivateKey
	pub ed25519.PublicKey
}

// verifier implements sdk.Verifier
type verifier struct{}

func init() {
	// Register this shim with the SDK under a placeholder algo name.
	// Use whatever constant/alias your SDK exposes. Example:
	//   sdk.Register(sdk.AlgoShimSig, newSigner, newVerifier)
	// If your SDK exposes a string name instead, do:
	//   sdk.RegisterByName("shim-ed25519", newSigner, newVerifier)
	sdk.Register(sdk.AlgoShimSig, newSigner, newVerifier)
}

func newSigner() (sdk.Signer, error) {
	pub, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &signer{pri: pri, pub: pub}, nil
}

func (s *signer) PublicKey() []byte { return s.pub }

func (s *signer) Sign(msg []byte) ([]byte, error) {
	sig := ed25519.Sign(s.pri, msg)
	return sig, nil
}

func newVerifier() (sdk.Verifier, error) { return &verifier{}, nil }

func (v *verifier) Verify(pub, msg, sig []byte) bool {
	return ed25519.Verify(ed25519.PublicKey(pub), msg, sig)
}
