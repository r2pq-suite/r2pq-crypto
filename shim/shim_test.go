package shim

import (
	"testing"

	// Importing the SDK package ensures the module graph is correct
	_ "github.com/r2pq-suite/r2pq-sdk/crypto/pq"
)

func TestSmoke(t *testing.T) {
	// Build/link smoke test: if imports resolve and this test runs, we're good.
	// We'll expand this to real sign/verify round-trips once algorithms land.
}
