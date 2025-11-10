# r2pq-crypto

Reference crypto shim(s) for the R2PQ suite. This repo compiles standalone and registers a
signature/verifier with `r2pq-sdk` at init time.

## Layout
- `shim/` — ed25519-based placeholder shim (drop-in PQ backends later)
- `.github/workflows/ci.yml` — Go build + test

### Quick check
go run ./examples/basic
R2PQ wiring OK (shim linked).
