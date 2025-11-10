package shim

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"

	sdk "github.com/r2pq-suite/r2pq-sdk/keys"
)

type signer struct {
	pri ed25519.PrivateKey
	pub ed25519.PublicKey
}

type verifier struct{}

func init() {
	sdk.Register(sdk.AlgoShimSig, newSigner, verifier{})
}

func newSigner() (sdk.Signer, error) {
	pub, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &signer{pri: pri, pub: pub}, nil
}

func (s *signer) Algorithm() sdk.Algorithm {
	return sdk.AlgoShimSig
}

func (s *signer) Public() sdk.PublicKey {
	return sdk.PublicKey(s.pub)
}

func (s *signer) Address() sdk.Address {
	return sdk.Address(hex.EncodeToString(s.pub))
}

func (s *signer) Sign(msg []byte) (sdk.Signature, error) {
	sig := ed25519.Sign(s.pri, msg)
	return sdk.Signature(sig), nil
}

// ---------------------
// Verifier implementation
// ---------------------

func (verifier) Algorithm() sdk.Algorithm {
	return sdk.AlgoShimSig
}

func (verifier) Verify(pub sdk.PublicKey, msg []byte, sig sdk.Signature) bool {
	return ed25519.Verify(ed25519.PublicKey(pub), msg, sig)
}

func (verifier) AddressFrom(pub sdk.PublicKey) sdk.Address {
	return sdk.Address(hex.EncodeToString(pub))
}
