# r2pq-crypto

Reference crypto backends for the R2PQ suite.

- ✅ Buildable shim backend (ed25519) used only for wiring/tests
- 🔌 Registers itself with `r2pq-sdk` on import
- 🧪 `go test ./...` passes on fresh clone

> Real PQ algos (Kyber/Dilithium/etc.) will drop in later behind the same interfaces.
