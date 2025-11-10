package shim

import (
	"testing"
)

func TestSignVerify(t *testing.T) {
	s, err := newSigner()
	if err != nil {
		t.Fatalf("newSigner error: %v", err)
	}

	msg := []byte("hello R2PQ")
	sig, err := s.Sign(msg)
	if err != nil {
		t.Fatalf("sign error: %v", err)
	}

	v := verifier{}
	ok := v.Verify(s.Public(), msg, sig)
	if !ok {
		t.Fatalf("verify failed")
	}
}
