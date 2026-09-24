> **Historical 2025 R2PQ prototype lineage.** This repository preserves
> earlier exploratory work. It does not represent the architecture or
> capabilities of the current R2PQ cryptographic-assessment research system.
> Documentation below records historical implementation, limitations,
> and intentions; planned features must not be treated as delivered.
> Current public scope and evidence boundaries: [r2pq.dev](https://r2pq.dev).

## Historical prototype documentation — 2025

# r2pq-crypto

Reference crypto shim(s) for the R2PQ suite. This repo compiles standalone and registers a
signature/verifier with `r2pq-sdk` at init time.

## Layout
- `shim/` — ed25519-based placeholder shim (drop-in PQ backends later)
- `.github/workflows/ci.yml` — Go build + test

### Quick check
go run ./examples/basic
R2PQ wiring OK (shim linked).
