package shim

import "testing"

func TestRoundTrip(t *testing.T) {
	s, err := newSigner()
	if err != nil {
		t.Fatal(err)
	}
	v, err := newVerifier()
	if err != nil {
		t.Fatal(err)
	}

	msg := []byte("r2pq")
	sig, err := s.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}
	if !v.Verify(s.PublicKeyHex(), msg, sig) {
		t.Fatal("verify failed")
	}
}
